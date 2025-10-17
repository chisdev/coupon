package progress

import (
	"context"

	"github.com/chisdev/coupon/internal/utiils/tx"
	"github.com/chisdev/coupon/pkg/ent"
)

func (p *progress) CreateBulk(ctx context.Context, tx tx.Tx, entites []ProgressEntity) ([]*ent.Progress, error) {
	creates := []*ent.ProgressCreate{}
	for _, entity := range entites {
		creates = append(creates, tx.Client().Progress.Create().
			SetMilestoneID(entity.MilestoneId).
			SetCustomerID(entity.CustomerId).
			SetPassCount(entity.PassCount).
			SetProgress(entity.Progress),
		)
	}
	return tx.Client().Progress.CreateBulk(creates...).Save(ctx)
}
