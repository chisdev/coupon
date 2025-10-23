package currency

import (
	"github.com/chisdev/coupon/pkg/ent"
)

var (
	initData = []ent.Currency{
		ent.Currency{
			Name: "Vietnamese Dong",
			Code: "VND",
		},
		ent.Currency{
			Name: "Euro",
			Code: "EUR",
		},
		ent.Currency{
			Name: "United States Dollar",
			Code: "USD",
		},
	}
)
