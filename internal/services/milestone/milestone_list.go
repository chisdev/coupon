package milestone

import (
	"context"

	coupon "github.com/chisdev/coupon/api"
	milestonerepo "github.com/chisdev/coupon/internal/repository/milestone"
	"github.com/chisdev/coupon/internal/utiils/convert"
)

func (m *milestone) ListMileStone(ctx context.Context, req *coupon.ListMileStoneRequest) (*coupon.ListMileStoneResponse, error) {
	storeId := m.extractor.GetStoreID(ctx)
	if storeId == "" {
		return nil, errStoreIdNotFound
	}

	opts := []milestonerepo.Option{}

	opts = append(opts, milestonerepo.WithStoreIDs([]string{storeId}))
	opts = append(opts, milestonerepo.WithSortMethods(req.SortMethods))
	opts = append(opts, milestonerepo.WithReward(true))

	if req.PageSize != 0 {
		opts = append(opts, milestonerepo.WithPaging(req.PageSize, req.PageIndex))
	}

	milestones, totalCount, totalPage, err := m.repository.MileStoneRepository.List(ctx, opts...)

	return &coupon.ListMileStoneResponse{
		Milestones: convert.ConvertMilestones(milestones),
		TotalCount: totalCount,
		TotalPages: totalPage,
		Request:    req,
	}, err
}
