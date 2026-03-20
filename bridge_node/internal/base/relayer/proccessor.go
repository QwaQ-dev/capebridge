package relayer

import (
	"context"

	"log/slog"

	"github.com/QwaQ-dev/stacks-base-bridge/internal/models"
	"github.com/QwaQ-dev/stacks-base-bridge/pkg/logger/sl"
	"github.com/QwaQ-dev/stacks-base-bridge/pkg/parsebigint"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

// Takes pending events from db
func (b *BaseRelayer) processPending(sourceChain string, handler func(ctx context.Context, event models.BridgeEvent) error) error {
	events, err := b.db.GetPendingEvents(b.ctx, sourceChain)
	if err != nil {
		return err
	}

	for _, event := range events {
		if err := b.processEvent(event, handler); err != nil {
			b.log.Error("Failed to process event",
				slog.Int64("event_id", event.ID),
				slog.String("tx_hash", event.TxHash),
				slog.String("nonce", event.Nonce),
				sl.Err(err),
			)
		}
	}
	return nil
}

func (b *BaseRelayer) processEvent(event models.BridgeEvent, handler func(ctx context.Context, event models.BridgeEvent) error) error {
	nonce, err := parsebigint.ParseBigInt("nonce", event.Nonce)
	if err != nil {
		return err
	}

	req, err := b.federation.Requests(&bind.CallOpts{Context: b.ctx}, nonce)
	if err != nil {
		return err
	}

	if req.TransferMade {
		b.log.Info("Transfer executed on-chain, marking relayed",
			slog.Int64("event_id", event.ID),
			slog.String("nonce", event.Nonce),
		)
		_ = b.db.UpdateBridgeEventStatus(b.ctx, event.ID, "relayed")
		return nil
	}

	if err := handler(b.ctx, event); err != nil {
		_ = b.db.UpdateBridgeEventStatus(b.ctx, event.ID, "failed")
		return err
	}

	return nil
}
