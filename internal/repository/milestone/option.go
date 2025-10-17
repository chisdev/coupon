package milestone

import (
	coupon "github.com/chisdev/coupon/api"
)

type MilestoneOption struct {
	Name          *string
	Step          int32
	Threshold     int32
	MilestoneType coupon.MilestoneType
	StoreIDs      []string
}

type Option interface {
	Apply(*MilestoneOption)
}

type funcOption func(*MilestoneOption)

func (f funcOption) Apply(o *MilestoneOption) {
	f(o)
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

func WithName(name *string) Option {
	return funcOption(func(mo *MilestoneOption) {
		mo.Name = name
	})
}

func WithMilestoneType(milestoneType coupon.MilestoneType) Option {
	return funcOption(func(mo *MilestoneOption) {
		mo.MilestoneType = milestoneType
	})
}

func WhereStoreIDs(storeIDs []string) Option {
	return funcOption(func(mo *MilestoneOption) {
		mo.StoreIDs = storeIDs
	})
}
