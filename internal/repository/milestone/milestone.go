package milestone

import "github.com/chisdev/coupon/pkg/ent"

type Milestone interface {
}

type milestone struct {
	ent *ent.Client
}

func New(ent *ent.Client) Milestone {
	return &milestone{ent: ent}
}
