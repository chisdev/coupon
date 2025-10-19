package reward

type RewardEntity struct {
	MilestoneID     uint64
	ExpiredDuration *float64
	CurrencyID      *uint64
	UsageLimit      *int32
	CouponValue     float64
	ServiceIds      []uint64
}
