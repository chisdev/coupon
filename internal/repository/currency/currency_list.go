package coupontype

import (
	"context"

	coupon "github.com/chisdev/coupon/api"
	"github.com/chisdev/coupon/internal/utiils/paging"
	utils "github.com/chisdev/coupon/internal/utiils/sort"
	"github.com/chisdev/coupon/pkg/ent"
	entcurrency "github.com/chisdev/coupon/pkg/ent/currency"
)

func (c *currency) List(ctx context.Context, opts ...Option) ([]*ent.Currency, int32, int32, error) {
	query := c.ent.Currency.Query()

	listOptions := &CurrencyOpts{
		Names:       []string{},
		SortMethods: []*coupon.SortMethod{},
		Limit:       0,
		PageIndex:   0,
	}

	for _, o := range opts {
		o.Apply(listOptions)
	}

	if len(listOptions.Names) > 0 {
		query = query.Where(entcurrency.NameIn(listOptions.Names...))
	}

	if listOptions.SearchContent != "" {
		query = query.Where(entcurrency.NameContainsFold(listOptions.SearchContent))
	}

	totalCount, err := query.Count(ctx)
	if err != nil {
		return nil, 0, 0, err
	}

	totalPage := paging.GetPagingData(int32(totalCount), listOptions.Limit)

	if listOptions.Limit > 0 {
		query = query.Offset(int(listOptions.PageIndex) * int(listOptions.Limit)).Limit(int(listOptions.Limit))
	}

	if len(listOptions.SortMethods) != 0 {
		sort, err := utils.GetSort(entcurrency.Columns, entcurrency.Table, listOptions.SortMethods)
		if err != nil {
			return nil, 0, 0, err
		}

		query = query.Modify(sort).Clone()
	}

	coupontTypes, err := query.All(ctx)
	if err != nil {
		return nil, 0, 0, err
	}

	return coupontTypes, int32(totalCount), totalPage, nil

}
