package reward

import (
	"context"

	"github.com/chisdev/coupon/internal/utiils/tx"
	entreward "github.com/chisdev/coupon/pkg/ent/reward"
)

func (r *reward) Delete(ctx context.Context, tx tx.Tx, opts ...Option) error {
	query := tx.Client().Reward.Delete()

	rOpt := RewardOption{}
	for _, opt := range opts {
		opt.Apply(&rOpt)
	}

	if len(rOpt.IDs) > 0 {
		query = query.Where(entreward.IDIn(rOpt.IDs...))
	}

	if rOpt.MilestoneID != nil {
		query = query.Where(entreward.MilestoneID(*rOpt.MilestoneID))
	}

	if _, err := query.Exec(ctx); err != nil {
		return err
	}

	return nil
}
