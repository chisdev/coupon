package coupon

import (
	"context"
	"fmt"
	"time"

	api "github.com/chisdev/coupon/api"
	"github.com/chisdev/coupon/internal/utiils/generator"
	"github.com/chisdev/coupon/internal/utiils/tx"
	"github.com/chisdev/coupon/pkg/ent"
	entcoupon "github.com/chisdev/coupon/pkg/ent/coupon"
)

func (c *coupon) Create(ctx context.Context, value float64, opts ...Option) (*ent.Coupon, error) {
	var couponOpts CouponOpts
	var code string
	var err error

	for _, opt := range opts {
		opt.Apply(&couponOpts)
	}

	for i := 0; i < c.maxRetry; i++ {
		code, err = generator.GenCode(c.codeLen)
		if err != nil {
			return nil, err
		}

		exist, err := c.ent.Coupon.
			Query().
			Where(entcoupon.CodeEQ(code)).
			Exist(ctx)
		if err != nil {
			return nil, err
		}

		if !exist {
			break
		}
	}

	query := c.ent.Coupon.Create().
		SetValue(value).
		SetStatus(api.CouponStatus_COUPON_STATUS_ACTIVE).
		SetCode(code)

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
	var code string
	var err error

	for _, opt := range opts {
		opt.Apply(&couponOpts)
	}

	for i := 0; i < c.maxRetry; i++ {
		code, err = generator.GenCode(c.codeLen)
		if err != nil {
			return err
		}

		exist, err := c.ent.Coupon.
			Query().
			Where(entcoupon.CodeEQ(code)).
			Exist(ctx)
		if err != nil {
			return err
		}

		if !exist {
			break
		}
	}

	if value < 0 {
		return fmt.Errorf("coupon value cannot be negative")
	}

	query := tx.Client().Coupon.Create().
		SetValue(value).
		SetStatus(api.CouponStatus_COUPON_STATUS_ACTIVE).
		SetCode(code)

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
