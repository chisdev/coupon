package progress

import (
	"context"

	"github.com/chisdev/coupon/internal/utiils/tx"
	"github.com/chisdev/coupon/pkg/ent"
)

func (p *progress) CreateTx(ctx context.Context, tx tx.Tx, milestoneId uint64, customerId string, passCount, progress int32) (*ent.Progress, error) {
	return tx.Client().Progress.Create().
		SetMilestoneID(milestoneId).
		SetCustomerID(customerId).
		SetPassCount(passCount).
		SetProgress(progress).
		Save(ctx)
}

func (p *progress) CreateBulkTx(ctx context.Context, tx tx.Tx, entites []*ent.Progress) ([]*ent.Progress, error) {
	creates := []*ent.ProgressCreate{}
	for _, entity := range entites {
		creates = append(creates, tx.Client().Progress.Create().
			SetMilestoneID(entity.MilestoneID).
			SetCustomerID(entity.CustomerID).
			SetPassCount(entity.PassCount).
			SetProgress(entity.Progress),
		)
	}
	return tx.Client().Progress.CreateBulk(creates...).Save(ctx)
}
