package coupon

import (
	"context"

	"github.com/chisdev/coupon/internal/utiils/tx"
	"github.com/chisdev/coupon/pkg/ent"
)

type Coupon interface {
	Get(ctx context.Context, opts ...Option) (*ent.Coupon, error)
	List(ctx context.Context, opts ...Option) ([]*ent.Coupon, int32, int32, error)
	Create(ctx context.Context, value float64, opts ...Option) error
	Update(ctx context.Context, tx tx.Tx, id uint64, opts ...Option) error
	Delete(ctx context.Context, tx tx.Tx, opts ...Option) error
}

type coupon struct {
	ent *ent.Client
}

func New(ent *ent.Client) Coupon {
	return &coupon{ent: ent}
}
