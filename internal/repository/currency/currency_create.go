package currency

import (
	"context"

	"github.com/chisdev/coupon/pkg/ent"
)

func (c *currency) Create(ctx context.Context, name, code string) (*ent.Currency, error) {
	return c.ent.Currency.Create().
		SetName(name).
		SetCode(code).
		Save(ctx)
}
