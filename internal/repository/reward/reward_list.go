package reward

import (
	"context"

	"github.com/chisdev/coupon/pkg/ent"
	entreward "github.com/chisdev/coupon/pkg/ent/reward"
)

func (r *reward) List(ctx context.Context, milestoneID uint64, opts ...Option) ([]*ent.Reward, error) {
	query := r.ent.Reward.Query().Where(entreward.MilestoneID(milestoneID))

	rOpt := RewardOption{}
	for _, opt := range opts {
		opt.Apply(&rOpt)
	}

	if rOpt.CouponType != 0 {
		query = query.Where(entreward.CouponType(rOpt.CouponType))
	}

	return query.Order(entreward.ByCreatedAt()).All(ctx)
}
