package progress

import (
	"context"

	coupon "github.com/chisdev/coupon/api"
	progressrepo "github.com/chisdev/coupon/internal/repository/progress"
	"github.com/chisdev/coupon/internal/utiils/convert"
)

func (p *progress) List(ctx context.Context, req *coupon.ListProgressRequest) (*coupon.ListProgressResponse, error) {
	customerId, ok := p.extractor.GetCustomerID(ctx)
	if !ok || customerId == "" {
		return nil, errMissingCustomerId
	}

	opts := []progressrepo.Option{
		progressrepo.WithCustomerIDs([]string{customerId}),
		progressrepo.WithStoreIds(req.StoreIds),
		progressrepo.WithPaging(req.PageSize, req.PageIndex),
		progressrepo.WithSortMethods(req.SortMethods),
	}

	listProgres, totalCount, totalPage, err := p.repository.ProgressRepository.List(ctx, opts...)
	if err != nil {
		return nil, err
	}

	return &coupon.ListProgressResponse{
		ProgressList: convert.ConvertProgresses(listProgres),
		Request:      req,
		TotalPage:    totalPage,
		TotalCount:   totalCount,
	}, nil
}
