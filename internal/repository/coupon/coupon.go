package coupon

import (
	"context"

	"github.com/chisdev/coupon/pkg/ent"
)

type Coupon interface {
	Get(ctx context.Context, opts ...Option) (*ent.Coupon, error)
	List(ctx context.Context, opts ...Option) ([]*ent.Coupon, error)
	Create(ctx context.Context, coupon *ent.Coupon) (*ent.Coupon, error)
	Update(ctx context.Context, coupon *ent.Coupon, opts ...Option) (*ent.Coupon, error)
	Delete(ctx context.Context, opts ...Option) error
}

type coupon struct {
	ent *ent.Client
}
