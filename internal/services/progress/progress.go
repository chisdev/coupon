package progress

import (
	"context"

	coupon "github.com/chisdev/coupon/api"
	"github.com/chisdev/coupon/internal/repository"
)

type Progress interface {
	AddPoints(ctx context.Context, req *coupon.AddPointRequest) error
}

type progress struct {
	repository *repository.Repository
}

func New(repository *repository.Repository) Progress {
	return &progress{
		repository: repository,
	}
}
