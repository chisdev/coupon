package progress

import (
	"context"

	"github.com/chisdev/coupon/internal/utiils/tx"
)

func (p *progress) AddPoint(ctx context.Context, tx tx.Tx, customerId, storeId string, points int32) error {
	return nil
}
