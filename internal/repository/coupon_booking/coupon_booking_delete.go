package couponbooking

import (
	"context"

	"github.com/chisdev/coupon/internal/utiils/tx"
	entcoupon "github.com/chisdev/coupon/pkg/ent/coupon"
	entcouponbooking "github.com/chisdev/coupon/pkg/ent/couponbooking"
)

func (c *couponBooking) Delete(ctx context.Context, couponCode string, bookingID uint64) error {
	return tx.WithTransaction(ctx, c.ent, func(ctx context.Context, tx tx.Tx) error {
		query := tx.Client().CouponBooking.Delete().Where(entcouponbooking.BookingID(bookingID)).
			Where(entcouponbooking.HasCouponWith(entcoupon.Code(couponCode)))

		_, err := query.Exec(ctx)
		return err
	})
}
