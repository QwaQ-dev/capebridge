package indexer

import (
	"time"

	"log/slog"

	"github.com/QwaQ-dev/stacks-base-bridge/pkg/logger/sl"
)

const historicalLimit = 50

func (i *StacksIndexer) syncHistorical() error {
	offset := 0
	var maxNonce int64 = i.lastMaxNonce

	i.log.Info("Stacks historical sync started",
		slog.String("contract", i.cfg.Stacks.Contract),
		slog.Int("from_offset", offset),
	)

	for {
		select {
		case <-i.ctx.Done():
			return i.ctx.Err()
		default:
		}

		resp, err := i.hiro.FetchContractEvents(i.ctx, i.cfg.Stacks.Contract, offset, historicalLimit)
		if err != nil {
			i.log.Error("fetch stacks events", slog.Int("offset", offset), sl.Err(err))
			time.Sleep(stacksBackoff)
			continue
		}

		i.log.Info("Stacks historical chunk",
			slog.Int("offset", offset),
			slog.Int("got", len(resp.Results)),
			slog.Int("total", resp.Total),
		)

		for _, raw := range resp.Results {
			nonce, err := processIfNewEvent(i, raw)
			if err != nil {
				i.log.Error("process historical event failed",
					slog.String("tx_id", raw.TxID),
					sl.Err(err),
				)
				continue
			}
			if nonce > maxNonce {
				maxNonce = nonce
			}
		}

		offset += len(resp.Results)

		if err := i.db.SetLastBlock(i.ctx, "stacks", maxNonce); err != nil {
			i.log.Error("save max nonce during historical", sl.Err(err))
		}

		if len(resp.Results) == 0 || offset >= resp.Total {
			break
		}
	}

	i.lastMaxNonce = maxNonce
	i.log.Info("Stacks historical sync completed",
		slog.Int64("max_nonce_processed", maxNonce),
		slog.Int("total_offset_processed", offset),
	)
	return nil
}
