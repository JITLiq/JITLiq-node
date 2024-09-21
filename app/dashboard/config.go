package dashboard

import (
	"github.com/JITLiq/node/config"
)

type Config struct {
	BaseRPCURL string   `envconfig:"BASE_RPC_URL"`
	ArbRPCURL  string   `envconfig:"ARB_RPC_URL"`
	Fillers    []string `envconfig:"FILLERS"`
	Operators  []string `envconfig:"OPERATORS"`
}

func LoadConfig(configPath string) (*Config, error) {
	cfg := &Config{}
	err := config.LoadEnvConfig(cfg, config.WithEnvPath(configPath))
	return cfg, err
}
