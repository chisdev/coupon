package coupon

import (
	"context"

	coupon "github.com/chisdev/coupon/api"
	"go.uber.org/zap"
)

func (s *couponServer) CreateMileStone(ctx context.Context, request *coupon.CreateMileStoneRequest) (*coupon.CreateMileStoneResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	rsp, err := s.service.MileStoneService.CreateMilestone(ctx, request)
	if err != nil {
		s.logger.Error("failed to create milestone", zap.Error(err))
		return nil, err
	}
	return rsp, nil
}
