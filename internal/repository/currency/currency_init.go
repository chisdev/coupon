package currency

import (
	"context"

	"github.com/chisdev/coupon/pkg/ent"
)

func (c *currency) Init(ctx context.Context) {
	creates := []*ent.CurrencyCreate{}
	for _, v := range initData {
		query := c.ent.Currency.Create().
			SetName(v.Name).
			SetCode(v.Code)
		creates = append(creates, query)
	}

	c.ent.Currency.CreateBulk(creates...).OnConflict().DoNothing().Exec(ctx)

}
