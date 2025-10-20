package couponbooking

import (
	"context"

	coupon "github.com/chisdev/coupon/api"
	"github.com/chisdev/coupon/pkg/ent"
)

type CouponBooking interface {
	Create(ctx context.Context, couponCode string, bookingID uint64) error
	UpdateStatus(ctx context.Context, couponCode string, bookingID uint64, status coupon.CouponUsedStatus) error
	Delete(ctx context.Context, couponCode string, bookingID uint64) error
}

type couponBooking struct {
	ent *ent.Client
}

func New(entClient *ent.Client) CouponBooking {
	return &couponBooking{
		ent: entClient,
	}
}
