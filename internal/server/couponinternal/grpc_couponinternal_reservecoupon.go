package couponinternal

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"

	coupon "github.com/chisdev/coupon/api"
)

func (s *couponInternalServer) ReserveCoupon(ctx context.Context, request *coupon.ReserveCouponRequest) (*emptypb.Empty, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	if err := s.service.CouponService.Reserve(ctx, request); err != nil {
		s.logger.Error("error while calling Reserve", zap.Error(err))
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
