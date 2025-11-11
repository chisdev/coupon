package couponinternal

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"

	coupon "github.com/chisdev/coupon/api"
)

func (s *couponInternalServer) UnReserveCoupon(ctx context.Context, request *coupon.UnReserveCouponRequest) (*emptypb.Empty, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	if err := s.service.CouponService.UnReserve(ctx, request); err != nil {
		s.logger.Error("error while calling UnReserve", zap.Error(err))
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
