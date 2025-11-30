package coupon

import (
	"context"

	api "github.com/chisdev/coupon/api"
	"github.com/chisdev/coupon/internal/utiils/tx"
)

func (c *coupon) ConfirmUsage(ctx context.Context, req *api.ConfirmCouponUsageRequest) error {
	return tx.WithTransaction(ctx, c.repo.GetEntClient(), func(ctx context.Context, tx tx.Tx) error {
		if len(req.Codes) > 0 {
			if err := c.repo.CouponBookingRepository.CreateListTx(ctx, tx, req.StoreId, req.BookingId, req.CustomerId, req.Codes, []string{}, api.CouponUsedStatus_COUPON_USED_STATUS_USED); err != nil {
				return err
			}
		}
		return c.repo.CouponBookingRepository.UpdateStatusTx(ctx, tx, req.StoreId, req.BookingId, api.CouponUsedStatus_COUPON_USED_STATUS_RESERVED, api.CouponUsedStatus_COUPON_USED_STATUS_USED)
	})
}
