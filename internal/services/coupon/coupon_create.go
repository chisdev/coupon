package coupon

import (
	"context"
	"time"

	api "github.com/chisdev/coupon/api"
	couponrepo "github.com/chisdev/coupon/internal/repository/coupon"
	"github.com/chisdev/coupon/internal/utiils/convert"
)

func (c *coupon) CreateCoupon(ctx context.Context, req *api.CreateCouponRequest) (*api.CreateCouponResponse, error) {

	storeId := c.extractor.GetStoreID(ctx)
	if storeId == "" {
		return nil, errMissingStoreID
	}

	opts := []couponrepo.Option{
		couponrepo.WithCurrencyID(req.CurrencyId),
		couponrepo.WithStoreIDs([]string{storeId}),
		// couponrepo.WithServiceIds(req.ServiceIds),
		couponrepo.WithUsageLimit(req.UsageLimit),
	}

	if req.CustomerId != nil {
		opts = append(opts, couponrepo.WithUserIDs([]string{*req.CustomerId}))
	}

	if req.ExpiredDuration != nil {
		expiredAt := time.Now().Add(time.Duration(*req.ExpiredDuration) * time.Second)
		opts = append(opts, couponrepo.WithExpiredAt(&expiredAt))
	}

	ec, err := c.repo.CouponRepository.Create(ctx, req.CouponValue, opts...)
	if err != nil {
		return nil, err
	}

	return &api.CreateCouponResponse{
		Coupon: convert.ConvertCoupon(ec),
	}, nil
}
