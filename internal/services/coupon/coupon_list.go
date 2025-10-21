package coupon

import (
	"context"

	api "github.com/chisdev/coupon/api"
	couponrepo "github.com/chisdev/coupon/internal/repository/coupon"
	"github.com/chisdev/coupon/internal/utiils/convert"
)

func (c *coupon) ListCoupons(ctx context.Context, request *api.ListCouponRequest) (*api.ListCouponResponse, error) {
	customerId, ok := c.extractor.GetCustomerID(ctx)
	if !ok {
		return nil, errMissingCustomerID
	}

	opts := []couponrepo.Option{
		couponrepo.WithStoreIDs(request.StoreId),
		couponrepo.WithServiceIds(request.ServiceIds),
		couponrepo.WithPaging(request.PageSize, request.PageIndex),
		couponrepo.WithUserIDs([]string{customerId}),
		couponrepo.WithStatus(request.Status),
		couponrepo.WithUsage(true),
	}

	coupons, totalCount, totalPage, err := c.repo.CouponRepository.List(ctx, opts...)
	if err != nil {
		return nil, err
	}

	return &api.ListCouponResponse{
		Coupons:    convert.ConvertCoupons(coupons),
		TotalCount: totalCount,
		TotalPages: totalPage,
		Request:    request,
	}, nil

}
