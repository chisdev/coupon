package milestone

import (
	"context"

	"github.com/chisdev/coupon/internal/utiils/paging"
	utils "github.com/chisdev/coupon/internal/utiils/sort"
	"github.com/chisdev/coupon/pkg/ent"
	entmilestone "github.com/chisdev/coupon/pkg/ent/milestone"
)

func (m *milestone) List(ctx context.Context, opts ...Option) ([]*ent.Milestone, int32, int32, error) {

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

	totalCount, err := query.Count(ctx)
	if err != nil {
		return nil, 0, 0, err
	}

	totalPage := int32(1)

	if len(mOpts.SortMethods) != 0 {
		sort, err := utils.GetSort(entmilestone.Columns, entmilestone.Table, mOpts.SortMethods)
		if err != nil {
			return nil, 0, 0, err
		}

		query = query.Modify(sort).Clone()
	}

	if mOpts.Limit > 0 {
		query = query.Offset(int(mOpts.PageIndex) * int(mOpts.Limit)).Limit(int(mOpts.Limit))
		totalPage = paging.GetPagingData(int32(totalCount), mOpts.Limit)
	}

	milestones, err := query.WithReward().All(ctx)
	if err != nil {
		return nil, 0, 0, err
	}

	return milestones, int32(totalCount), totalPage, nil
}
