package coupon

import (
	"context"

	coupon "github.com/chisdev/coupon/api"
	"go.uber.org/zap"
)

func (s *couponServer) ListCoupon(ctx context.Context, request *coupon.ListCouponRequest) (*coupon.ListCouponResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	resp, err := s.service.CouponService.ListCoupons(ctx, request)
	if err != nil {
		s.logger.Error("failed to list coupons", zap.Error(err))
		return nil, err
	}

	return resp, nil
}
