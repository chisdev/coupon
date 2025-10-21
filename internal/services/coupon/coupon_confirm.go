package coupon

import (
	"context"

	api "github.com/chisdev/coupon/api"
)

func (c *coupon) ConfirmUsage(ctx context.Context, req *api.ConfirmCouponUsageRequest) error {
	return c.repo.CouponBookingRepository.UpdateStatusV2(ctx, req.StoreId, req.BookingId, api.CouponUsedStatus_COUPON_USED_STATUS_RESERVED, api.CouponUsedStatus_COUPON_USED_STATUS_USED)
}
