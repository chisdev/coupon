package coupon

import (
	"context"

	couponrepo "github.com/chisdev/coupon/internal/repository/coupon"
	"github.com/chisdev/coupon/internal/utiils/tx"

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

	return tx.WithTransaction(ctx, c.repo.GetEntClient(), func(ctx context.Context, tx tx.Tx) error {
		return c.repo.CouponRepository.Delete(ctx, tx, opts...)
	})
}
