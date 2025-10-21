package services

import (
	"github.com/chisdev/coupon/internal/repository"
	"github.com/chisdev/coupon/internal/services/coupon"
	"github.com/chisdev/coupon/internal/services/milestone"
	"github.com/chisdev/coupon/internal/services/progress"
	"github.com/chisdev/coupon/internal/utiils/extractor"
)

type Services struct {
	MileStoneService milestone.MileStone
	ProgressService  progress.Progress
	CouponService    coupon.Coupon
}

func New(repository *repository.Repository, extractor extractor.Extractor) *Services {
	return &Services{
		MileStoneService: milestone.New(repository, extractor),
		ProgressService:  progress.New(repository),
		CouponService:    coupon.New(repository, extractor),
	}
}
