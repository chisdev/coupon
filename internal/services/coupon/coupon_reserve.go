package coupon

import (
	"context"

	api "github.com/chisdev/coupon/api"
)

func (c *coupon) Reserve(ctx context.Context, req *api.ReserveCouponRequest) error {
	return c.repo.CouponBookingRepository.Create(ctx, req.StoreId, req.Code, req.BookingId, req.CustomerId, []string{})
}
