package couponcms

import (
	"context"

	coupon "github.com/chisdev/coupon/api"
	"go.uber.org/zap"
)

func (s *couponCmsServer) ListMileStone(ctx context.Context, request *coupon.ListMileStoneRequest) (*coupon.ListMileStoneResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	resp, err := s.service.MileStoneService.ListMileStone(ctx, request)
	if err != nil {
		s.logger.Error("failed to list milestones", zap.Error(err))
		return nil, err
	}
	return resp, nil
}
