package coupon

import (
	"context"

	"github.com/chisdev/coupon/api"
)

func (s *couponServer) CheckCoupon(ctx context.Context, request *coupon.CheckCouponRequest) (*coupon.CheckCouponResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	return &coupon.CheckCouponResponse{}, nil
}
