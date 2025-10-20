package couponbooking

import (
	"context"

	coupon "github.com/chisdev/coupon/api"
	"github.com/chisdev/coupon/internal/utiils/checker"
	"github.com/chisdev/coupon/internal/utiils/tx"
	entcoupon "github.com/chisdev/coupon/pkg/ent/coupon"
)

func (c *couponBooking) Create(ctx context.Context, storeId, couponCode, bookingID, userID string, serviceIds []string) error {
	return tx.WithTransaction(ctx, c.ent, func(ctx context.Context, tx tx.Tx) error {
		eCoupon, err := tx.Client().Coupon.Query().
			Where(entcoupon.Code(couponCode)).
			Where(entcoupon.StoreID(storeId)).
			WithCouponBookings().Only(ctx)
		if err != nil {
			return err
		}

		if eCoupon.CustomerID != nil && *eCoupon.CustomerID != userID {
			return errCustomerIdNotMatch
		}

		if len(eCoupon.ServiceIds) > 0 {
			acceptedServiceIds, ok := checker.IsContains(eCoupon.ServiceIds, serviceIds)
			if !ok {
				return errServiceIdsNotAccepted
			}
			serviceIds = acceptedServiceIds
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
			SetServiceIds(serviceIds).
			SetCustomerID(userID).
			Save(ctx)
		return err
	})
}
