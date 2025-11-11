package coupon

import (
	"context"

	api "github.com/chisdev/coupon/api"
)

func (c *coupon) Reserve(ctx context.Context, req *api.ReserveCouponRequest) error {
	return c.repo.CouponBookingRepository.CreateList(ctx, req.StoreId, req.BookingId, req.CustomerId, req.Codes, []string{}, api.CouponUsedStatus_COUPON_USED_STATUS_RESERVED)
}
