package milestone

import (
	"context"

	coupon "github.com/chisdev/coupon/api"
	"github.com/chisdev/coupon/internal/repository"
	"github.com/chisdev/coupon/internal/utiils/extractor"
)

type MileStone interface {
	CreateMilestone(ctx context.Context, req *coupon.CreateMileStoneRequest) (*coupon.CreateMileStoneResponse, error)
	ListMileStone(ctx context.Context, req *coupon.ListMileStoneRequest) (*coupon.ListMileStoneResponse, error)
	DeleteMilestone(ctx context.Context, req *coupon.DeleteMileStoneRequest) error
}

type milestone struct {
	repository *repository.Repository
	extractor  extractor.Extractor
}

func New(repository *repository.Repository, extractor extractor.Extractor) MileStone {
	return &milestone{
		repository: repository,
		extractor:  extractor,
	}
}
