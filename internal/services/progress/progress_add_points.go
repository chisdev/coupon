package progress

import (
	"context"
	"errors"

	coupon "github.com/chisdev/coupon/api"
	"github.com/chisdev/coupon/internal/repository/milestone"
	"github.com/chisdev/coupon/internal/utiils/tx"
	"github.com/chisdev/coupon/pkg/ent"
)

func (p *progress) AddPoints(ctx context.Context, req *coupon.AddPointRequest) error {
	return tx.WithTransaction(ctx, p.repository.GetEntClient(), func(ctx context.Context, tx tx.Tx) error {

		eMilestones, _, _, err := p.repository.MileStoneRepository.ListTx(ctx, tx, milestone.WithStoreIDs([]string{req.String()}))
		if err != nil && !ent.IsNotFound(err) {
			return err
		}

		cusProList := []*ent.Progress{}

		for _, m := range eMilestones {
			cusPro, err := p.repository.ProgressRepository.GetOrCreateTx(ctx, tx, req.CustomerId, m.ID)
			if err != nil {
				return err
			}

			var goal int32 = 0
			switch {
			case m.MilestoneType == coupon.MilestoneType_MILESTONE_TYPE_FIXED && m.Threshold != nil:
				goal = *m.Threshold
			case m.MilestoneType == coupon.MilestoneType_MILESTONE_TYPE_RECURRING && m.Step != nil:
				goal = *m.Step
			default:
				return errors.New("")
			}

			pass := (req.Points + cusPro.Progress) / goal
			for range pass {
				//Implement flow create coupon here
				println("pass test")
			}

			cusPro.PassCount += pass
			cusPro.Progress = (req.Points + cusPro.Progress) % goal
			cusProList = append(cusProList, cusPro)
		}

		return p.repository.ProgressRepository.UpdateBulkTx(ctx, tx, cusProList)
	})
}
