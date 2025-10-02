package coupontype

import (
	"context"

	"github.com/chisdev/coupon/internal/utiils/tx"
	"github.com/chisdev/coupon/pkg/ent"
)

func (c *currency) Update(ctx context.Context, tx tx.Tx, id uint64, name string) (*ent.Currency, error) {
	return tx.Client().Currency.UpdateOneID(id).
		SetName(name).
		Save(ctx)
}
