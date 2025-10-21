package reward

import (
	"context"

	coupon "github.com/chisdev/coupon/api"
	"github.com/chisdev/coupon/internal/utiils/tx"
	"github.com/chisdev/coupon/pkg/ent"
)

func (r *reward) Create(ctx context.Context, tx tx.Tx, milestoneID uint64, couponValue float64, opts ...Option) (*ent.Reward, error) {
	if couponValue <= 0 {
		return nil, errInvalidCouponValue
	}

	query := tx.Client().Reward.Create().
		SetMilestoneID(milestoneID).
		SetCouponValue(couponValue)

	rOpt := RewardOption{
		CouponType: coupon.CouponType_COUPON_TYPE_PERCENTAGE,
	}

	for _, opt := range opts {
		opt.Apply(&rOpt)
	}

	if rOpt.ExpiredDuration != nil {
		if *rOpt.ExpiredDuration < 0 {
			return nil, errInvalidExpiredDuration
		}
		query = query.SetExpiredDuration(*rOpt.ExpiredDuration)
	}

	if rOpt.UsageLimit != nil && *rOpt.UsageLimit > 1 {
		query = query.SetUsageLimit(*rOpt.UsageLimit)
	}

	if len(rOpt.ServiceIds) != 0 {
		query = query.SetServiceIds(rOpt.ServiceIds)
	}

	switch rOpt.CouponType {
	case coupon.CouponType_COUPON_TYPE_FIXED:
		if rOpt.CurrencyID == nil {
			return nil, errMissingCurrencyID
		}
		query = query.SetCurrencyID(*rOpt.CurrencyID)
	case coupon.CouponType_COUPON_TYPE_UNKNOWN:
		return nil, errInvalidCouponType
	}
	query = query.SetCouponType(rOpt.CouponType)

	return query.Save(ctx)
}

func (r *reward) CreateBulk(ctx context.Context, tx tx.Tx, entities []*ent.Reward) ([]*ent.Reward, error) {
	creates := []*ent.RewardCreate{}

	for _, entity := range entities {
		if entity.CouponValue <= 0 {
			return nil, errInvalidCouponValue
		}

		query := tx.Client().Reward.Create().
			SetMilestoneID(entity.MilestoneID).
			SetCouponValue(entity.CouponValue).
			SetCouponType(coupon.CouponType_COUPON_TYPE_PERCENTAGE).
			SetServiceIds(entity.ServiceIds)

		if entity.ExpiredDuration != nil {
			if *entity.ExpiredDuration < 0 {
				return nil, errInvalidExpiredDuration
			}
			query = query.SetExpiredDuration(*entity.ExpiredDuration)
		}

		if entity.UsageLimit != nil {
			if *entity.UsageLimit < 1 {
				return nil, errInvaliedUsageLimit
			}
		}

		if entity.CurrencyID != nil {
			query = query.SetCurrencyID(*entity.CurrencyID)
			query = query.SetCouponType(coupon.CouponType_COUPON_TYPE_FIXED)
		}

		creates = append(creates, query)

	}

	return tx.Client().Reward.CreateBulk(creates...).Save(ctx)
}
