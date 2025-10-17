package milestone

import (
	"context"

	coupon "github.com/chisdev/coupon/api"
	"github.com/chisdev/coupon/internal/utiils/tx"
	"github.com/chisdev/coupon/pkg/ent"
)

func (m *milestone) Create(ctx context.Context, tx tx.Tx, storeID string, opts ...Option) (*ent.Milestone, error) {

	milestoneOpts := &MilestoneOption{
		MilestoneType: coupon.MilestoneType_MILESTONE_TYPE_FIXED,
	}
	for _, opt := range opts {
		opt.Apply(milestoneOpts)
	}

	query := m.ent.Milestone.Create().
		SetStoreID(storeID)

	if milestoneOpts.Name != nil {
		query = query.SetName(*milestoneOpts.Name)
	}

	switch milestoneOpts.MilestoneType {
	case coupon.MilestoneType_MILESTONE_TYPE_RECURRING:
		if milestoneOpts.Step == 0 {
			return nil, errMissingStep
		}
		query = query.SetStep(milestoneOpts.Step)
	case coupon.MilestoneType_MILESTONE_TYPE_FIXED:
		if milestoneOpts.Threshold == 0 {
			return nil, errMisstingThreshold
		}
	}
	query = query.SetThreshold(milestoneOpts.Threshold)

	return query.Save(ctx)
}
