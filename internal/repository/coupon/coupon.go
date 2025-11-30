package coupon

import (
	"context"

	api "github.com/chisdev/coupon/api"
	"github.com/chisdev/coupon/internal/utiils/tx"
	"github.com/chisdev/coupon/pkg/ent"
)

type Coupon interface {
	Get(ctx context.Context, opts ...Option) (*ent.Coupon, error)
	List(ctx context.Context, tx tx.Tx, opts ...Option) ([]*ent.Coupon, int32, int32, error)
	Create(ctx context.Context, value float64, opts ...Option) (*ent.Coupon, error)
	CreateTx(ctx context.Context, tx tx.Tx, value float64, opts ...Option) error
	Update(ctx context.Context, tx tx.Tx, id uint64, opts ...Option) error
	Delete(ctx context.Context, tx tx.Tx, opts ...Option) error
	Check(ctx context.Context, opts ...Option) ([]*api.CheckCouponsResponse_Result, error)
	ListAppliedCoupons(ctx context.Context, bookingID string) ([]*ent.Coupon, error)
}

type coupon struct {
	ent      *ent.Client
	maxRetry int
	codeLen  int
}

func New(ent *ent.Client) Coupon {
	return &coupon{ent: ent, maxRetry: 10, codeLen: 8}
}
