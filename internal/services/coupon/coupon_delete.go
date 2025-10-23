package coupon

import (
	"context"

	couponrepo "github.com/chisdev/coupon/internal/repository/coupon"

	api "github.com/chisdev/coupon/api"
)

func (c *coupon) DeleteCoupon(ctx context.Context, req *api.DeleteCouponRequest) error {

	storeId := c.extractor.GetStoreID(ctx)
	if storeId == "" {
		return errMissingStoreID
	}
	opts := []couponrepo.Option{
		couponrepo.WithStoreIDs([]string{storeId}),
		couponrepo.WithIDs(req.Ids),
	}

	err := c.repo.CouponRepository.Delete(ctx, nil, opts...)
	if err != nil {
		return err
	}

	return nil
}
