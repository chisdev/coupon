package couponbooking

import (
	"context"

	coupon "github.com/chisdev/coupon/api"
	"github.com/chisdev/coupon/internal/utiils/tx"
	entcoupon "github.com/chisdev/coupon/pkg/ent/coupon"
	entcouponbooking "github.com/chisdev/coupon/pkg/ent/couponbooking"
)

func (c *couponBooking) UpdateStatus(ctx context.Context, couponCode string, bookingID uint64, status coupon.CouponUsedStatus) error {
	return tx.WithTransaction(ctx, c.ent, func(ctx context.Context, tx tx.Tx) error {
		query := tx.Client().CouponBooking.Update().Where(entcouponbooking.BookingID(bookingID)).
			Where(entcouponbooking.HasCouponWith(entcoupon.Code(couponCode))).
			SetStatus(status)

		return query.Exec(ctx)
	})
}
