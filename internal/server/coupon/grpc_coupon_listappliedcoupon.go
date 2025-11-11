package coupon

import (
	"context"

	"github.com/chisdev/coupon/api"
)

func (s *couponServer) ListAppliedCoupon(ctx context.Context, request *coupon.ListAppliedCouponRequest) (*coupon.ListAppliedCouponResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	return &coupon.ListAppliedCouponResponse{}, nil
}
