package milestone

import (
	"context"

	coupon "github.com/chisdev/coupon/api"
	"github.com/chisdev/coupon/pkg/ent"
)

func (m *milestone) Create(ctx context.Context, storeID string, couponValue float64, milestoneType coupon.MilestoneType, opts ...Option) (*ent.Milestone, error) {
	milestoneOpts := &MilestoneOption{}
	for _, opt := range opts {
		opt.Apply(milestoneOpts)
	}

	query := m.ent.Milestone.Create().
		SetStoreID(storeID).
		SetCouponValue(couponValue).
		SetMilestoneType(milestoneType)

	if milestoneOpts.Name != nil {
		query = query.SetName(*milestoneOpts.Name)
	}

	return query.Save(ctx)
}
