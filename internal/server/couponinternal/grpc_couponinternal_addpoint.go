package couponinternal

import (
	"context"
	"fmt"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"

	coupon "github.com/chisdev/coupon/api"
)

func (s *couponInternalServer) AddPoint(ctx context.Context, request *coupon.AddPointRequest) (*emptypb.Empty, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	if request.Points < 0 {
		s.logger.Error("points to add cannot be negative", zap.Int32("points", request.Points))
		return nil, fmt.Errorf("points to add cannot be negative")
	}

	if err := s.service.ProgressService.AddPoints(ctx, request); err != nil {
		s.logger.Error("error while calling AddPoints", zap.Error(err))
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
