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

func (c *coupon) Get(ctx context.Context, opts ...Option) (*ent.Coupon, error) {

	var couponOpts CouponOpts
	for _, opt := range opts {
		opt.Apply(&couponOpts)
	}

	query := c.ent.Coupon.Query().ForUpdate()

	if couponOpts.Code != "" {
		query = query.Where(entcoupon.CodeEQ(couponOpts.Code))
	}

	if len(couponOpts.UserIDs) != 0 {
		query = query.Where(entcoupon.CustomerIDIn(couponOpts.UserIDs...))
	}

	if len(couponOpts.StoreIDs) != 0 {
		query = query.Where(entcoupon.StoreIDIn(couponOpts.StoreIDs...))
	}

	for _, serviceID := range couponOpts.ServiceIds {
		query = query.Where(func(s *sql.Selector) {
			s.Where(sqljson.ValueContains(
				s.C(entcoupon.FieldServiceIds),
				serviceID,
			))
		})
	}

	entity, err := query.Only(ctx)
	if err != nil {
		return nil, err
	}

	switch couponOpts.Status {
	case api.CouponStatus_COUPON_STATUS_ACTIVE:
		if entity.ExpireAt != nil && entity.ExpireAt.Before(time.Now()) {
			entity.Status = api.CouponStatus_COUPON_STATUS_EXPIRED
			if entity, err = entity.Update().Save(ctx); err != nil {
				return nil, err
			}
		}
	}

	return entity, nil
}
