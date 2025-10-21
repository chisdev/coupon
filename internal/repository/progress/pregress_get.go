package progress

import (
	"context"

	"github.com/chisdev/coupon/internal/utiils/tx"
	"github.com/chisdev/coupon/pkg/ent"
	entprogress "github.com/chisdev/coupon/pkg/ent/progress"
)

func (p *progress) GetTx(ctx context.Context, tx tx.Tx, customerId string, milestonerId uint64) (*ent.Progress, error) {
	return tx.Client().Progress.Query().
		Where(entprogress.CustomerID(customerId)).
		Where(entprogress.MilestoneID(milestonerId)).
		Only(ctx)
}

func (p *progress) GetOrCreateTx(ctx context.Context, tx tx.Tx, customerId string, milestonerId uint64) (*ent.Progress, error) {
	cp, err := tx.Client().Progress.Query().
		Where(entprogress.CustomerID(customerId)).
		Where(entprogress.MilestoneID(milestonerId)).
		Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	if cp == nil {
		if cp, err = p.CreateTx(ctx, tx, milestonerId, customerId, 0, 0); err != nil {
			return nil, err
		}
	}

	return cp, nil

}
