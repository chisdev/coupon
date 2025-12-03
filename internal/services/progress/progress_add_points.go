package progress

import (
	"context"
	"errors"
	"fmt"
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
			milestone.WithStoreIDs([]string{req.StoreId}),
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

			var (
				goal     int32 = 0
				pass     int32 = 0
				progress int32 = 0
			)
			switch {
			case m.MilestoneType == coupon.MilestoneType_MILESTONE_TYPE_FIXED && m.Threshold != nil:
				if cusPro.PassCount >= 1 {
					continue
				}
				goal = *m.Threshold
				progress = (req.Points + cusPro.Progress) % goal
				if (req.Points+cusPro.Progress)/goal > 0 {
					pass = 1
					progress = goal
				}
			case m.MilestoneType == coupon.MilestoneType_MILESTONE_TYPE_RECURRING && m.Step != nil:
				goal = *m.Step
				pass = (req.Points + cusPro.Progress) / goal
				progress = (req.Points + cusPro.Progress) % goal
			default:
				return errors.New("")
			}

			for range pass {
				for _, reward := range m.Edges.Reward {
					opts := []couponrepo.Option{
						couponrepo.WithUserIDs([]string{req.CustomerId}),
						couponrepo.WithStoreIDs([]string{m.StoreID}),
						couponrepo.WithType(reward.CouponType),
						couponrepo.WithCurrencyID(reward.CurrencyID),
						couponrepo.WithUsageLimit(reward.UsageLimit),
					}
					var expiredAt *time.Time = nil
					if reward.ExpiredDuration != nil {
						t := time.Now().Add(time.Duration(*reward.ExpiredDuration * float64(time.Second)))
						expiredAt = &t
					}
					opts = append(opts, couponrepo.WithExpiredAt(expiredAt))

					if err := p.repository.CouponRepository.CreateTx(ctx, tx, reward.CouponValue,
						opts...,
					); err != nil {
						return err
					}
				}
			}

			cusPro.PassCount += pass
			cusPro.Progress = progress
			cusProList = append(cusProList, cusPro)
		}

		return p.repository.ProgressRepository.UpdateBulkTx(ctx, tx, cusProList)
	})
}

func (p *progress) AddSecretPoints(ctx context.Context, req *coupon.AddSecretPointsRequest) error {
	customerID, _ := p.extractor.GetCustomerID(ctx)
	if customerID == "" {
		return fmt.Errorf("customer id is missing in context")
	}

	return tx.WithTransaction(ctx, p.repository.GetEntClient(), func(ctx context.Context, tx tx.Tx) error {

		eMilestones, _, _, err := p.repository.MileStoneRepository.ListTx(ctx, tx,
			milestone.WithStoreIDs([]string{req.StoreId}),
			milestone.WithReward(true))
		if err != nil && !ent.IsNotFound(err) {
			return err
		}

		cusProList := []*ent.Progress{}

		for _, m := range eMilestones {
			cusPro, err := p.repository.ProgressRepository.GetOrCreateTx(ctx, tx, customerID, m.ID)
			if err != nil {
				return err
			}

			var (
				goal     int32 = 0
				pass     int32 = 0
				progress int32 = 0
			)
			switch {
			case m.MilestoneType == coupon.MilestoneType_MILESTONE_TYPE_FIXED && m.Threshold != nil:
				if cusPro.PassCount >= 1 {
					continue
				}
				goal = *m.Threshold
				progress = (req.Points + cusPro.Progress) % goal
				if (req.Points+cusPro.Progress)/goal > 0 {
					pass = 1
					progress = goal
				}
			case m.MilestoneType == coupon.MilestoneType_MILESTONE_TYPE_RECURRING && m.Step != nil:
				goal = *m.Step
				pass = (req.Points + cusPro.Progress) / goal
				progress = (req.Points + cusPro.Progress) % goal
			default:
				return errors.New("")
			}

			for range pass {
				for _, reward := range m.Edges.Reward {
					opts := []couponrepo.Option{
						couponrepo.WithUserIDs([]string{customerID}),
						couponrepo.WithStoreIDs([]string{m.StoreID}),
						couponrepo.WithType(reward.CouponType),
						couponrepo.WithCurrencyID(reward.CurrencyID),
						couponrepo.WithUsageLimit(reward.UsageLimit),
					}
					var expiredAt *time.Time = nil
					if reward.ExpiredDuration != nil {
						t := time.Now().Add(time.Duration(*reward.ExpiredDuration * float64(time.Second)))
						expiredAt = &t
					}
					opts = append(opts, couponrepo.WithExpiredAt(expiredAt))

					if err := p.repository.CouponRepository.CreateTx(ctx, tx, reward.CouponValue,
						opts...,
					); err != nil {
						return err
					}
				}
			}

			cusPro.PassCount += pass
			cusPro.Progress = progress
			cusProList = append(cusProList, cusPro)
		}

		return p.repository.ProgressRepository.UpdateBulkTx(ctx, tx, cusProList)
	})
}
