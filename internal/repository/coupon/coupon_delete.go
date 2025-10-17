package coupon

import (
	"context"

	"github.com/chisdev/coupon/internal/utiils/tx"
	entcoupon "github.com/chisdev/coupon/pkg/ent/coupon"
)

func (c *coupon) Delete(ctx context.Context, tx tx.Tx, opts ...Option) error {
	var couponOpts CouponOpts
	for _, opt := range opts {
		opt.Apply(&couponOpts)
	}

	query := tx.Client().Coupon.Delete()

	if couponOpts.Code != "" {
		query = query.Where(entcoupon.CodeEQ(couponOpts.Code))
	}

	if len(couponOpts.UserIDs) > 0 {
		query = query.Where(entcoupon.CustomerIDIn(couponOpts.UserIDs...))
	}

	if len(couponOpts.StoreIDs) > 0 {
		query = query.Where(entcoupon.StoreIDIn(couponOpts.StoreIDs...))
	}

	if _, err := query.Exec(ctx); err != nil {
		return err
	}

	return nil
}
