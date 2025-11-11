package coupon

import (
	"context"

	api "github.com/chisdev/coupon/api"
	couponrepo "github.com/chisdev/coupon/internal/repository/coupon"
)

func (c *coupon) CheckCoupons(ctx context.Context, req *api.CheckCouponsRequest) (*api.CheckCouponsResponse, error) {
	opts := []couponrepo.Option{
		couponrepo.WithCodes(req.Codes),
		couponrepo.WithServiceIds(req.Services),
	}

	if req.CustomerId != nil {
		opts = append(opts, couponrepo.WithUserIDs([]string{*req.CustomerId}))
	}

	out, err := c.repo.CouponRepository.Check(ctx,
		opts...)

	if err != nil {
		return nil, err
	}

	return &api.CheckCouponsResponse{
		CheckResult: out,
	}, nil
}
