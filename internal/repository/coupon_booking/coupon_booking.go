package couponbooking

import (
	"context"

	coupon "github.com/chisdev/coupon/api"
	"github.com/chisdev/coupon/pkg/ent"
)

type CouponBooking interface {
	Create(ctx context.Context, storeId, couponCode, bookingID string, customerID *string, serviceIds []string) error
	UpdateStatus(ctx context.Context, storeId, customerID, couponCode, bookingID string, status coupon.CouponUsedStatus) error
	UpdateStatusV2(ctx context.Context, storeId, bookingID string, status, newStatus coupon.CouponUsedStatus) error
	Delete(ctx context.Context, storeId, couponCode, bookingID, userID string) error
	DeleteV2(ctx context.Context, storeId, bookingID string) error
}

type couponBooking struct {
	ent *ent.Client
}

func New(entClient *ent.Client) CouponBooking {
	return &couponBooking{
		ent: entClient,
	}
}
