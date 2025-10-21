package reward

import (
	coupon "github.com/chisdev/coupon/api"
)

type RewardOption struct {
	ExpiredDuration *float64
	CurrencyID      *uint64
	UsageLimit      *int32
	CouponValue     float64
	ServiceIds      []uint64
	CouponType      coupon.CouponType
	IDs             []uint64
	MilestoneID     *uint64
}

type Option interface {
	Apply(*RewardOption)
}

type funcOption func(*RewardOption)

func (f funcOption) Apply(o *RewardOption) {
	f(o)
}

func WithExpireAt(expiredDuration *float64) Option {
	return funcOption(func(mo *RewardOption) {
		mo.ExpiredDuration = expiredDuration
	})
}

func WithCurrencyID(currencyID *uint64) Option {
	return funcOption(func(mo *RewardOption) {
		mo.CurrencyID = currencyID
	})
}
func WithUsageLimit(usageLimit *int32) Option {
	return funcOption(func(mo *RewardOption) {
		mo.UsageLimit = usageLimit
	})
}

func WithCouponValue(couponValue float64) Option {
	return funcOption(func(mo *RewardOption) {
		mo.CouponValue = couponValue
	})
}

func WithServiceIds(serviceIds []uint64) Option {
	return funcOption(func(mo *RewardOption) {
		mo.ServiceIds = serviceIds
	})
}

func WithCouponType(CouponType coupon.CouponType) Option {
	return funcOption(func(mo *RewardOption) {
		mo.CouponType = CouponType
	})
}

func WithIDs(ids []uint64) Option {
	return funcOption(func(mo *RewardOption) {
		mo.IDs = ids
	})
}

func WithMilestoneID(id *uint64) Option {
	return funcOption(func(mo *RewardOption) {
		mo.MilestoneID = id
	})
}
