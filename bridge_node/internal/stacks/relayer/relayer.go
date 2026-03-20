package relayer

import (
	"context"
	"log/slog"
	"time"

	"github.com/QwaQ-dev/stacks-base-bridge/internal/config"
	"github.com/QwaQ-dev/stacks-base-bridge/internal/models"
	"github.com/QwaQ-dev/stacks-base-bridge/internal/repository/db"
	"github.com/QwaQ-dev/stacks-base-bridge/internal/stacks/api"
	"github.com/QwaQ-dev/stacks-base-bridge/pkg/logger/sl"
)

const pollInterval = 5 * time.Second

type StacksRelayer struct {
	cfg  *config.Config
	log  *slog.Logger
	db   *db.DB
	ctx  context.Context
	hiro *api.HiroClient
}

func NewStacksRelayer(
	cfg *config.Config,
	log *slog.Logger,
	database *db.DB,
	ctx context.Context,
) *StacksRelayer {
	return &StacksRelayer{
		cfg:  cfg,
		log:  log,
		db:   database,
		ctx:  ctx,
		hiro: api.NewHiroClient(cfg.Stacks.RPCURL, cfg.Stacks.SignerURL),
	}
}

func (s *StacksRelayer) Start() {
	go s.runLoop("base", s.confirmOnStacks)
}

func (s *StacksRelayer) runLoop(sourceChain string, handler func(ctx context.Context, event models.BridgeEvent) error) {
	ticker := time.NewTicker(pollInterval)
	defer ticker.Stop()

	s.log.Info("Stacks relayer loop started", slog.String("source_chain", sourceChain))

	for {
		select {
		case <-s.ctx.Done():
			s.log.Info("Stacks relayer loop stopped")
			return
		case <-ticker.C:
			if err := s.processPending(sourceChain, handler); err != nil {
				s.log.Error("processPending error", sl.Err(err))
			}
		}
	}
}

func (s *StacksRelayer) processPending(sourceChain string, handler func(ctx context.Context, event models.BridgeEvent) error) error {
	events, err := s.db.GetPendingEvents(s.ctx, sourceChain)
	if err != nil {
		return err
	}

	for _, event := range events {
		if err := handler(s.ctx, event); err != nil {
			s.log.Error("Failed to process event",
				slog.Int64("event_id", event.ID),
				sl.Err(err),
			)
			_ = s.db.UpdateBridgeEventStatus(s.ctx, event.ID, "failed")
		}
	}
	return nil
}
