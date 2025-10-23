package currency

import coupon "github.com/chisdev/coupon/api"

type CurrencyOpts struct {
	Names         []string
	Codes         []string
	SortMethods   []*coupon.SortMethod
	Limit         int32
	PageIndex     int32
	SearchContent string
}

type Option interface {
	Apply(*CurrencyOpts)
}

type funcOption func(*CurrencyOpts)

func (f funcOption) Apply(o *CurrencyOpts) {
	f(o)
}

func WithPaging(limit, pageIndex int32) Option {
	return funcOption(func(o *CurrencyOpts) {
		o.Limit = limit
		o.PageIndex = pageIndex
	})
}

func WithNames(names []string) Option {
	return funcOption(func(o *CurrencyOpts) {
		o.Names = names
	})
}

func WithSortMethods(sortMethods []*coupon.SortMethod) Option {
	return funcOption(func(o *CurrencyOpts) {
		o.SortMethods = sortMethods
	})
}

func WithSearchContent(content string) Option {
	return funcOption(func(o *CurrencyOpts) {
		o.SearchContent = content
	})
}

func WithCodes(codes []string) Option {
	return funcOption(func(co *CurrencyOpts) {
		co.Codes = codes
	})
}
