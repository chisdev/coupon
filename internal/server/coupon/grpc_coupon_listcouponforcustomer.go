package coupon

import (
	"context"

	coupon "github.com/chisdev/coupon/api"
	"go.uber.org/zap"
)

func (s *couponServer) ListCouponForCustomer(ctx context.Context, request *coupon.ListCouponForCustomerRequest) (*coupon.ListCouponForCustomerResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	resp, err := s.service.CouponService.ListCouponForCustomer(ctx, request)
	if err != nil {
		s.logger.Error("failed to list coupons for customer", zap.Error(err))
		return nil, err
	}

	return resp, nil

}
