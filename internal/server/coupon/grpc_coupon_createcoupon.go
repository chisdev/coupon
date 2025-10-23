package coupon

import (
	"context"

	coupon "github.com/chisdev/coupon/api"
	"go.uber.org/zap"
)

func (s *couponServer) CreateCoupon(ctx context.Context, request *coupon.CreateCouponRequest) (*coupon.CreateCouponResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	resp, err := s.service.CouponService.CreateCoupon(ctx, request)
	if err != nil {
		s.logger.Error("failed to create coupon", zap.Error(err))
		return nil, err
	}
	return resp, nil
}
