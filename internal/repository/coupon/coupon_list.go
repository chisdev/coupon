package coupon

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqljson"
	api "github.com/chisdev/coupon/api"
	"github.com/chisdev/coupon/internal/utiils/paging"
	utils "github.com/chisdev/coupon/internal/utiils/sort"
	"github.com/chisdev/coupon/pkg/ent"
	entcoupon "github.com/chisdev/coupon/pkg/ent/coupon"
)

func (c *coupon) List(ctx context.Context, opts ...Option) ([]*ent.Coupon, int32, int32, error) {
	var couponOpts CouponOpts
	for _, opt := range opts {
		opt.Apply(&couponOpts)
	}

	query := c.ent.Coupon.Query().ForUpdate()

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

	totalCount, err := query.Count(ctx)
	if err != nil {
		return nil, 0, 0, err
	}

	totalPage := int32(1)

	if len(couponOpts.SortMethods) != 0 {
		sort, err := utils.GetSort(entcoupon.Columns, entcoupon.Table, couponOpts.SortMethods)
		if err != nil {
			return nil, 0, 0, err
		}
		query = query.Modify(sort).Clone()
	}

	if couponOpts.Limit > 0 {
		query = query.Offset(int(couponOpts.PageIndex) * int(couponOpts.Limit)).Limit(int(couponOpts.Limit))
		totalPage = paging.GetPagingData(int32(totalCount), couponOpts.Limit)
	}

	if couponOpts.WithUsage {
		query = query.WithCouponBookings()
	}

	entCoupons, err := query.All(ctx)
	if err != nil {
		return nil, 0, 0, err
	}

	for _, entCoupon := range entCoupons {
		if couponOpts.Status == api.CouponStatus_COUPON_STATUS_ACTIVE {
			if entCoupon.ExpireAt != nil && entCoupon.ExpireAt.Before(time.Now()) {
				entCoupon.Status = api.CouponStatus_COUPON_STATUS_EXPIRED
				if entCoupon, err = entCoupon.Update().Save(ctx); err != nil {
					return nil, 0, 0, err
				}
			}
		}
	}

	return entCoupons, int32(totalCount), totalPage, nil
}
