package relayer

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"log/slog"

	federationsync "github.com/QwaQ-dev/stacks-base-bridge/contracts/federationsync"
	"github.com/QwaQ-dev/stacks-base-bridge/internal/config"
	"github.com/QwaQ-dev/stacks-base-bridge/internal/models"
	"github.com/QwaQ-dev/stacks-base-bridge/internal/repository/db"
	"github.com/QwaQ-dev/stacks-base-bridge/pkg/logger/sl"
)

const (
	pollInterval = 5 * time.Second
)

type BaseRelayer struct {
	cfg        *config.Config
	log        *slog.Logger
	db         *db.DB
	ctx        context.Context
	ethClient  *ethclient.Client
	federation *federationsync.Federationsync
	authOpts   *bind.TransactOpts
	nodeAddr   common.Address
}

func NewRelayer(cfg *config.Config, log *slog.Logger, database *db.DB, ctx context.Context, ethClient *ethclient.Client) (*BaseRelayer, error) {
	privKey, err := crypto.HexToECDSA(cfg.Base.RelayerPK)
	if err != nil {
		return nil, fmt.Errorf("parse base relayer pk: %w", err)
	}

	chainID, err := ethClient.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("get chain id: %w", err)
	}

	authOpts, err := bind.NewKeyedTransactorWithChainID(privKey, chainID)
	if err != nil {
		return nil, fmt.Errorf("build transactor: %w", err)
	}

	contractAddr := common.HexToAddress(cfg.Base.FederationSyncAddress)
	federation, err := federationsync.NewFederationsync(contractAddr, ethClient)
	if err != nil {
		return nil, fmt.Errorf("bind federation sync: %w", err)
	}

	nodeAddr := crypto.PubkeyToAddress(*privKey.Public().(*ecdsa.PublicKey))

	log.Info("Relayer initialized",
		slog.String("node_id", cfg.Relayer.NodeID),
		slog.String("node_eth_address", nodeAddr.Hex()),
		slog.String("federation_contract", contractAddr.Hex()),
	)

	return &BaseRelayer{
		cfg:        cfg,
		log:        log,
		db:         database,
		ctx:        ctx,
		ethClient:  ethClient,
		federation: federation,
		authOpts:   authOpts,
		nodeAddr:   nodeAddr,
	}, nil
}

func (b *BaseRelayer) Start() {
	go b.runLoop("stacks", b.confirmOnChain)
}

func (b *BaseRelayer) runLoop(sourceChain string, handler func(ctx context.Context, event models.BridgeEvent) error) {
	ticker := time.NewTicker(pollInterval)
	defer ticker.Stop()

	b.log.Info("Relayer loop started", slog.String("source_chain", sourceChain))

	for {
		select {
		case <-b.ctx.Done():
			b.log.Info("Relayer loop stopped", slog.String("source_chain", sourceChain))
			return
		case <-ticker.C:
			if err := b.processPending(sourceChain, handler); err != nil {
				b.log.Error("processPending error", slog.String("chain", sourceChain), sl.Err(err))
			}
		}
	}
}
