package currency

import (
	"context"

	coupon "github.com/chisdev/coupon/api"
	currencyrepo "github.com/chisdev/coupon/internal/repository/currency"
	"github.com/chisdev/coupon/internal/utiils/convert"
)

func (c *currency) ListCurrency(ctx context.Context, req *coupon.ListCurrencyRequest) (*coupon.ListCurrencyResponse, error) {
	opts := []currencyrepo.Option{
		currencyrepo.WithPaging(req.PageSize, req.PageIndex),
		currencyrepo.WithSearchContent(req.SearchContent),
		currencyrepo.WithSortMethods(req.SortMethods),
	}

	currencies, totalCount, totalPage, err := c.repo.CurrencyRepository.List(ctx, opts...)
	if err != nil {
		return nil, err
	}

	return &coupon.ListCurrencyResponse{
		Currencies: convert.ConvertCurrencies(currencies),
		TotalPage:  totalPage,
		TotalCount: totalCount,
		Request:    &coupon.ListProgressRequest{},
	}, nil
}
