package coupon

import (
	"time"

	api "github.com/chisdev/coupon/api"
)

type CouponOpts struct {
	Code          string
	UserIDs       []string
	StoreIDs      []string
	ExpiredAt     *time.Time
	Limit         int32
	PageIndex     int32
	UsageLimit    int32
	UsageCount    int32
	ReservedCount int32
	Status        api.CouponStatus
	Type          api.CouponType
	CurrencyID    *uint64
	ServiceIds    []uint64
	SortMethods   []*api.SortMethod
}

type Option interface {
	Apply(*CouponOpts)
}

type funcOption func(*CouponOpts)

func (f funcOption) Apply(o *CouponOpts) {
	f(o)
}

func WithExpiredAt(expiredAt *time.Time) Option {
	return funcOption(func(co *CouponOpts) {
		co.ExpiredAt = expiredAt
	})
}

func WithPaging(limit, pageIndex int32) Option {
	return funcOption(func(co *CouponOpts) {
		co.Limit = limit
		co.PageIndex = pageIndex
	})
}

func WithUserIDs(userIDs []string) Option {
	return funcOption(func(co *CouponOpts) {
		co.UserIDs = userIDs
	})
}

func WithStoreIDs(storeIDs []string) Option {
	return funcOption(func(co *CouponOpts) {
		co.StoreIDs = storeIDs
	})
}

func WithCode(code string) Option {
	return funcOption(func(co *CouponOpts) {
		co.Code = code
	})
}

func WithUsageLimit(usageLimit int32) Option {
	return funcOption(func(co *CouponOpts) {
		co.UsageLimit = usageLimit
	})
}

func WithStatus(status api.CouponStatus) Option {
	return funcOption(func(co *CouponOpts) {
		co.Status = status
	})
}

func WithType(cType api.CouponType) Option {
	return funcOption(func(co *CouponOpts) {
		co.Type = cType
	})
}

func WithCurrencyID(currencyID *uint64) Option {
	return funcOption(func(co *CouponOpts) {
		co.CurrencyID = currencyID
	})
}

func WithReservedCount(reservedCount int32) Option {
	return funcOption(func(co *CouponOpts) {
		co.ReservedCount = reservedCount
	})
}

func WithUsageCount(usageCount int32) Option {
	return funcOption(func(co *CouponOpts) {
		co.UsageCount = usageCount
	})
}

func WithServiceIds(serviceIds []uint64) Option {
	return funcOption(func(co *CouponOpts) {
		co.ServiceIds = serviceIds
	})
}
