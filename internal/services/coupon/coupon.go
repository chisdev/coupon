package coupon

import (
	"context"

	api "github.com/chisdev/coupon/api"
	"github.com/chisdev/coupon/internal/repository"
	"github.com/chisdev/coupon/internal/utiils/extractor"
)

type Coupon interface {
	ConfirmUsage(ctx context.Context, req *api.ConfirmCouponUsageRequest) error
	Reserve(ctx context.Context, req *api.ReserveCouponRequest) error
	UnReserve(ctx context.Context, req *api.UnReserveCouponRequest) error
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
