package relayer

import (
	"context"
	"fmt"
	"log/slog"
	"strconv"
	"strings"
	"time"

	"github.com/QwaQ-dev/stacks-base-bridge/internal/models"
	"github.com/QwaQ-dev/stacks-base-bridge/pkg/logger/sl"
	"github.com/QwaQ-dev/stacks-base-bridge/pkg/uinttocv"
	"github.com/QwaQ-dev/stacks-base-bridge/pkg/votedelay"
)

// Token used in the transfer call (SIP-010 trait reference)
const bridgeTokenTrait = "ST1PQHQKV0RJXZFY1DGX8MNSNYVE3VGZJSRTPGZGM.usdcx"

func (s *StacksRelayer) confirmOnStacks(ctx context.Context, event models.BridgeEvent) error {
	// 1. Parse nonce and check if this request was already executed on-chain
	nonceUint, _ := strconv.ParseUint(event.Nonce, 10, 64)

	isExecuted, err := s.isAlreadyRelayed(ctx, nonceUint)
	if err != nil {
		return fmt.Errorf("check contract state: %w", err)
	}

	// Early exit if transfer already happened (idempotency guard)
	if isExecuted {
		s.log.Info("Transfer already executed on Stacks, skipping", slog.String("nonce", event.Nonce))
		_ = s.db.UpdateBridgeEventStatus(ctx, event.ID, "relayed")
		return nil
	}

	// 2. Apply staggered delay between nodes (anti-race / coordination mechanism)
	delay := votedelay.VoteDelay(s.cfg.Relayer.NodeID)
	if delay > 0 {
		s.log.Info("Waiting before vote", slog.Duration("delay", delay))
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(delay):
		}
	}

	// 3. Send vote (confirm-request)
	// This does NOT execute transfer, only registers node's approval
	_ = s.db.UpdateBridgeEventStatus(ctx, event.ID, "relaying")

	txID, err := s.broadcastMint(ctx, event)
	if err != nil {
		// These error codes mean "already voted" - not a failure condition
		if strings.Contains(err.Error(), "u201") ||
			strings.Contains(err.Error(), "u202") ||
			strings.Contains(err.Error(), "u203") {

			s.log.Info("Node already confirmed this request", slog.String("nonce", event.Nonce))
		} else {
			// Real failure - rollback status
			s.log.Error("Stacks broadcastMint failed", sl.Err(err))
			_ = s.db.UpdateBridgeEventStatus(ctx, event.ID, "detected")
			return err
		}
	}

	// 4. Wait until vote transaction is confirmed on-chain
	if txID != "" {
		s.log.Info("Wait for vote confirmation", slog.String("txid", txID))

		confirmed, err := s.waitConfirmed(ctx, txID)
		if err != nil || !confirmed {
			return fmt.Errorf("vote transaction failed or timed out: %w", err)
		}
	}

	// 5. Re-check state after vote (another node might have executed transfer already)
	isDoneFinal, _ := s.isAlreadyRelayed(ctx, nonceUint)
	if isDoneFinal {
		s.log.Info("Transfer already finished by someone else", slog.String("nonce", event.Nonce))
		_ = s.db.UpdateBridgeEventStatus(ctx, event.ID, "relayed")
		return nil
	}

	// At this point:
	// - We voted
	// - Transfer is still not executed
	// - Try to execute it (if consensus is reached)

	s.log.Info("Checking for consensus to execute transfer", slog.String("nonce", event.Nonce))

	// FINAL STEP: attempt to execute transfer
	transferTxID, err := s.broadcastTransfer(ctx, nonceUint)
	if err != nil {
		// u206 = ERR-NO-CONSENSUS - expected in 1/3 or 2/3 states
		if strings.Contains(err.Error(), "u206") {
			s.log.Info("Consensus not reached yet, waiting for other nodes", slog.String("nonce", event.Nonce))
			_ = s.db.UpdateBridgeEventStatus(ctx, event.ID, "voted")
			return nil
		}
		return fmt.Errorf("transfer execution failed: %w", err)
	}

	// Wait for final transfer confirmation
	if transferTxID != "" {
		s.log.Info("Transfer transaction broadcasted", slog.String("txid", transferTxID))

		success, _ := s.waitConfirmed(ctx, transferTxID)
		if success {
			s.log.Info("Transfer SUCCESSFUL", slog.String("nonce", event.Nonce))
			_ = s.db.UpdateBridgeEventStatus(ctx, event.ID, "relayed")
		}
	}

	return nil
}

