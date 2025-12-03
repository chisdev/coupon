package progress

import (
	"context"

	coupon "github.com/chisdev/coupon/api"
	"github.com/chisdev/coupon/internal/repository"
	"github.com/chisdev/coupon/internal/utiils/extractor"
)

type Progress interface {
	AddPoints(ctx context.Context, req *coupon.AddPointRequest) error
	AddSecretPoints(ctx context.Context, req *coupon.AddSecretPointsRequest) error
	List(ctx context.Context, req *coupon.ListProgressRequest) (*coupon.ListProgressResponse, error)
}

type progress struct {
	repository *repository.Repository
	extractor  extractor.Extractor
}

func New(repository *repository.Repository, extractor extractor.Extractor) Progress {
	return &progress{
		repository: repository,
		extractor:  extractor,
	}
}
