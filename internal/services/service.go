package services

import (
	"github.com/chisdev/coupon/internal/repository"
	"github.com/chisdev/coupon/internal/services/milestone"
	"github.com/chisdev/coupon/internal/utiils/extractor"
)

type Services struct {
	MileStoneService milestone.MileStone
}

func New(repository *repository.Repository) *Services {
	extractor := extractor.New()
	return &Services{
		MileStoneService: milestone.New(repository, extractor),
	}
}
