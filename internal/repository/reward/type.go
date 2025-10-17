package reward

type RewardEntity struct {
	ID              uint64
	MilestoneID     uint64
	ExpiredDuration *float64
	CurrencyID      *uint64
	UsageLimit      *int32
	CouponValue     float64
}
