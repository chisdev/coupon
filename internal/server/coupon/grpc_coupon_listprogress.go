package coupon

import (
	"context"

	coupon "github.com/chisdev/coupon/api"
	"go.uber.org/zap"
)

func (s *couponServer) ListProgress(ctx context.Context, request *coupon.ListProgressRequest) (*coupon.ListProgressResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	resp, err := s.service.ProgressService.List(ctx, request)
	if err != nil {
		s.logger.Error("error while calling list progress", zap.Error(err))
	}

	return resp, nil
}
