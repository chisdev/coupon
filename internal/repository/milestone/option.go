package milestone

import "time"

type MilestoneOption struct {
	Name        *string
	ExpireAt    *time.Time
	CurrencyID  *uint64
	UsageLimit  int32
	Step        int32
	Threshold   int32
	CouponValue float64
	ServiceIds  []string
}

type Option interface {
	Apply(*MilestoneOption)
}

type funcOption func(*MilestoneOption)

func (f funcOption) Apply(o *MilestoneOption) {
	f(o)
}

func WithExpireAt(expireAt *time.Time) Option {
	return funcOption(func(mo *MilestoneOption) {
		mo.ExpireAt = expireAt
	})
}

func WithCurrencyID(currencyID *uint64) Option {
	return funcOption(func(mo *MilestoneOption) {
		mo.CurrencyID = currencyID
	})
}
func WithUsageLimit(usageLimit int32) Option {
	return funcOption(func(mo *MilestoneOption) {
		mo.UsageLimit = usageLimit
	})
}

func WithStep(step int32) Option {
	return funcOption(func(mo *MilestoneOption) {
		mo.Step = step
	})
}

func WithThreshold(threshold int32) Option {
	return funcOption(func(mo *MilestoneOption) {
		mo.Threshold = threshold
	})
}

func WithCouponValue(couponValue float64) Option {
	return funcOption(func(mo *MilestoneOption) {
		mo.CouponValue = couponValue
	})
}

func WithServiceIds(serviceIds []string) Option {
	return funcOption(func(mo *MilestoneOption) {
		mo.ServiceIds = serviceIds
	})
}

func WithName(name *string) Option {
	return funcOption(func(mo *MilestoneOption) {
		mo.Name = name
	})
}
