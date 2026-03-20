package indexer

import (
	"time"

	"log/slog"

	"github.com/QwaQ-dev/stacks-base-bridge/pkg/logger/sl"
)

const liveLimit = 1

func (i *StacksIndexer) pollLive() {
	for {
		select {
		case <-i.ctx.Done():
			i.log.Info("Stacks live polling stopped")
			return
		case <-time.After(livePollInterval):
		}

		if err := i.pollOnce(); err != nil {
			i.log.Error("stacks poll error", sl.Err(err))
		}
	}
}

func (i *StacksIndexer) pollOnce() error {
	resp, err := i.hiro.FetchContractEvents(i.ctx, i.cfg.Stacks.Contract, 0, liveLimit)
	if err != nil {
		return err
	}

	if len(resp.Results) == 0 {
		i.log.Debug("no recent events at offset=0")
		return nil
	}

	i.log.Debug("checked most recent event",
		slog.Int("got", len(resp.Results)),
		slog.Int64("last_max_nonce", i.lastMaxNonce),
	)

	updated := false

	for _, raw := range resp.Results {
		nonce, err := processIfNewEvent(i, raw)
		if err != nil {
			i.log.Error("process recent event failed",
				slog.String("tx_id", raw.TxID),
				sl.Err(err),
			)
			continue
		}

		if nonce > i.lastMaxNonce {
			i.lastMaxNonce = nonce
			updated = true
		}
	}

	if updated {
		if err := i.db.SetLastBlock(i.ctx, "stacks", i.lastMaxNonce); err != nil {
			i.log.Error("failed to save new max nonce", sl.Err(err))
		} else {
			i.log.Info("updated max nonce after live poll",
				slog.Int64("new_max_nonce", i.lastMaxNonce))
		}
	}

	return nil
}
