package indexer

import (
	"fmt"
	"log/slog"

	"github.com/QwaQ-dev/stacks-base-bridge/pkg/logger/sl"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

// Fetching historical chunk
func (i *BaseIndexer) syncHistorical() error {
	fromBlock, err := i.db.GetLastBlock(i.ctx, i.cfg.Base.ChanID)
	if err != nil {
		return fmt.Errorf("change fromblock")
	}

	if fromBlock == 0 {
		fromBlock = i.cfg.Base.StartBlock
	}

	i.log.Debug("From block", fromBlock)

	header, err := i.client.HeaderByNumber(i.ctx, nil)
	if err != nil {
		return fmt.Errorf("get latest block: %w", err)
	}
	latestBlock := header.Number.Int64()

	if fromBlock >= latestBlock {
		i.log.Info("Historical sync: already up to date", slog.Int64("block", latestBlock))
		return nil
	}

	i.log.Info("Historical sync started",
		slog.Int64("from", fromBlock),
		slog.Int64("to", latestBlock),
	)

	for start := fromBlock; start <= latestBlock; start += chunkSize {
		select {
		case <-i.ctx.Done():
			return i.ctx.Err()
		default:
		}

		end := start + chunkSize - 1
		if end > latestBlock {
			end = latestBlock
		}

		startU := uint64(start)
		endU := uint64(end)

		i.log.Info("Fetching historical chunk",
			slog.Int64("from", start),
			slog.Int64("to", end),
		)

		iter, err := i.bridgeInstance.FilterRequestApproved(&bind.FilterOpts{
			Start:   startU,
			End:     &endU,
			Context: i.ctx,
		}, nil)
		if err != nil {
			return fmt.Errorf("filter logs [%d-%d]: %w", start, end, err)
		}

		for iter.Next() {
			event := iter.Event
			i.log.Info("[Historical] RequestApproved",
				slog.String("sender", event.Sender.Hex()),
				slog.String("amount", event.Amount.String()),
				slog.String("receiver", event.Receiver),
				slog.Uint64("block", event.Raw.BlockNumber),
			)

			err = i.db.SaveBridgeEvent(i.ctx,
				"base", "stacks",
				event.Raw.TxHash.Hex(),
				int(event.Raw.Index),
				int64(event.Raw.BlockNumber),
				event.Raw.BlockHash.Hex(),
				event.Sender.Hex(),
				event.Receiver,
				event.Amount.String(),
				event.Nonce.String(),
				"detected",
			)
			if err != nil {
				i.log.Error("failed to save historical event", sl.Err(err))
			}
		}

		if err := iter.Error(); err != nil {
			return fmt.Errorf("iter error [%d-%d]: %w", start, end, err)
		}
		iter.Close()

		if err := i.db.SetLastBlock(i.ctx, "base", end); err != nil {
			i.log.Error("failed to save last block", sl.Err(err))
		}
	}

	i.log.Info("Historical sync completed", slog.Int64("last_block", latestBlock))
	return nil
}
