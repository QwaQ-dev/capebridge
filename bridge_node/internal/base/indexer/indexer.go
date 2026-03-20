package indexer

import (
	"context"
	"fmt"

	"log/slog"

	bridge "github.com/QwaQ-dev/stacks-base-bridge/contracts/bridge"
	"github.com/QwaQ-dev/stacks-base-bridge/internal/config"
	"github.com/QwaQ-dev/stacks-base-bridge/internal/repository/db"
	"github.com/QwaQ-dev/stacks-base-bridge/pkg/logger/sl"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const chunkSize = 49_000

// BaseIndexer - indexer base events (Request-approved)
type BaseIndexer struct {
	cfg            *config.Config
	log            *slog.Logger
	ctx            context.Context
	db             *db.DB
	client         *ethclient.Client
	bridgeInstance *bridge.Bridge
}

// NewBaseIndexer creates new BaseIndexer
func NewBaseIndexer(cfg *config.Config, log *slog.Logger, client *ethclient.Client, ctx context.Context, db *db.DB) (*BaseIndexer, error) {
	bridgeAddr := common.HexToAddress(cfg.Base.Contract)
	bridgeInstance, err := bridge.NewBridge(bridgeAddr, client)
	if err != nil {
		return nil, fmt.Errorf("failed to load bridge contract: %w", err)
	}

	return &BaseIndexer{
		cfg:            cfg,
		log:            log,
		ctx:            ctx,
		db:             db,
		client:         client,
		bridgeInstance: bridgeInstance,
	}, nil
}

func (i *BaseIndexer) Start() {
	go func() {
		if err := i.syncHistorical(); err != nil {
			i.log.Error("base historical sync failed", sl.Err(err))
		}
		i.subscribeLive()
	}()
}
