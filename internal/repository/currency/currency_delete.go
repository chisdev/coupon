package currency

import (
	"context"

	"github.com/chisdev/coupon/internal/utiils/tx"
	entcurrency "github.com/chisdev/coupon/pkg/ent/currency"
)

func (c *currency) Delete(ctx context.Context, tx tx.Tx, ids []uint64) error {
	_, err := tx.Client().Currency.Delete().
		Where(entcurrency.IDIn(ids...)).Exec(ctx)
	return err
}
