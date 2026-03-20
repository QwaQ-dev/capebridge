package indexer

import (
	"time"

	"log/slog"

	bridge "github.com/QwaQ-dev/stacks-base-bridge/contracts/bridge"
	"github.com/QwaQ-dev/stacks-base-bridge/pkg/logger/sl"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

const baseBackoff = 2 * time.Second

func (i *BaseIndexer) subscribeLive() {
	for {
		logs := make(chan *bridge.BridgeRequestApproved)
		sub, err := i.bridgeInstance.WatchRequestApproved(&bind.WatchOpts{
			Start:   nil,
			Context: i.ctx,
		}, logs, nil)
		if err != nil {
			i.log.Error("watch error, retrying...", sl.Err(err))
			time.Sleep(baseBackoff)
			continue
		}

		i.log.Info("Live subscription started")

		// Loop for reconnecting to RPC dial
	readLoop:
		for {
			select {
			case <-i.ctx.Done():
				sub.Unsubscribe()
				return
			case event := <-logs:
				i.log.Info("[Live] RequestApproved",
					slog.String("sender", event.Sender.Hex()),
					slog.String("amount", event.Amount.String()),
					slog.String("receiver", event.Receiver),
					slog.Uint64("block", event.Raw.BlockNumber),
				)

				err := i.db.SaveBridgeEvent(i.ctx,
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
					i.log.Error("failed to save live event", sl.Err(err))
				}

				_ = i.db.SetLastBlock(i.ctx, "base", int64(event.Raw.BlockNumber))

			case err := <-sub.Err():
				i.log.Error("subscription closed, reconnecting...", sl.Err(err))
				sub.Unsubscribe()
				break readLoop
			}
		}

		backoff := baseBackoff
		for {
			select {
			case <-i.ctx.Done():
				return
			case <-time.After(backoff):
			}
			backoff *= 2
			if backoff > 5*time.Minute {
				backoff = 5 * time.Minute
			}
		}
	}
}
