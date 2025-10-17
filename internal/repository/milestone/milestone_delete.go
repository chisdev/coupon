package milestone

import (
	"context"

	"github.com/chisdev/coupon/internal/utiils/tx"
	entmilestone "github.com/chisdev/coupon/pkg/ent/milestone"
)

func (m *milestone) Delete(ctx context.Context, tx tx.Tx, storeID string, ids []uint64) error {
	if _, err := tx.Client().Milestone.Delete().Where(entmilestone.IDIn(ids...)).Exec(ctx); err != nil {
		return err
	}

	return nil
}
