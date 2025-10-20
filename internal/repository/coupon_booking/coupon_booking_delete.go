package couponbooking

import (
	"context"

	coupon "github.com/chisdev/coupon/api"
	"github.com/chisdev/coupon/internal/utiils/tx"
	entcoupon "github.com/chisdev/coupon/pkg/ent/coupon"
	entcouponbooking "github.com/chisdev/coupon/pkg/ent/couponbooking"
)

func (c *couponBooking) Delete(ctx context.Context, storeId, couponCode, bookingID, userID string) error {
	return tx.WithTransaction(ctx, c.ent, func(ctx context.Context, tx tx.Tx) error {
		query := tx.Client().CouponBooking.Delete().
			Where(entcouponbooking.BookingID(bookingID)).
			Where(entcouponbooking.HasCouponWith(entcoupon.Code(couponCode))).
			Where(entcouponbooking.HasCouponWith(entcoupon.StoreID(storeId))).
			Where(entcouponbooking.Status(coupon.CouponUsedStatus_COUPON_USED_STATUS_RESERVED)).
			Where(entcouponbooking.CustomerID(userID))

		_, err := query.Exec(ctx)
		return err
	})
}

func (c *couponBooking) DeleteV2(ctx context.Context, storeId, bookingID string) error {
	return tx.WithTransaction(ctx, c.ent, func(ctx context.Context, tx tx.Tx) error {
		query := tx.Client().CouponBooking.Delete().
			Where(entcouponbooking.BookingID(bookingID)).
			Where(entcouponbooking.HasCouponWith(entcoupon.StoreID(storeId))).
			Where(entcouponbooking.Status(coupon.CouponUsedStatus_COUPON_USED_STATUS_RESERVED))
		_, err := query.Exec(ctx)
		return err
	})
}
