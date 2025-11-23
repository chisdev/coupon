package progress

import (
	"context"

	"github.com/chisdev/coupon/internal/utiils/paging"
	utils "github.com/chisdev/coupon/internal/utiils/sort"
	"github.com/chisdev/coupon/internal/utiils/tx"
	"github.com/chisdev/coupon/pkg/ent"
	"github.com/chisdev/coupon/pkg/ent/milestone"
	entprogress "github.com/chisdev/coupon/pkg/ent/progress"
)

func (p *progress) List(ctx context.Context, opts ...Option) ([]*ent.Progress, int32, int32, error) {
	option := ProgrestOption{}
	for _, opt := range opts {
		opt.Apply(&option)
	}

	query := p.ent.Progress.Query()

	if len(option.StoreIds) != 0 {
		query = query.Where(entprogress.HasMilestoneWith(milestone.StoreIDIn(option.StoreIds...)))
	}

	if len(option.CustomerIds) != 0 {
		query = query.Where(entprogress.CustomerIDIn(option.CustomerIds...))
	}

	if len(option.MilestoneIds) != 0 {
		query = query.Where(entprogress.MilestoneIDIn(option.MilestoneIds...))
	}

	totalCount, err := query.Count(ctx)
	if err != nil {
		return nil, 0, 0, err
	}

	totalPage := int32(1)

	if len(option.SortMethods) != 0 {
		sort, err := utils.GetSort(entprogress.Columns, entprogress.Table, option.SortMethods)
		if err != nil {
			return nil, 0, 0, err
		}
		query = query.Modify(sort).Clone()
	}

	if option.Limit > 0 {
		query = query.Offset(int(option.PageIndex) * int(option.Limit)).Limit(int(option.Limit))
		totalPage = paging.GetPagingData(int32(totalCount), option.Limit)
	}

	progs, err := query.WithMilestone().All(ctx)
	if err != nil {
		return nil, 0, 0, err
	}

	return progs, int32(totalCount), totalPage, nil
}

func (p *progress) ListTx(ctx context.Context, tx tx.Tx, opts ...Option) ([]*ent.Progress, int32, int32, error) {
	option := ProgrestOption{}
	for _, opt := range opts {
		opt.Apply(&option)
	}

	query := tx.Client().Progress.Query()

	if len(option.StoreIds) != 0 {
		query = query.Where(entprogress.HasMilestoneWith(milestone.StoreIDIn(option.StoreIds...)))
	}

	if len(option.CustomerIds) != 0 {
		query = query.Where(entprogress.CustomerIDIn(option.CustomerIds...))
	}

	if len(option.MilestoneIds) != 0 {
		query = query.Where(entprogress.MilestoneIDIn(option.MilestoneIds...))
	}

	totalCount, err := query.Count(ctx)
	if err != nil {
		return nil, 0, 0, err
	}

	totalPage := int32(1)

	if len(option.SortMethods) != 0 {
		sort, err := utils.GetSort(entprogress.Columns, entprogress.Table, option.SortMethods)
		if err != nil {
			return nil, 0, 0, err
		}
		query = query.Modify(sort).Clone()
	}

	if option.Limit > 0 {
		query = query.Offset(int(option.PageIndex) * int(option.Limit)).Limit(int(option.Limit))
		totalPage = paging.GetPagingData(int32(totalCount), option.Limit)
	}

	progs, err := query.All(ctx)
	if err != nil {
		return nil, 0, 0, err
	}

	return progs, int32(totalCount), totalPage, nil
}
