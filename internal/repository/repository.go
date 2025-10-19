package repository

import (
	"github.com/chisdev/coupon/internal/repository/milestone"
	"github.com/chisdev/coupon/internal/repository/reward"
	"github.com/chisdev/coupon/pkg/ent"
)

type Repository struct {
	ent                 *ent.Client
	MileStoneRepository milestone.Milestone
	RewardRepository    reward.Reward
}

func New(ent *ent.Client) *Repository {
	return &Repository{
		MileStoneRepository: milestone.New(ent),
		RewardRepository:    reward.New(ent),
		ent:                 ent,
	}
}

func (r *Repository) GetEntClient() *ent.Client {
	return r.ent
}
