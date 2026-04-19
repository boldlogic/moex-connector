package service

import (
	"context"

	moexparser "github.com/boldlogic/moex-connector/internal/moex"
	"go.uber.org/zap"
)

type Client interface {
	GetSecurityInfo(ctx context.Context, ticker string) ([]byte, error)
}

type Service struct {
	parser *moexparser.Parser
	client Client
	logger *zap.Logger
}

func NewService(
	client Client,
	parser *moexparser.Parser,
	logger *zap.Logger) *Service {

	s := &Service{
		client: client,
		parser: parser,
		logger: logger,
	}
	return s
}
