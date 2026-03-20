package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Global   GlobalConfig   `yaml:"global"`
	Database DatabaseConfig `yaml:"database"`
	Stacks   StacksConfig   `yaml:"stacks"`
	Base     BaseConfig     `yaml:"base"`
	Relayer  RelayerConfig  `yaml:"relayer"`
}

type GlobalConfig struct {
	Env string `yaml:"env"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
	SSLMode  string `yaml:"sslmode"`
}

type StacksConfig struct {
	RPCURL           string `yaml:"rpc_url"`
	Contract         string `yaml:"contract"`
	SignerURL        string `yaml:"signer_url"`
	BridgeTokenTrait string `yaml:"bridge_token_trait"`
}

type BaseConfig struct {
	RPCURL                string `yaml:"rpc_url"`
	RelayerPK             string `yaml:"relayer_pk"`
	Contract              string `yaml:"contract"`
	ChanID                string `yaml:"chan_id"`
	StartBlock            int64  `yaml:"start_block"`
	FederationSyncAddress string `yaml:"federation_sync_address"`
}

type RelayerConfig struct {
	NodeID string `yaml:"node_id"`
}

func envOverride(yamlVal, envKey string) string {
	if v := os.Getenv(envKey); v != "" {
		return v
	}
	return yamlVal
}

func LoadConfig(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("cannot open config file: %w", err)
	}
	defer f.Close()

	var cfg Config
	if err := yaml.NewDecoder(f).Decode(&cfg); err != nil {
		return nil, fmt.Errorf("cannot decode config yaml: %w", err)
	}

	// Base
	cfg.Base.RPCURL = envOverride(cfg.Base.RPCURL, "BASE_RPC_URL")
	cfg.Base.Contract = envOverride(cfg.Base.Contract, "BASE_CONTRACT")
	cfg.Base.RelayerPK = envOverride(cfg.Base.RelayerPK, "BASE_RELAYER_PK")

	// Database
	cfg.Database.Host = envOverride(cfg.Database.Host, "DATABASE_HOST")
	cfg.Database.Port = envOverride(cfg.Database.Port, "DATABASE_PORT")
	cfg.Database.Username = envOverride(cfg.Database.Username, "DATABASE_USER")
	cfg.Database.Password = envOverride(cfg.Database.Password, "DATABASE_PASSWORD")
	cfg.Database.Name = envOverride(cfg.Database.Name, "DATABASE_NAME")
	cfg.Database.SSLMode = envOverride(cfg.Database.SSLMode, "DATABASE_SSLMODE")

	// Stacks
	cfg.Stacks.RPCURL = envOverride(cfg.Stacks.RPCURL, "STACKS_RPC_URL")
	cfg.Stacks.Contract = envOverride(cfg.Stacks.Contract, "STACKS_CONTRACT")
	cfg.Stacks.SignerURL = envOverride(cfg.Stacks.SignerURL, "STACKS_SIGNER_URL")

	// Relayer
	cfg.Relayer.NodeID = envOverride(cfg.Relayer.NodeID, "NODE_ID")

	if cfg.Base.RPCURL == "" || cfg.Base.RelayerPK == "" || cfg.Base.Contract == "" {
		return nil, fmt.Errorf("base: rpc_url, relayer_pk, contract are required")
	}
	if cfg.Stacks.RPCURL == "" || cfg.Stacks.Contract == "" {
		return nil, fmt.Errorf("stacks: rpc_url, contract are required")
	}
	if cfg.Database.Host == "" || cfg.Database.Password == "" {
		return nil, fmt.Errorf("database config is required")
	}

	return &cfg, nil
}

func MustLoadConfig() *Config {
	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		panic("CONFIG_PATH env variable is not set")
	}

	cfg, err := LoadConfig(path)
	if err != nil {
		panic(fmt.Sprintf("config load failed: %v", err))
	}
	return cfg
}
