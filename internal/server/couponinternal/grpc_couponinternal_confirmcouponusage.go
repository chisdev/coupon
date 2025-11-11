package couponinternal

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"

	coupon "github.com/chisdev/coupon/api"
)

func (s *couponInternalServer) ConfirmCouponUsage(ctx context.Context, request *coupon.ConfirmCouponUsageRequest) (*emptypb.Empty, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	if err := s.service.CouponService.ConfirmUsage(ctx, request); err != nil {
		s.logger.Error("error while calling ConfirmUsage", zap.Error(err))
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
