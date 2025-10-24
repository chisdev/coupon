package coupon

import (
	"context"
	"time"

	api "github.com/chisdev/coupon/api"
	"github.com/chisdev/coupon/internal/utiils/tx"
	"github.com/chisdev/coupon/pkg/ent"
	"github.com/google/uuid"
)

func (c *coupon) Create(ctx context.Context, value float64, opts ...Option) (*ent.Coupon, error) {
	var couponOpts CouponOpts
	for _, opt := range opts {
		opt.Apply(&couponOpts)
	}

	query := c.ent.Coupon.Create().
		SetValue(value).
		SetStatus(api.CouponStatus_COUPON_STATUS_ACTIVE).
		SetCode(uuid.NewString())

	if len(couponOpts.UserIDs) != 0 {
		query = query.SetCustomerID(couponOpts.UserIDs[0])
	}

	if len(couponOpts.StoreIDs) != 0 {
		query = query.SetStoreID(couponOpts.StoreIDs[0])
	}

	if couponOpts.ExpiredAt != nil {
		if couponOpts.ExpiredAt.IsZero() || couponOpts.ExpiredAt.Before(time.Now()) {
			return nil, errInvalidExpiredAt
		}
		query = query.SetExpireAt(*couponOpts.ExpiredAt)
	}

	if couponOpts.UsageLimit != nil {
		query = query.SetUsageLimit(*couponOpts.UsageLimit)
	}

	if len(couponOpts.ServiceIds) > 0 {
		query = query.SetServiceIds(couponOpts.ServiceIds)
	}

	switch {
	case couponOpts.CurrencyID != nil:
		query = query.SetType(api.CouponType_COUPON_TYPE_FIXED)
	default:
		query = query.SetType(api.CouponType_COUPON_TYPE_PERCENTAGE)
	}

	return query.Save(ctx)
}

func (c *coupon) CreateTx(ctx context.Context, tx tx.Tx, value float64, opts ...Option) error {
	var couponOpts CouponOpts
	for _, opt := range opts {
		opt.Apply(&couponOpts)
	}

	query := tx.Client().Coupon.Create().
		SetValue(value).
		SetStatus(api.CouponStatus_COUPON_STATUS_ACTIVE).
		SetCode(uuid.NewString())

	if len(couponOpts.UserIDs) != 0 {
		query = query.SetCustomerID(couponOpts.UserIDs[0])
	}

	if len(couponOpts.StoreIDs) != 0 {
		query = query.SetStoreID(couponOpts.StoreIDs[0])
	}

	if couponOpts.ExpiredAt != nil {
		if couponOpts.ExpiredAt.IsZero() || couponOpts.ExpiredAt.Before(time.Now()) {
			return errInvalidExpiredAt
		}
		query = query.SetExpireAt(*couponOpts.ExpiredAt)
	}

	if couponOpts.UsageLimit != nil {
		query = query.SetUsageLimit(*couponOpts.UsageLimit)
	}

	if len(couponOpts.ServiceIds) > 0 {
		query = query.SetServiceIds(couponOpts.ServiceIds)
	}

	switch couponOpts.Type {
	case api.CouponType_COUPON_TYPE_PERCENTAGE:
		if couponOpts.CurrencyID != nil {
			return errConfigCouponTypeWithCurrencyID
		}
		query = query.SetType(couponOpts.Type)
	case api.CouponType_COUPON_TYPE_FIXED:
		if couponOpts.CurrencyID == nil {
			return errMissingCurrencyID
		}
		query = query.SetType(couponOpts.Type)
		query = query.SetCurrencyID(*couponOpts.CurrencyID)
	default:
		return errInvalidCouponType
	}

	return query.Exec(ctx)
}
