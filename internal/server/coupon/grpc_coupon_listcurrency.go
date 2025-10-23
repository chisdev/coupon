package coupon

import (
	"context"

	coupon "github.com/chisdev/coupon/api"
	"go.uber.org/zap"
)

func (s *couponServer) ListCurrency(ctx context.Context, request *coupon.ListCurrencyRequest) (*coupon.ListCurrencyResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	resp, err := s.service.CurrencyService.ListCurrency(ctx, request)
	if err != nil {
		s.logger.Error("failed to list currency", zap.Error(err))
		return nil, err
	}

	return resp, nil
}
