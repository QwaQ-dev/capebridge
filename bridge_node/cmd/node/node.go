package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/QwaQ-dev/stacks-base-bridge/internal/base/client"
	baseindexer "github.com/QwaQ-dev/stacks-base-bridge/internal/base/indexer"
	baserelayer "github.com/QwaQ-dev/stacks-base-bridge/internal/base/relayer"
	"github.com/QwaQ-dev/stacks-base-bridge/internal/config"
	postgres "github.com/QwaQ-dev/stacks-base-bridge/internal/repository"
	"github.com/QwaQ-dev/stacks-base-bridge/internal/repository/db"
	stacksindexer "github.com/QwaQ-dev/stacks-base-bridge/internal/stacks/indexer"
	stacksrelayer "github.com/QwaQ-dev/stacks-base-bridge/internal/stacks/relayer"

	"github.com/QwaQ-dev/stacks-base-bridge/pkg/logger"
)

func main() {
	cfg := config.MustLoadConfig()

	log := logger.SetupLogger(cfg.Global.Env)

	log.Info("Config loaded", slog.String("env", cfg.Global.Env))

	sqlDB, err := postgres.InitDatabase(cfg.Database, log)
	if err != nil {
		log.Error("failed to init database", "err", err)
		os.Exit(1)
	}

	database, err := db.NewFromSQL(sqlDB)
	if err != nil {
		log.Error("failed to create db wrapper", "err", err)
		os.Exit(1)
	}
	defer database.Close()

	bc := client.NewClient(cfg, log)
	if bc == nil {
		log.Error("failed to create base client")
		os.Exit(1)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		log.Info("Shutting down...")
		cancel()
	}()

	baseIdx, err := baseindexer.NewBaseIndexer(cfg, log, bc.Eth(), ctx, database)
	if err != nil {
		log.Error("failed to create indexer", "err", err)
		os.Exit(1)
	}

	baseRel, err := baserelayer.NewRelayer(cfg, log, database, ctx, bc.Eth())
	if err != nil {
		log.Error("Failed to create base relayer", "err", err)
		os.Exit(1)
	}

	stacksIdx := stacksindexer.NewStacksIndexer(cfg, log, ctx, database)

	stacksRel := stacksrelayer.NewStacksRelayer(cfg, log, database, ctx)

	log.Info("Starting Bridge", "node_id", cfg.Relayer.NodeID)

	baseIdx.Start()
	stacksIdx.Start()
	baseRel.Start()
	stacksRel.Start()

	<-ctx.Done()
	fmt.Println("Exited")
}
