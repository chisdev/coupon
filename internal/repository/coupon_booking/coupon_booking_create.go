package couponbooking

import (
	"context"

	coupon "github.com/chisdev/coupon/api"
	"github.com/chisdev/coupon/internal/utiils/checker"
	"github.com/chisdev/coupon/internal/utiils/tx"
	entcoupon "github.com/chisdev/coupon/pkg/ent/coupon"
)

func (c *couponBooking) Create(ctx context.Context, storeId, couponCode, bookingID string, customerID *string, serviceIds []string) error {
	return tx.WithTransaction(ctx, c.ent, func(ctx context.Context, tx tx.Tx) error {
		eCoupon, err := tx.Client().Coupon.Query().
			Where(entcoupon.Code(couponCode)).
			Where(entcoupon.StoreID(storeId)).
			WithCouponBookings().Only(ctx)
		if err != nil {
			return err
		}

		query := tx.Client().CouponBooking.
			Create().
			SetCouponID(eCoupon.ID).
			SetBookingID(bookingID)

		if eCoupon.CustomerID != nil {
			if customerID == nil || *customerID != *eCoupon.CustomerID {
				return errCustomerIdNotMatch
			}
			query = query.SetCustomerID(*customerID)
		}

		if len(eCoupon.ServiceIds) > 0 {
			acceptedServiceIds, ok := checker.IsContains(eCoupon.ServiceIds, serviceIds)
			if !ok {
				return errServiceIdsNotAccepted
			}
			query = query.SetServiceIds(acceptedServiceIds)
		}

		if eCoupon.Status != coupon.CouponStatus_COUPON_STATUS_ACTIVE {
			return errCouponNotActive
		}

		if eCoupon.UsageLimit != nil && *eCoupon.UsageLimit > 0 && int32(len(eCoupon.Edges.CouponBookings)) >= *eCoupon.UsageLimit {
			return errCouponUsageLimitExceeded
		}

		return query.Exec(ctx)
	})
}
