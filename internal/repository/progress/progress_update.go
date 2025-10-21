package progress

import (
	"context"

	"github.com/chisdev/coupon/internal/utiils/tx"
	"github.com/chisdev/coupon/pkg/ent"
)

func (p *progress) UpdateBulkTx(ctx context.Context, tx tx.Tx, pList []*ent.Progress) error {
	for _, v := range pList {
		if err := tx.Client().Progress.UpdateOneID(v.ID).
			SetPassCount(v.PassCount).
			SetProgress(v.Progress).
			Exec(ctx); err != nil {
			return err
		}
	}
	return nil
}
