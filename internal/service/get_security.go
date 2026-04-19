package service

import (
	"context"
	"fmt"

	"go.uber.org/zap"
)

// test
func (s *Service) GetSecurity(ctx context.Context) error {
	ticker := "SBER"
	resp, err := s.client.GetSecurityInfo(ctx, ticker)
	if err != nil {
		s.logger.Error("", zap.Error(err))
		return err
	}
	info, err := s.parser.SecurityDescriptionXML(resp)
	if err != nil {
		s.logger.Error("", zap.Error(err))
		return err
	}
	fmt.Println(info)
	return nil

}
