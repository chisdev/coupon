package coupon

import (
	"context"
	"time"

	api "github.com/chisdev/coupon/api"
	"github.com/chisdev/coupon/internal/utiils/tx"
	"github.com/chisdev/coupon/pkg/ent"
	"github.com/chisdev/coupon/pkg/ent/couponbooking"
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

func (c *coupon) UpdateAllCouponStatus(ctx context.Context, tx tx.Tx, query *ent.CouponQuery) error {
	entcoupons, err := query.ForUpdate().WithCouponBookings(func(cbq *ent.CouponBookingQuery) {
		cbq.Select(couponbooking.FieldID)
		cbq.Where(couponbooking.StatusEQ(api.CouponUsedStatus_COUPON_USED_STATUS_USED))
	}).All(ctx)
	if err != nil {
		return err
	}

	for _, entcoupon := range entcoupons {
		if entcoupon.ExpireAt != nil && entcoupon.ExpireAt.Before(time.Now()) && entcoupon.Status != api.CouponStatus_COUPON_STATUS_EXPIRED {
			if err := tx.Client().Coupon.UpdateOne(entcoupon).SetStatus(api.CouponStatus_COUPON_STATUS_EXPIRED).Exec(ctx); err != nil {
				return err
			}
		}
		if entcoupon.UsageLimit != nil && len(entcoupon.Edges.CouponBookings) >= int(*entcoupon.UsageLimit) && entcoupon.Status != api.CouponStatus_COUPON_STATUS_USED {
			if err := tx.Client().Coupon.UpdateOne(entcoupon).SetStatus(api.CouponStatus_COUPON_STATUS_USED).Exec(ctx); err != nil {
				return err
			}
		}
	}

	return nil
}
