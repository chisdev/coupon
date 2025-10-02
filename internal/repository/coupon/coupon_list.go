package coupon

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqljson"
	api "github.com/chisdev/coupon/api"
	"github.com/chisdev/coupon/pkg/ent"
	entcoupon "github.com/chisdev/coupon/pkg/ent/coupon"
)

func (c *coupon) List(ctx context.Context, opts ...Option) ([]*ent.Coupon, error) {
	var couponOpts CouponOpts
	for _, opt := range opts {
		opt.Apply(&couponOpts)
	}

	query := c.ent.Coupon.Query().ForUpdate()

	if couponOpts.UserID != "" {
		query = query.Where(entcoupon.CustomerID(couponOpts.UserID))
	}

	if couponOpts.StoreID != "" {
		query = query.Where(entcoupon.StoreID(couponOpts.StoreID))
	}

	if len(couponOpts.UserIDs) > 0 {
		query = query.Where(entcoupon.CustomerIDIn(couponOpts.UserIDs...))
	}

	if len(couponOpts.StoreIDs) > 0 {
		query = query.Where(entcoupon.StoreIDIn(couponOpts.StoreIDs...))
	}

	if couponOpts.Status != 0 {
		query = query.Where(entcoupon.StatusEQ(couponOpts.Status))
	}

	for _, serviceID := range couponOpts.ServiceIds {
		query = query.Where(func(s *sql.Selector) {
			s.Where(sqljson.ValueContains(
				s.C(entcoupon.FieldServiceIds),
				serviceID,
			))
		})
	}

	entCoupons, err := query.Order(entcoupon.ByID()).All(ctx)
	if err != nil {
		return nil, err
	}

	for _, entCoupon := range entCoupons {
		if couponOpts.Status == api.CouponStatus_COUPON_STATUS_ACTIVE {
			if entCoupon.ExpireAt != nil && entCoupon.ExpireAt.Before(time.Now()) {
				entCoupon.Status = api.CouponStatus_COUPON_STATUS_EXPIRED
				if entCoupon, err = entCoupon.Update().Save(ctx); err != nil {
					return nil, err
				}
			}
		}
	}

	return entCoupons, nil
}
