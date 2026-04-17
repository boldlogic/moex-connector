package application

import (
	"github.com/boldlogic/moex-connector/internal/config"
	"github.com/boldlogic/packages/commonconfig"
	logger "github.com/boldlogic/packages/logger/zaplog"
	"go.uber.org/zap"
)

type Application struct {
	cfg    *config.Config
	logger *zap.Logger
}

const (
	defaultConfigPath = "config.yaml"
)

func New() (*Application, error) {
	configPath := commonconfig.GetConfigPath(defaultConfigPath)
	cfg, err := config.Load(configPath)
	if err != nil {
		return &Application{}, err
	}
	log := logger.New(cfg.Log)
	return &Application{
		cfg:    cfg,
		logger: log,
	}, nil
}
