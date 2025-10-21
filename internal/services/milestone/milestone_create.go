package milestone

import (
	"context"

	coupon "github.com/chisdev/coupon/api"
	milestonerepo "github.com/chisdev/coupon/internal/repository/milestone"
	"github.com/chisdev/coupon/internal/utiils/convert"
	"github.com/chisdev/coupon/internal/utiils/tx"
	"github.com/chisdev/coupon/pkg/ent"
)

func (m *milestone) CreateMilestone(ctx context.Context, req *coupon.CreateMileStoneRequest) (*coupon.CreateMileStoneResponse, error) {
	storeId := m.extractor.GetStoreID(ctx)
	if storeId == "" {
		return nil, errStoreIdNotFound
	}

	var (
		milestone *ent.Milestone
		err       error
	)

	if err := tx.WithTransaction(ctx, m.repository.GetEntClient(), func(ctx context.Context, tx tx.Tx) error {
		options := []milestonerepo.Option{}

		if req.Name == "" {
			return errMissingNames
		}

		if req.Threshold == nil && req.Step == nil {
			return errInvalidMilestoneConfig
		}

		if req.Threshold != nil && req.Step != nil {
			return errInvalidMilestoneConfig
		}

		options = append(options, milestonerepo.WithName(req.GetName()))

		if req.Step != nil {
			options = append(options, milestonerepo.WithStep(req.Step))
			options = append(options, milestonerepo.WithMilestoneType(coupon.MilestoneType_MILESTONE_TYPE_RECURRING))
		}

		if req.Threshold != nil {
			options = append(options, milestonerepo.WithThreshold(req.Threshold))
			options = append(options, milestonerepo.WithMilestoneType(coupon.MilestoneType_MILESTONE_TYPE_FIXED))
		}

		milestone, err = m.repository.MileStoneRepository.Create(ctx, tx, storeId, options...)
		if err != nil {
			return err
		}

		rewardEntities := []*ent.Reward{}
		for _, reward := range req.Rewards {
			rewardEntities = append(rewardEntities, &ent.Reward{
				MilestoneID:     milestone.ID,
				ExpiredDuration: reward.ExpiredDuration,
				CurrencyID:      reward.CurrencyId,
				UsageLimit:      reward.UsageLimit,
				CouponValue:     reward.CouponValue,
				ServiceIds:      reward.ServiceIds,
			})
		}

		if milestone.Edges.Reward, err = m.repository.RewardRepository.CreateBulk(ctx, tx, rewardEntities); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &coupon.CreateMileStoneResponse{
		Milestone: convert.ConvertMilestone(milestone),
	}, nil
}
