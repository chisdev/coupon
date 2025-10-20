package coupon

import (
	"context"

	api "github.com/chisdev/coupon/api"
)

func (c *coupon) UnReserve(ctx context.Context, req *api.UnReserveCouponRequest) error {
	return c.repo.CouponBookingRepository.DeleteV2(ctx, req.StoreId, req.BookingId)
}
