package currency

import (
	"context"

	coupon "github.com/chisdev/coupon/api"
	"github.com/chisdev/coupon/internal/repository"
)

type Currency interface {
	ListCurrency(ctx context.Context, req *coupon.ListCurrencyRequest) (*coupon.ListCurrencyResponse, error)
}

type currency struct {
	repo *repository.Repository
}

func New(repo *repository.Repository) Currency {
	return &currency{
		repo: repo,
	}
}
