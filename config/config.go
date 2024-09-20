package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Signer  string `envconfig:"SIGNER"`
	ChainID int64  `envconfig:"CHAINID"`
	P2PHost string `envconfig:"P2PHOST"`
}

type Option func(*envConfig)

type envConfig struct {
	envFilePath string
}

func defaultEnvConfig() *envConfig {
	return &envConfig{
		envFilePath: ".env",
	}
}

func WithEnvPath(path string) Option {
	return func(cfg *envConfig) {
		cfg.envFilePath = path
	}
}

func NewConfigFromEnv() (*Config, error) {
	cfg := &Config{}
	err := LoadEnvConfig(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func LoadEnvConfig(cfg any, options ...Option) error {
	defaultCfg := defaultEnvConfig()
	for _, option := range options {
		option(defaultCfg)
	}

	err := godotenv.Overload(defaultCfg.envFilePath)
	if err != nil {
		return fmt.Errorf("failed to read env file, %w", err)
	}

	err = envconfig.Process("", cfg)
	if err != nil {
		return fmt.Errorf("failed to fill config structure.en, %w", err)
	}

	return nil
}
