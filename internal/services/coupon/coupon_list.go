package coupon

import (
	"context"

	api "github.com/chisdev/coupon/api"
	couponrepo "github.com/chisdev/coupon/internal/repository/coupon"
	"github.com/chisdev/coupon/internal/utiils/convert"
	"github.com/chisdev/coupon/internal/utiils/tx"
	"github.com/chisdev/coupon/pkg/ent"
)

func (c *coupon) ListCouponForCustomer(ctx context.Context, request *api.ListCouponForCustomerRequest) (*api.ListCouponForCustomerResponse, error) {
	customerId, ok := c.extractor.GetCustomerID(ctx)
	if !ok {
		return nil, errMissingCustomerID
	}

	opts := []couponrepo.Option{
		couponrepo.WithStoreIDs(request.StoreId),
		// couponrepo.WithServiceIds(request.ServiceIds),
		couponrepo.WithPaging(request.PageSize, request.PageIndex),
		couponrepo.WithUserIDs([]string{customerId}),
		couponrepo.WithStatus(request.Status),
		couponrepo.WithUsage(true),
	}

	var (
		coupons    []*ent.Coupon
		totalCount int32
		totalPage  int32
		err        error
	)
	if err := tx.WithTransaction(ctx, c.repo.GetEntClient(), func(ctx context.Context, tx tx.Tx) error {
		coupons, totalCount, totalPage, err = c.repo.CouponRepository.List(ctx, tx, opts...)
		return err
	}); err != nil {
		return nil, err
	}

	return &api.ListCouponForCustomerResponse{
		Coupons:    convert.ConvertCoupons(coupons),
		TotalCount: totalCount,
		TotalPages: totalPage,
		Request:    request,
	}, nil

}

func (c *coupon) ListCouponForCms(ctx context.Context, request *api.ListCouponForCmsRequest) (*api.ListCouponForCmsResponse, error) {
	storeId := c.extractor.GetStoreID(ctx)
	if storeId == "" {
		return nil, errMissingStoreID
	}

	opts := []couponrepo.Option{
		couponrepo.WithStoreIDs([]string{storeId}),
		// couponrepo.WithServiceIds(request.ServiceIds),
		couponrepo.WithPaging(request.PageSize, request.PageIndex),
		couponrepo.WithUserIDs(request.CustomerIds),
		couponrepo.WithStatus(request.Status),
		couponrepo.WithUsage(true),
		couponrepo.WithSearchContent(request.SearchContent),
	}

	var (
		coupons    []*ent.Coupon
		totalCount int32
		totalPage  int32
		err        error
	)
	if err := tx.WithTransaction(ctx, c.repo.GetEntClient(), func(ctx context.Context, tx tx.Tx) error {
		coupons, totalCount, totalPage, err = c.repo.CouponRepository.List(ctx, tx, opts...)
		return err
	}); err != nil {
		return nil, err
	}

	return &api.ListCouponForCmsResponse{
		Coupons:    convert.ConvertCoupons(coupons),
		TotalCount: totalCount,
		TotalPages: totalPage,
		Request:    request,
	}, nil

}

func (c *coupon) ListAppliedCoupons(ctx context.Context, request *api.ListAppliedCouponRequest) (*api.ListAppliedCouponResponse, error) {
	coupons, err := c.repo.CouponRepository.ListAppliedCoupons(ctx, request.BookingId)
	if err != nil {
		return nil, err
	}

	return &api.ListAppliedCouponResponse{
		Coupons: convert.ConvertCoupons(coupons),
	}, nil
}
