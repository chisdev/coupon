package reward

import "github.com/chisdev/coupon/pkg/ent"

type Reward interface {
}

type reward struct {
	ent *ent.Client
}

func New(ent *ent.Client) Reward {
	return &reward{
		ent: ent,
	}
}
