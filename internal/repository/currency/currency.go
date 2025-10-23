package currency

import (
	"context"

	"github.com/chisdev/coupon/internal/utiils/tx"
	"github.com/chisdev/coupon/pkg/ent"
)

type Currency interface {
	Create(ctx context.Context, name, code string) (*ent.Currency, error)
	Delete(ctx context.Context, tx tx.Tx, ids []uint64) error
	List(ctx context.Context, opts ...Option) ([]*ent.Currency, int32, int32, error)
}

type currency struct {
	ent *ent.Client
}

func New(ent *ent.Client) Currency {
	c := &currency{ent: ent}
	c.Init(context.Background())
	return c
}
