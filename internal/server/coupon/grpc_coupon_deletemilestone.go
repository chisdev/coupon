package coupon

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"

	coupon "github.com/chisdev/coupon/api"
)

func (s *couponServer) DeleteMileStone(ctx context.Context, request *coupon.DeleteMileStoneRequest) (*emptypb.Empty, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	if err := s.service.MileStoneService.DeleteMilestone(ctx, request); err != nil {
		s.logger.Error("failed to delete milestone", zap.Error(err))
		return nil, err
	}
	
	return &emptypb.Empty{}, nil
}
