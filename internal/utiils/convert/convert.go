package convert

import (
	coupon "github.com/chisdev/coupon/api"
	"github.com/chisdev/coupon/pkg/ent"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertReward(e *ent.Reward) *coupon.Reward {
	return &coupon.Reward{
		Id:              e.ID,
		CouponValue:     e.CouponValue,
		ExpiredDuration: e.ExpiredDuration,
		CurrencyId:      e.CurrencyID,
		UsageLimit:      e.UsageLimit,
		MilestoneId:     e.MilestoneID,
		CouponType:      e.CouponType,
		// ServiceIds:      e.ServiceIds,
	}
}

func ConvertRewards(ents []*ent.Reward) []*coupon.Reward {
	rewards := make([]*coupon.Reward, 0, len(ents))
	for _, e := range ents {
		rewards = append(rewards, ConvertReward(e))
	}
	return rewards
}

func ConvertMilestone(e *ent.Milestone) *coupon.Milestone {
	return &coupon.Milestone{
		Id:        e.ID,
		Type:      e.MilestoneType,
		Threshold: e.Threshold,
		Step:      e.Step,
		StoreId:   e.StoreID,
		Rewards:   ConvertRewards(e.Edges.Reward),
	}
}

func ConvertMilestones(ents []*ent.Milestone) []*coupon.Milestone {
	milestones := make([]*coupon.Milestone, 0, len(ents))
	for _, e := range ents {
		milestones = append(milestones, ConvertMilestone(e))
	}
	return milestones
}

func ConvertUsage(couponBooking *ent.CouponBooking) *coupon.CouponUsage {
	out := &coupon.CouponUsage{
		Id:        couponBooking.ID,
		BookingId: couponBooking.BookingID,
		// ServiceIds: couponBooking.ServiceIds,
		CustomerId: couponBooking.CustomerID,
		Status:     couponBooking.Status,
		ReservedAt: timestamppb.New(couponBooking.CreatedAt),
	}
	if couponBooking.Status == coupon.CouponUsedStatus_COUPON_USED_STATUS_USED {
		out.UsedAt = timestamppb.New(couponBooking.UpdatedAt)
	}

	return out
}

func ConvertUsages(ents []*ent.CouponBooking) []*coupon.CouponUsage {
	usages := make([]*coupon.CouponUsage, 0, len(ents))
	for _, e := range ents {
		usages = append(usages, ConvertUsage(e))
	}
	return usages
}

func ConvertCoupon(e *ent.Coupon) *coupon.StoreCoupon {
	out := &coupon.StoreCoupon{
		Id:         e.ID,
		Code:       e.Code,
		CustomerId: e.CustomerID,
		StoreId:    e.StoreID,
		UsageLimit: e.UsageLimit,
		Status:     e.Status,
		CouponType: e.Type,
		CurrencyId: e.CurrencyID,
		// ServiceIds:    e.ServiceIds,
		CouponUsages:  ConvertUsages(e.Edges.CouponBookings),
		UsedCount:     0,
		ReservedCount: 0,
	}

	if e.ExpireAt != nil {
		out.ExpiredAt = timestamppb.New(*e.ExpireAt)
	}

	for _, usage := range out.CouponUsages {
		switch usage.Status {
		case coupon.CouponUsedStatus_COUPON_USED_STATUS_USED:
			out.UsedCount++
		case coupon.CouponUsedStatus_COUPON_USED_STATUS_RESERVED:
			out.ReservedCount++
		}
	}

	return out
}

func ConvertCoupons(ents []*ent.Coupon) []*coupon.StoreCoupon {
	coupons := make([]*coupon.StoreCoupon, 0, len(ents))
	for _, e := range ents {
		coupons = append(coupons, ConvertCoupon(e))
	}
	return coupons
}

func ConvertCurrency(ent *ent.Currency) *coupon.Currency {
	return &coupon.Currency{
		Id:   ent.ID,
		Name: ent.Name,
		Code: ent.Code,
	}
}

func ConvertCurrencies(ents []*ent.Currency) []*coupon.Currency {
	currencies := make([]*coupon.Currency, 0, len(ents))
	for _, e := range ents {
		currencies = append(currencies, ConvertCurrency(e))
	}
	return currencies
}

func ConvertProgress(ent *ent.Progress) *coupon.Progress {
	return &coupon.Progress{
		MilestoneId: ent.ID,
		Progress:    ent.Progress,
		PassCount:   ent.PassCount,
	}
}

func ConvertProgresses(ents []*ent.Progress) []*coupon.Progress {
	progresses := make([]*coupon.Progress, 0, len(ents))
	for _, e := range ents {
		progresses = append(progresses, ConvertProgress(e))
	}

	return progresses
}
