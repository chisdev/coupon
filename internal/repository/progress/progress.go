package progress

import (
	"context"

	"github.com/chisdev/coupon/internal/utiils/tx"
	"github.com/chisdev/coupon/pkg/ent"
)

type Progress interface {
	CreateTx(ctx context.Context, tx tx.Tx, milestoneId uint64, customerId string, passCount, progress int32) (*ent.Progress, error)
	GetOrCreateTx(ctx context.Context, tx tx.Tx, customerId string, milestonerId uint64) (*ent.Progress, error)
	GetTx(ctx context.Context, tx tx.Tx, customerId string, milestonerId uint64) (*ent.Progress, error)
	CreateBulkTx(ctx context.Context, tx tx.Tx, entites []*ent.Progress) ([]*ent.Progress, error)
	UpdateBulkTx(ctx context.Context, tx tx.Tx, pList []*ent.Progress) error
	List(ctx context.Context, opts ...Option) ([]*ent.Progress, int32, int32, error)
	ListTx(ctx context.Context, tx tx.Tx, opts ...Option) ([]*ent.Progress, int32, int32, error)
}

type progress struct {
	ent *ent.Client
}

func New(ent *ent.Client) Progress {
	return &progress{
		ent: ent,
	}
}
