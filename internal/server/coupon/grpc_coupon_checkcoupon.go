package coupon

import (
	"context"

	coupon "github.com/chisdev/coupon/api"
	"go.uber.org/zap"
)

func (s *couponServer) CheckCoupon(ctx context.Context, request *coupon.CheckCouponsRequest) (*coupon.CheckCouponsResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	resp, err := s.service.CouponService.CheckCoupons(ctx, request)
	if err != nil {
		s.logger.Error("error while checking coupons", zap.Error(err))
		return nil, err
	}

	return resp, nil
}
