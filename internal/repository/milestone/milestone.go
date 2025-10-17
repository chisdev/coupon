package milestone

import (
	"context"

	"github.com/chisdev/coupon/internal/utiils/tx"
	"github.com/chisdev/coupon/pkg/ent"
)

type Milestone interface {
	Create(ctx context.Context, tx tx.Tx, storeID string, opts ...Option) (*ent.Milestone, error)
	Delete(ctx context.Context, tx tx.Tx, storeID string, ids []uint64) error
	List(ctx context.Context, opts ...Option) ([]*ent.Milestone, int32, int32, error)
	Update(ctx context.Context, tx tx.Tx, id uint64, opts ...Option) error
}

type milestone struct {
	ent *ent.Client
}

func New(ent *ent.Client) Milestone {
	return &milestone{ent: ent}
}
