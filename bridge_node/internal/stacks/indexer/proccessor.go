package indexer

import (
	"fmt"
	"log/slog"
	"regexp"
	"strconv"
	"strings"

	"github.com/QwaQ-dev/stacks-base-bridge/internal/stacks/api"
)

func processIfNewEvent(i *StacksIndexer, raw api.HiroEvent) (int64, error) {
	if raw.EventType != "smart_contract_log" || raw.ContractLog == nil || raw.ContractLog.Topic != "print" {
		return -1, nil
	}

	parsed, err := parseRepr(raw.ContractLog.Value.Repr)
	if err != nil {
		i.log.Debug("skip non-bridge print event",
			slog.String("tx_id", raw.TxID),
			slog.String("repr", raw.ContractLog.Value.Repr),
		)
		return -1, nil
	}

	if parsed.EventName != "request-approved" {
		return -1, nil
	}

	nonceInt, err := strconv.ParseInt(parsed.Nonce, 10, 64)
	if err != nil {
		return -1, err
	}

	if nonceInt <= i.lastMaxNonce {
		i.log.Debug("skipping old nonce", slog.Int64("nonce", nonceInt))
		return nonceInt, nil
	}

	txInfo, err := i.hiro.FetchTxInfo(i.ctx, raw.TxID)
	if err != nil {
		return nonceInt, err
	}

	i.log.Info("Stacks request-approved event (NEW)",
		slog.String("tx_id", raw.TxID),
		slog.String("sender", txInfo.SenderAddress),
		slog.String("receiver", parsed.Receiver),
		slog.String("amount", parsed.Amount),
		slog.String("nonce", parsed.Nonce),
		slog.Int64("block", txInfo.BlockHeight),
	)

	err = i.db.SaveBridgeEvent(
		i.ctx,
		"base", "stacks", raw.TxID, raw.EventIndex,
		txInfo.BlockHeight, "", txInfo.SenderAddress,
		parsed.Receiver, parsed.Amount, parsed.Nonce, "detected",
	)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") || strings.Contains(err.Error(), "unique constraint") {
			i.log.Info("event already exists in local DB",
				slog.String("tx_id", raw.TxID),
				slog.String("nonce", parsed.Nonce),
			)
			return nonceInt, nil
		}
		return nonceInt, err
	}

	return nonceInt, nil
}

var (
	reEvent    = regexp.MustCompile(`\(event\s+"([^"]+)"\)`)
	reAmount   = regexp.MustCompile(`\(amount\s+u(\d+)\)`)
	reSender   = regexp.MustCompile(`\(sender\s+u(\d+)\)`)
	reReceiver = regexp.MustCompile(`\(receiver\s+"([^"]+)"\)`)
	reNonce    = regexp.MustCompile(`\(nonce\s+u(\d+)\)`)
)

type StacksPrintEvent struct {
	EventName string
	Amount    string
	Receiver  string
	Nonce     string
}

func parseRepr(repr string) (*StacksPrintEvent, error) {
	if !strings.Contains(repr, "(tuple") {
		return nil, fmt.Errorf("not a tuple")
	}

	extract := func(re *regexp.Regexp) string {
		m := re.FindStringSubmatch(repr)
		if len(m) < 2 {
			return ""
		}
		return m[1]
	}

	res := &StacksPrintEvent{
		EventName: extract(reEvent),
		Amount:    extract(reAmount),
		Receiver:  extract(reReceiver),
		Nonce:     extract(reNonce),
	}

	if res.EventName == "" || res.Amount == "" || res.Receiver == "" || res.Nonce == "" {
		return nil, fmt.Errorf("missing fields in repr: %s (parsed: %+v)", repr, res)
	}

	return res, nil
}
