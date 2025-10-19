package milestone

import (
	"context"

	coupon "github.com/chisdev/coupon/api"
	"github.com/chisdev/coupon/internal/utiils/tx"
)

func (m *milestone) DeleteMilestone(ctx context.Context, req *coupon.DeleteMileStoneRequest) error {
	storeId := m.extractor.GetStoreID(ctx)
	if storeId == "" {
		return errStoreIdNotFound
	}

	return tx.WithTransaction(ctx, m.repository.GetEntClient(), func(ctx context.Context, tx tx.Tx) error {
		return m.repository.MileStoneRepository.Delete(ctx, tx, storeId, req.GetIds())
	})
}
