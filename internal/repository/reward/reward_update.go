package reward

import (
	"context"

	coupon "github.com/chisdev/coupon/api"
	"github.com/chisdev/coupon/internal/utiils/tx"
	entreward "github.com/chisdev/coupon/pkg/ent/reward"
)

func (r *reward) Update(ctx context.Context, tx tx.Tx, id uint64, opts ...Option) error {
	reward, err := tx.Client().Reward.Query().ForUpdate().Where(entreward.ID(id)).Only(ctx)
	if err != nil {
		return err
	}

	opt := RewardOption{}

	for _, v := range opts {
		v.Apply(&opt)
	}

	if opt.CouponValue > 0 && reward.CouponValue != opt.CouponValue {
		reward.CouponValue = opt.CouponValue
	}

	if opt.UsageLimit > 1 && reward.UsageLimit != opt.UsageLimit {
		reward.UsageLimit = opt.UsageLimit
	}

	if len(opt.ServiceIds) > 0 {
		reward.ServiceIds = opt.ServiceIds
	}

	if opt.ExpiredDuration != nil && *reward.ExpiredDuration > 0 && *reward.ExpiredDuration != *opt.ExpiredDuration {
		reward.ExpiredDuration = opt.ExpiredDuration
	}

	if opt.CouponType != 0 && opt.CouponType != reward.CouponType {
		switch opt.CouponType {
		case coupon.CouponType_COUPON_TYPE_FIXED:
			if opt.CurrencyID == nil {
				return errMissingCurrencyID
			}
			reward.CurrencyID = opt.CurrencyID
			reward.CouponType = opt.CouponType
			return reward.Update().Exec(ctx)
		case coupon.CouponType_COUPON_TYPE_PERCENTAGE:
			reward.CouponType = opt.CouponType
			return reward.Update().ClearCurrency().Exec(ctx)
		default:
			return errInvalidCouponType
		}
	}

	return reward.Update().Exec(ctx)
}
