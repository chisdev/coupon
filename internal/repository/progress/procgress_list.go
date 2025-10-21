package progress

import (
	"context"

	"github.com/chisdev/coupon/internal/utiils/tx"
	"github.com/chisdev/coupon/pkg/ent"
	entprogress "github.com/chisdev/coupon/pkg/ent/progress"
)

func (p *progress) List(ctx context.Context, opts ...Option) ([]*ent.Progress, error) {
	option := ProgrestOption{}
	for _, opt := range opts {
		opt.Apply(&option)
	}

	query := p.ent.Progress.Query()

	if len(option.CustomerIds) != 0 {
		query = query.Where(entprogress.CustomerIDIn(option.CustomerIds...))
	}

	if len(option.MilestoneIds) != 0 {
		query = query.Where(entprogress.MilestoneIDIn(option.MilestoneIds...))
	}

	return query.Order(entprogress.ByID(), ent.Asc()).All(ctx)
}

func (p *progress) ListTx(ctx context.Context, tx tx.Tx, opts ...Option) ([]*ent.Progress, error) {
	option := ProgrestOption{}
	for _, opt := range opts {
		opt.Apply(&option)
	}

	query := tx.Client().Progress.Query()

	if len(option.CustomerIds) != 0 {
		query = query.Where(entprogress.CustomerIDIn(option.CustomerIds...))
	}

	if len(option.MilestoneIds) != 0 {
		query = query.Where(entprogress.MilestoneIDIn(option.MilestoneIds...))
	}

	return query.Order(entprogress.ByID(), ent.Asc()).All(ctx)
}
