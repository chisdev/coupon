package couponbooking

import (
	"context"

	coupon "github.com/chisdev/coupon/api"
	"github.com/chisdev/coupon/internal/utiils/tx"
	entcoupon "github.com/chisdev/coupon/pkg/ent/coupon"
	entcouponbooking "github.com/chisdev/coupon/pkg/ent/couponbooking"
)

func (c *couponBooking) UpdateStatus(ctx context.Context, storeId, customerID, couponCode, bookingID string, status coupon.CouponUsedStatus) error {
	return tx.WithTransaction(ctx, c.ent, func(ctx context.Context, tx tx.Tx) error {
		query := tx.Client().CouponBooking.Update().
			Where(entcouponbooking.BookingID(bookingID)).
			Where(entcouponbooking.CustomerID(customerID)).
			Where(entcouponbooking.HasCouponWith(entcoupon.Code(couponCode))).
			Where(entcouponbooking.HasCouponWith(entcoupon.StoreID(storeId))).
			SetStatus(status)

		return query.Exec(ctx)
	})
}

func (c *couponBooking) UpdateStatusV2(ctx context.Context, storeId, bookingID string, status, newStatus coupon.CouponUsedStatus) error {
	return tx.WithTransaction(ctx, c.ent, func(ctx context.Context, tx tx.Tx) error {
		query := tx.Client().CouponBooking.Update().
			Where(entcouponbooking.BookingID(bookingID)).
			Where(entcouponbooking.HasCouponWith(entcoupon.StoreID(storeId))).
			Where(entcouponbooking.Status(status)).
			SetStatus(newStatus)

		return query.Exec(ctx)
	})
}
