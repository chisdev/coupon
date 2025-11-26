package coupon

import (
	"context"

	api "github.com/chisdev/coupon/api"
	"github.com/chisdev/coupon/internal/repository"
	"github.com/chisdev/coupon/internal/utiils/extractor"
)

type Coupon interface {
	CheckCoupons(ctx context.Context, req *api.CheckCouponsRequest) (*api.CheckCouponsResponse, error)
	CreateCoupon(ctx context.Context, req *api.CreateCouponRequest) (*api.CreateCouponResponse, error)
	DeleteCoupon(ctx context.Context, req *api.DeleteCouponRequest) error
	ConfirmUsage(ctx context.Context, req *api.ConfirmCouponUsageRequest) error
	Reserve(ctx context.Context, req *api.ReserveCouponRequest) error
	UnReserve(ctx context.Context, req *api.UnReserveCouponRequest) error
	ListCouponForCustomer(ctx context.Context, request *api.ListCouponForCustomerRequest) (*api.ListCouponForCustomerResponse, error)
	ListCouponForCms(ctx context.Context, request *api.ListCouponForCmsRequest) (*api.ListCouponForCmsResponse, error)
	ListAppliedCoupons(ctx context.Context, request *api.ListAppliedCouponRequest) (*api.ListAppliedCouponResponse, error) 
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