func (s *StacksRelayer) broadcastMint(ctx context.Context, event models.BridgeEvent) (string, error) {
	// Parse and validate input values from event
	nonceUint, err := strconv.ParseUint(event.Nonce, 10, 64)
	if err != nil {
		return "", fmt.Errorf("invalid nonce: %w", err)
	}

	amountUint, err := strconv.ParseUint(event.Amount, 10, 64)
	if err != nil {
		return "", fmt.Errorf("invalid amount: %w", err)
	}

	// Arguments MUST strictly match contract signature:
	// (recipient, amount, request-nonce)
	args := []map[string]any{
		{"type": "principal", "value": event.Receiver},
		{"type": "uint", "value": amountUint},
		{"type": "uint", "value": nonceUint},
	}

	s.log.Info("Broadcasting confirm-request",
		slog.Uint64("nonce", nonceUint),
		slog.String("receiver", event.Receiver),
		slog.Uint64("amount", amountUint))

	return s.hiro.CallSigner(ctx, s.cfg.Stacks.Contract, "confirm-request", args)
}

func (s *StacksRelayer) broadcastTransfer(ctx context.Context, nonce uint64) (string, error) {
	// Contract signature:
	// (transfer (request-nonce uint) (token <sip-010-trait>))
	args := []map[string]any{
		{"type": "uint", "value": nonce},
		// IMPORTANT:
		// This must be a valid SIP-010 trait reference (contract principal)
		{"type": "principal", "value": bridgeTokenTrait},
	}

	s.log.Info("Broadcasting FINAL transfer", slog.Uint64("nonce", nonce))

	return s.hiro.CallSigner(ctx, s.cfg.Stacks.Contract, "transfer", args)
}

func (s *StacksRelayer) waitConfirmed(ctx context.Context, txID string) (bool, error) {
	const (
		interval    = 10 * time.Second
		maxAttempts = 60 // ~10 minutes total wait
	)

	for i := 0; i < maxAttempts; i++ {
		select {
		case <-ctx.Done():
			return false, ctx.Err()
		case <-time.After(interval):
		}

		tx, err := s.hiro.FetchTxInfo(ctx, txID)
		if err != nil {
			// Non-fatal: retry (network / API instability)
			s.log.Warn("FetchTxInfo error", slog.String("txid", txID), slog.String("err", err.Error()))
			continue
		}

		s.log.Debug("Checking tx status", slog.String("txid", txID), slog.String("status", tx.TxStatus))

		switch tx.TxStatus {
		case "success":
			return true, nil

		// Explicit failure states from Stacks VM
		case "abort_by_response", "abort_by_post_condition":
			return false, nil
		}
	}

	// Timeout reached without final state
	return false, fmt.Errorf("timeout waiting for txid %s", txID)
}

func (s *StacksRelayer) isAlreadyRelayed(ctx context.Context, nonce uint64) (bool, error) {
	// Convert uint → Clarity Value (hex-encoded)
	arg := uinttocv.UintToCV(nonce)

	parts := strings.Split(s.cfg.Stacks.Contract, ".")
	if len(parts) < 2 {
		return false, fmt.Errorf("invalid contract in config")
	}

	// Read contract state using read-only call
	res, err := s.hiro.ReadOnlyCall(ctx, parts[0], parts[1], "get-request", []string{arg})
	if err != nil {
		return false, err
	}

	// No request found
	if res.Result == "(none)" || res.Result == "" {
		return false, nil
	}

	// High-level string check (human-readable Clarity result)
	if strings.Contains(res.Result, "transfer-made true") {
		return true, nil
	}

	// Low-level fallback: decode hex result manually
	cleanRes := strings.TrimPrefix(res.Result, "0x")

	// Clarity tuple serialization detail:
	// bool true  → 0x03
	// bool false → 0x04
	// If tuple ends with 03 → transfer-made = true
	if strings.HasSuffix(cleanRes, "03") {
		s.log.Info("Transfer already executed (detected via hex suffix 03)", slog.Uint64("nonce", nonce))
		return true, nil
	}

	// Otherwise - transfer not executed yet
	return false, nil
}
