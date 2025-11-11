package coupon

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqljson"
	api "github.com/chisdev/coupon/api"
	"github.com/chisdev/coupon/pkg/ent"
	entcoupon "github.com/chisdev/coupon/pkg/ent/coupon"
	"github.com/chisdev/coupon/pkg/ent/couponbooking"
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

func (c *coupon) Check(ctx context.Context, opts ...Option) ([]*api.CheckCouponsResponse_Result, error) {
	var couponOpts CouponOpts
	for _, opt := range opts {
		opt.Apply(&couponOpts)
	}

	codes := couponOpts.Codes
	results := make([]*api.CheckCouponsResponse_Result, 0, len(codes))
	checkMap := make(map[string]*api.CheckCouponsResponse_Result, len(codes))

	for _, code := range codes {
		r := &api.CheckCouponsResponse_Result{
			Code:       code,
			Msg:        "Not Found",
			Ok:         false,
			StatusCode: int32(codeNotFound),
		}
		checkMap[code] = r
		results = append(results, r)
	}
	query := c.ent.Coupon.
		Query().
		Where(entcoupon.CodeIn(codes...)).
		Select(entcoupon.FieldCode, entcoupon.FieldCustomerID, entcoupon.FieldExpireAt, entcoupon.FieldUsageLimit, entcoupon.FieldServiceIds).
		WithCouponBookings(func(cbq *ent.CouponBookingQuery) {
			cbq.Select(couponbooking.FieldID)
		})

	// for _, serviceID := range couponOpts.ServiceIds {
	// 	query = query.Where(func(s *sql.Selector) {
	// 		s.Where(sqljson.ValueContains(
	// 			s.C(entcoupon.FieldServiceIds),
	// 			serviceID,
	// 		))
	// 	})
	// }

	coupons, err := query.All(ctx)
	if err != nil {
		return nil, err
	}

	now := time.Now()

	for _, cp := range coupons {
		result := checkMap[cp.Code] // chỉ lookup 1 lần
		if result == nil {
			continue
		}

		if cp.CustomerID != nil && (len(couponOpts.UserIDs) == 0 || couponOpts.UserIDs[0] != *cp.CustomerID) {
			result.Msg = "not owner"
			result.StatusCode = int32(codeInvalid)
			continue
		}

		if cp.ExpireAt != nil && cp.ExpireAt.Before(now) {
			result.Msg = "expired coupon"
			result.StatusCode = int32(codeInvalid)
			continue
		}

		if len(cp.Edges.CouponBookings) >= int(*cp.UsageLimit) {
			result.Msg = "limit"
			result.StatusCode = int32(codeInvalid)
			continue
		}

		result.Msg = "ok"
		result.Ok = true
		result.StatusCode = int32(codeValid)
	}

	return results, nil
}
