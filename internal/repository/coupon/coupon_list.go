package coupon

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqljson"
	api "github.com/chisdev/coupon/api"
	"github.com/chisdev/coupon/internal/utiils/paging"
	utils "github.com/chisdev/coupon/internal/utiils/sort"
	"github.com/chisdev/coupon/internal/utiils/tx"
	"github.com/chisdev/coupon/pkg/ent"
	entcoupon "github.com/chisdev/coupon/pkg/ent/coupon"
	"github.com/chisdev/coupon/pkg/ent/couponbooking"
)

func (c *coupon) List(ctx context.Context, tx tx.Tx, opts ...Option) ([]*ent.Coupon, int32, int32, error) {
	var couponOpts CouponOpts
	for _, opt := range opts {
		opt.Apply(&couponOpts)
	}

	query := tx.Client().Coupon.Query().ForUpdate()

	if len(couponOpts.UserIDs) > 0 {
		query = query.Where(entcoupon.Or(entcoupon.CustomerIDIn(couponOpts.UserIDs...), entcoupon.CustomerIDIsNil()))
	}

	if len(couponOpts.StoreIDs) > 0 {
		query = query.Where(entcoupon.StoreIDIn(couponOpts.StoreIDs...))
	}

	if couponOpts.BookingiD != "" {
		query = query.Where(entcoupon.HasCouponBookingsWith(couponbooking.BookingID(couponOpts.BookingiD), couponbooking.StatusEQ(api.CouponUsedStatus_COUPON_USED_STATUS_RESERVED)))
	}

	if couponOpts.SearchContent != "" {
		query = query.Where(entcoupon.CodeContainsFold(couponOpts.SearchContent))
	}

	for _, serviceID := range couponOpts.ServiceIds {
		query = query.Where(func(s *sql.Selector) {
			s.Where(sqljson.ValueContains(
				s.C(entcoupon.FieldServiceIds),
				serviceID,
			))
		})
	}

	if err := c.UpdateAllCouponStatus(ctx, tx, query); err != nil {
		return nil, 0, 0, err
	}

	if couponOpts.Status != 0 {
		query = query.Where(entcoupon.StatusEQ(couponOpts.Status))
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

	return entCoupons, int32(totalCount), totalPage, nil
}

func (c *coupon) ListAppliedCoupons(ctx context.Context, bookingID string) ([]*ent.Coupon, error) {
	entCoupons, err := c.ent.Coupon.Query().
		Where(
			entcoupon.HasCouponBookingsWith(
				couponbooking.BookingID(bookingID),
			),
		).
		Order(entcoupon.ByCreatedAt(), ent.Asc()).WithCouponBookings(func(cbq *ent.CouponBookingQuery) {
	}).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return entCoupons, nil
}
