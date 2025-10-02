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

	if couponOpts.ReservedCount != 0 {
		query = query.SetReservedCount(couponOpts.ReservedCount)
	}

	if couponOpts.UsageCount != 0 {
		query = query.SetUsedCount(couponOpts.UsageCount)
	}

	return query.Exec(ctx)
}
