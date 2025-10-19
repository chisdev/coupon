package reward

import (
	"context"

	"github.com/chisdev/coupon/internal/utiils/tx"
	"github.com/chisdev/coupon/pkg/ent"
)

type Reward interface {
	Create(ctx context.Context, tx tx.Tx, milestoneID uint64, couponValue float64, opts ...Option) (*ent.Reward, error)
	CreateBulk(ctx context.Context, tx tx.Tx, entities []*RewardEntity) ([]*ent.Reward, error)
	Delete(ctx context.Context, tx tx.Tx, opts ...Option) error
	Update(ctx context.Context, tx tx.Tx, id uint64, opts ...Option) error
	List(ctx context.Context, milestoneID uint64, opts ...Option) ([]*ent.Reward, error)
}

type reward struct {
	ent *ent.Client
}

func New(ent *ent.Client) Reward {
	return &reward{
		ent: ent,
	}
}
