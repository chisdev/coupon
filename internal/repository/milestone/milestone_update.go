package milestone

import (
	"context"

	coupon "github.com/chisdev/coupon/api"
	"github.com/chisdev/coupon/internal/utiils/tx"
	entmilestone "github.com/chisdev/coupon/pkg/ent/milestone"
)

func (m *milestone) Update(ctx context.Context, tx tx.Tx, id uint64, opts ...Option) error {
	mOpt := MilestoneOption{}

	for _, opt := range opts {
		opt.Apply(&mOpt)
	}

	milestone, err := tx.Client().Milestone.Query().ForUpdate().Where(entmilestone.ID(id)).Only(ctx)
	if err != nil {
		return err
	}

	if mOpt.Name != "" {
		milestone.Name = mOpt.Name
	}

	if mOpt.MilestoneType != 0 && milestone.MilestoneType != mOpt.MilestoneType {
		switch mOpt.MilestoneType {
		case coupon.MilestoneType_MILESTONE_TYPE_RECURRING:
			if mOpt.Step == nil {
				return errMissingStep
			}
			milestone.Step = mOpt.Step
			milestone.Threshold = nil
		case coupon.MilestoneType_MILESTONE_TYPE_FIXED:
			if mOpt.Threshold == nil {
				return errMissingThreshold
			}
			milestone.Threshold = mOpt.Threshold
			milestone.Step = nil
		}
		milestone.MilestoneType = mOpt.MilestoneType
	} else {
		switch mOpt.MilestoneType {
		case coupon.MilestoneType_MILESTONE_TYPE_RECURRING:
			if mOpt.Step != nil && *mOpt.Step != *milestone.Step {
				milestone.Step = mOpt.Step
			}
		case coupon.MilestoneType_MILESTONE_TYPE_FIXED:
			if mOpt.Threshold != nil && *mOpt.Threshold != *milestone.Threshold {
				milestone.Threshold = mOpt.Threshold
			}
		}
	}

	return milestone.Update().Exec(ctx)
}
