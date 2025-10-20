package coupon

import (
	"github.com/chisdev/coupon/internal/repository"
	"github.com/chisdev/coupon/internal/utiils/extractor"
)

type Coupon interface {
}

type coupon struct {
	repo      *repository.Repository
	extractor extractor.Extractor
}

func New(repo *repository.Repository, extractor extractor.Extractor) Coupon {
	return &coupon{
		repo:      repo,
		extractor: extractor,
	}
}
