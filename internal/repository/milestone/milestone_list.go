package milestone

import (
	"context"

	"github.com/chisdev/coupon/pkg/ent"
	entmilestone "github.com/chisdev/coupon/pkg/ent/milestone"
)

func (m *milestone) List(ctx context.Context, opts ...Option) ([]*ent.Milestone, error) {

	mOpts := &MilestoneOption{}
	for _, opt := range opts {
		opt.Apply(mOpts)
	}

	query := m.ent.Milestone.Query()

	if len(mOpts.StoreIDs) > 0 {
		query = query.Where(entmilestone.StoreIDIn(mOpts.StoreIDs...))
	}


	if mOpts.MilestoneType != 0 {
		query = query.Where(entmilestone.MilestoneType(mOpts.MilestoneType))
	}

	return query.Order(entmilestone.ByCreatedAt()).All(ctx)
}
