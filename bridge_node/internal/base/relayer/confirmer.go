package relayer

import (
	"context"
	"fmt"
	"time"

	"log/slog"
	"math/big"

	"github.com/QwaQ-dev/stacks-base-bridge/internal/models"
	"github.com/QwaQ-dev/stacks-base-bridge/pkg/parsebigint"
	"github.com/QwaQ-dev/stacks-base-bridge/pkg/votedelay"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

const transferExecutedTopic = "0xf8d62a11d316e69c1460662f86b7046003826db86fdc0d5ffa8c3adaa24f0a58"

func (b *BaseRelayer) confirmOnChain(ctx context.Context, event models.BridgeEvent) error {
	nonce, err := parsebigint.ParseBigInt("nonce", event.Nonce)
	if err != nil {
		return err
	}

	amount, err := parsebigint.ParseBigInt("amount", event.Amount)
	if err != nil {
		return err
	}

	receiver := common.HexToAddress(event.Receiver)
	if receiver == (common.Address{}) {
		return fmt.Errorf("invalid receiver address: %q", event.Receiver)
	}

	// 1. Delay for nodes
	delay := votedelay.VoteDelay(b.cfg.Relayer.NodeID)
	if delay > 0 {
		b.log.Info("Staggered vote delay",
			slog.Int64("event_id", event.ID),
			slog.String("nonce", event.Nonce),
			slog.String("node_id", b.cfg.Relayer.NodeID),
			slog.String("delay", delay.String()),
		)
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(delay):
		}
	}

	b.log.Info("Sending confirmRequest",
		slog.Int64("event_id", event.ID),
		slog.String("node_id", b.cfg.Relayer.NodeID),
		slog.String("nonce", event.Nonce),
	)

	_ = b.db.UpdateBridgeEventStatus(ctx, event.ID, "relaying")

	// 2. Sending Confirm request to smart contracr
	tx, err := b.federation.ConfirmRequest(b.authOpts, receiver, amount, nonce)
	if err != nil {
		b.log.Error("confirmRequest estimate failed (REVERT)",
			slog.Int64("event_id", event.ID),
			slog.String("error", err.Error()),
		)
		_ = b.db.UpdateBridgeEventStatus(ctx, event.ID, "detected")
		return fmt.Errorf("contract revert: %w", err)
	}

	if tx == nil {
		return fmt.Errorf("confirmRequest returned nil transaction")
	}

	b.log.Info("confirmRequest tx sent, waiting for mining",
		slog.String("base_txhash", tx.Hash().Hex()),
		slog.String("nonce", event.Nonce),
	)

	// 3. Mining waiting
	receipt, err := bind.WaitMined(ctx, b.ethClient, tx)
	if err != nil {
		return fmt.Errorf("waiting for tx mining failed: %w", err)
	}

	b.logContractState(ctx, event.ID, nonce, event.Nonce)

	// 4. Transaction status check
	if receipt.Status == 0 {
		b.log.Warn("confirmRequest transaction REVERTED on-chain",
			slog.String("tx_hash", tx.Hash().Hex()),
		)
		_ = b.db.UpdateBridgeEventStatus(ctx, event.ID, "fail")
		return nil
	}

	// 5. Check for consesus
	if transferExecutedInReceipt(receipt) {
		b.log.Info("Quorum reached, transfer executed on-chain!",
			slog.String("nonce", event.Nonce),
		)
		_ = b.db.UpdateBridgeEventStatus(ctx, event.ID, "relayed")
	} else {
		b.log.Info("Vote accepted, waiting for other nodes",
			slog.String("nonce", event.Nonce),
		)
		_ = b.db.UpdateBridgeEventStatus(ctx, event.ID, "voted")
	}

	return nil
}

func transferExecutedInReceipt(receipt *types.Receipt) bool {
	topic := common.HexToHash(transferExecutedTopic)
	for _, log := range receipt.Logs {
		if len(log.Topics) > 0 && log.Topics[0] == topic {
			return true
		}
	}
	return false
}

func (b *BaseRelayer) logContractState(ctx context.Context, eventID int64, nonce *big.Int, nonceStr string) {
	req, err := b.federation.Requests(&bind.CallOpts{Context: ctx}, nonce)
	if err != nil {
		b.log.Error("Failed to read contract state after vote",
			slog.Int64("event_id", eventID),
			slog.String("nonce", nonceStr),
			slog.String("error", err.Error()),
		)
		return
	}
	b.log.Info("Contract state after vote",
		slog.Int64("event_id", eventID),
		slog.String("nonce", nonceStr),
		slog.Bool("transfer_made", req.TransferMade),
		slog.Bool("node1_voted", req.Node1Confirmation),
		slog.Bool("node2_voted", req.Node2Confirmation),
		slog.Bool("node3_voted", req.Node3Confirmation),
		slog.String("node1_recipient", req.Node1Recipient.Hex()),
		slog.String("node2_recipient", req.Node2Recipient.Hex()),
		slog.String("node3_recipient", req.Node3Recipient.Hex()),
		slog.String("node1_amount", req.Node1Amount.String()),
		slog.String("node2_amount", req.Node2Amount.String()),
		slog.String("node3_amount", req.Node3Amount.String()),
	)
}
