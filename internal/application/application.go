package application

import (
	"context"

	"github.com/boldlogic/moex-connector/internal/client"
	"github.com/boldlogic/moex-connector/internal/config"
	moexparser "github.com/boldlogic/moex-connector/internal/moex"
	"github.com/boldlogic/moex-connector/internal/repository"
	"github.com/boldlogic/moex-connector/internal/service"
	"github.com/boldlogic/moex-connector/pkg/transport/httpclient"
	"github.com/boldlogic/moex-connector/pkg/transport/httpclient/clientmetrics"
	"github.com/boldlogic/packages/commonconfig"
	logger "github.com/boldlogic/packages/logger/zaplog"
	"github.com/boldlogic/packages/metrics"
	"go.uber.org/zap"
)

type Application struct {
	cfg    *config.Config
	logger *zap.Logger
	repo   *repository.Repository
	svc    *service.Service
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

func (a *Application) Start(ctx context.Context) error {
	dsn := a.cfg.Db.GetDSN()
	repo, err := repository.NewRepository(ctx, dsn, a.logger)
	if err != nil {
		a.logger.Error("не удалось запустить приложение", zap.Error(err))
		return err
	}
	a.repo = repo

	reg := metrics.New()
	commonClient := httpclient.NewClient(a.cfg.Client)
	moexMetrics := clientmetrics.NewMetrics(reg)
	client := client.NewClient(commonClient, moexMetrics, "moex", a.logger)

	parser := moexparser.NewParser(a.logger)

	a.svc = service.NewService(client, parser, a.logger)

	err = a.svc.GetSecurity(ctx)
	if err != nil {
		a.logger.Error("", zap.Error(err))
		return err
	}

	return nil
}
