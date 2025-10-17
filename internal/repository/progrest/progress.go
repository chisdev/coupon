package progress

import "github.com/chisdev/coupon/pkg/ent"

type Progress interface {
}

type progress struct {
	ent *ent.Client
}

func New(ent *ent.Client) Progress {
	return &progress{
		ent: ent,
	}
}
