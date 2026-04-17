package config

import (
	"github.com/boldlogic/packages/commonconfig"
	logger "github.com/boldlogic/packages/logger/zaplog"
)

type Config struct {
	Log logger.Config `yaml:"log" json:"log"`
}

func Load(configPath string) (*Config, error) {
	cfg, err := commonconfig.DecodeConfigStrict[Config](configPath)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
