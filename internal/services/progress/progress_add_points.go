package progress

import (
	"context"
	"errors"
	"time"

	coupon "github.com/chisdev/coupon/api"
	couponrepo "github.com/chisdev/coupon/internal/repository/coupon"
	"github.com/chisdev/coupon/internal/repository/milestone"
	"github.com/chisdev/coupon/internal/utiils/tx"
	"github.com/chisdev/coupon/pkg/ent"
)

func (p *progress) AddPoints(ctx context.Context, req *coupon.AddPointRequest) error {
	return tx.WithTransaction(ctx, p.repository.GetEntClient(), func(ctx context.Context, tx tx.Tx) error {

		eMilestones, _, _, err := p.repository.MileStoneRepository.ListTx(ctx, tx,
			milestone.WithStoreIDs([]string{req.String()}),
			milestone.WithReward(true))
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
				if cusPro.PassCount >= 1 {
					continue
				}
				goal = *m.Threshold
			case m.MilestoneType == coupon.MilestoneType_MILESTONE_TYPE_RECURRING && m.Step != nil:
				goal = *m.Step
			default:
				return errors.New("")
			}

			pass := (req.Points + cusPro.Progress) / goal
			for range pass {

				for _, reward := range m.Edges.Reward {
					opts := []couponrepo.Option{
						couponrepo.WithUserIDs([]string{req.CustomerId}),
						couponrepo.WithStoreIDs([]string{m.StoreID}),
						couponrepo.WithType(reward.CouponType),
						couponrepo.WithCurrencyID(reward.CurrencyID),
						couponrepo.WithUsageLimit(reward.UsageLimit),
					}

					if reward.ExpiredDuration != nil {
						expiredAt := time.Now().Add(time.Duration(*reward.ExpiredDuration * float64(time.Second)))
						opts = append(opts, couponrepo.WithExpiredAt(&expiredAt))
					}

					if err := p.repository.CouponRepository.CreateTx(ctx, tx, reward.CouponValue,
						opts...,
					); err != nil {
						return err
					}
				}
			}

			cusPro.PassCount += pass
			cusPro.Progress = (req.Points + cusPro.Progress) % goal
			cusProList = append(cusProList, cusPro)
		}

		return p.repository.ProgressRepository.UpdateBulkTx(ctx, tx, cusProList)
	})
}
