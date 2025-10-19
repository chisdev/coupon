package convert

import (
	coupon "github.com/chisdev/coupon/api"
	"github.com/chisdev/coupon/pkg/ent"
)

func ConvertReward(e *ent.Reward) *coupon.Reward {
	return &coupon.Reward{
		Id:              e.ID,
		CouponValue:     e.CouponValue,
		ExpiredDuration: e.ExpiredDuration,
		CurrencyId:      e.CurrencyID,
		UsageLimit:      e.UsageLimit,
		MilestoneId:     e.MilestoneID,
		CouponType:      e.CouponType,
		ServiceIds:      e.ServiceIds,
	}
}

func ConvertRewards(ents []*ent.Reward) []*coupon.Reward {
	rewards := make([]*coupon.Reward, 0, len(ents))
	for _, e := range ents {
		rewards = append(rewards, ConvertReward(e))
	}
	return rewards
}

func ConvertMilestone(e *ent.Milestone) *coupon.Milestone {
	return &coupon.Milestone{
		Id:        e.ID,
		Type:      e.MilestoneType,
		Threshold: e.Threshold,
		Step:      e.Step,
		StoreId:   e.StoreID,
		Rewards:   ConvertRewards(e.Edges.Reward),
	}
}

func ConvertMilestones(ents []*ent.Milestone) []*coupon.Milestone {
	milestones := make([]*coupon.Milestone, 0, len(ents))
	for _, e := range ents {
		milestones = append(milestones, ConvertMilestone(e))
	}
	return milestones
}
