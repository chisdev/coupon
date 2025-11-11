package coupon

import (
	"time"

	api "github.com/chisdev/coupon/api"
)

type CouponOpts struct {
	IDs         []uint64
	Code        string
	Codes       []string
	UserIDs     []string
	StoreIDs    []string
	ExpiredAt   *time.Time
	Limit       int32
	PageIndex   int32
	UsageLimit  *int32
	Status      api.CouponStatus
	Type        api.CouponType
	CurrencyID  *uint64
	ServiceIds  []string
	SortMethods []*api.SortMethod
	WithUsage   bool
}

type Option interface {
	Apply(*CouponOpts)
}

type funcOption func(*CouponOpts)

func (f funcOption) Apply(o *CouponOpts) {
	f(o)
}

func WithIDs(ids []uint64) Option {
	return funcOption(func(co *CouponOpts) {
		co.IDs = ids
	})
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

func WithCodes(codes []string) Option {
	return funcOption(func(co *CouponOpts) {
		co.Codes = codes
	})
}

func WithUsageLimit(usageLimit *int32) Option {
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

func WithServiceIds(serviceIds []string) Option {
	return funcOption(func(co *CouponOpts) {
		co.ServiceIds = serviceIds
	})
}

func WithUsage(withUsage bool) Option {
	return funcOption(func(co *CouponOpts) {
		co.WithUsage = withUsage
	})
}

func WithSortMethods(sortMethods []*api.SortMethod) Option {
	return funcOption(func(co *CouponOpts) {
		co.SortMethods = sortMethods
	})
}
