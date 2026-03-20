package client

import (
	"log/slog"

	"github.com/QwaQ-dev/stacks-base-bridge/internal/config"
	"github.com/ethereum/go-ethereum/ethclient"
)

type BaseClient struct {
	cfg    *config.Config
	log    *slog.Logger
	client *ethclient.Client
}

// Creating RPC dial with base nodes
func NewClient(cfg *config.Config, log *slog.Logger) *BaseClient {
	client, err := ethclient.Dial(cfg.Base.RPCURL)
	if err != nil {
		log.Error("Error dialing with ethereum")
		return nil
	}

	return &BaseClient{
		cfg:    cfg,
		log:    log,
		client: client,
	}
}

func (c *BaseClient) Eth() *ethclient.Client {
	return c.client
}
