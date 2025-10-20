package couponbooking

import (
	"context"

	coupon "github.com/chisdev/coupon/api"
	"github.com/chisdev/coupon/internal/utiils/tx"
	entcoupon "github.com/chisdev/coupon/pkg/ent/coupon"
)

func (c *couponBooking) Create(ctx context.Context, couponCode string, bookingID uint64) error {
	return tx.WithTransaction(ctx, c.ent, func(ctx context.Context, tx tx.Tx) error {
		eCoupon, err := tx.Client().Coupon.Query().Where(entcoupon.Code(couponCode)).WithCouponBookings().Only(ctx)
		if err != nil {
			return err
		}

		if eCoupon.Status != coupon.CouponStatus_COUPON_STATUS_ACTIVE {
			return errCouponNotActive
		}

		if eCoupon.UsageLimit > 0 && int32(len(eCoupon.Edges.CouponBookings)) >= eCoupon.UsageLimit {
			return errCouponUsageLimitExceeded
		}

		_, err = tx.Client().CouponBooking.
			Create().
			SetCouponID(eCoupon.ID).
			SetBookingID(bookingID).
			Save(ctx)
		return err
	})
}
