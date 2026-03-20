package indexer

import (
	"context"
	"time"

	"log/slog"

	"github.com/QwaQ-dev/stacks-base-bridge/internal/config"
	"github.com/QwaQ-dev/stacks-base-bridge/internal/repository/db"
	"github.com/QwaQ-dev/stacks-base-bridge/internal/stacks/api"
	"github.com/QwaQ-dev/stacks-base-bridge/pkg/logger/sl"
)

const (
	livePollInterval = 20 * time.Second
	stacksBackoff    = 5 * time.Second
)

type StacksIndexer struct {
	cfg          *config.Config
	log          *slog.Logger
	ctx          context.Context
	db           *db.DB
	hiro         *api.HiroClient
	lastMaxNonce int64
}

func NewStacksIndexer(cfg *config.Config, log *slog.Logger, ctx context.Context, database *db.DB) *StacksIndexer {
	lastNonce, err := database.GetLastBlock(ctx, "stacks")
	if err != nil {
		log.Warn("failed to load last max nonce, starting from -1", sl.Err(err))
		lastNonce = -1
	}

	return &StacksIndexer{
		cfg:          cfg,
		log:          log,
		ctx:          ctx,
		db:           database,
		hiro:         api.NewHiroClient(cfg.Global.Env, cfg.Stacks.SignerURL),
		lastMaxNonce: lastNonce,
	}
}

func (i *StacksIndexer) Start() {
	go func() {
		if err := i.syncHistorical(); err != nil {
			i.log.Error("base historical sync failed", sl.Err(err))
		}
		i.log.Info("Stacks starting live polling")
		i.pollLive()
	}()
}
