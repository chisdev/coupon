package coupon

import (
	"context"

	"github.com/chisdev/coupon/internal/utiils/tx"
)

func (c *coupon) Update(ctx context.Context, tx tx.Tx, id uint64, opts ...Option) error {
	query := tx.Client().Coupon.UpdateOneID(id)

	var couponOpts CouponOpts
	for _, opt := range opts {
		opt.Apply(&couponOpts)
	}

	if couponOpts.Status != 0 {
		query = query.SetStatus(couponOpts.Status)
	}

	return query.Exec(ctx)
}
