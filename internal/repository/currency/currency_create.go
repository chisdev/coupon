package coupontype

import (
	"context"

	"github.com/chisdev/coupon/pkg/ent"
)

func (c *currency) Create(ctx context.Context, name string) (*ent.Currency, error) {
	return c.ent.Currency.Create().
		SetName(name).
		Save(ctx)
}
