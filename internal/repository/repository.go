package repository

import (
	"github.com/chisdev/coupon/internal/repository/coupon"
	couponbooking "github.com/chisdev/coupon/internal/repository/coupon_booking"
	"github.com/chisdev/coupon/internal/repository/milestone"
	"github.com/chisdev/coupon/internal/repository/progress"
	"github.com/chisdev/coupon/internal/repository/reward"
	"github.com/chisdev/coupon/pkg/ent"
)

type Repository struct {
	ent                     *ent.Client
	MileStoneRepository     milestone.Milestone
	CouponRepository        coupon.Coupon
	CouponBookingRepository couponbooking.CouponBooking
	ProgressRepository      progress.Progress
	RewardRepository        reward.Reward
}

func New(ent *ent.Client) *Repository {
	return &Repository{
		MileStoneRepository:     milestone.New(ent),
		RewardRepository:        reward.New(ent),
		CouponRepository:        coupon.New(ent),
		CouponBookingRepository: couponbooking.New(ent),
		ProgressRepository:      progress.New(ent),
		ent:                     ent,
	}
}

func (r *Repository) GetEntClient() *ent.Client {
	return r.ent
}
