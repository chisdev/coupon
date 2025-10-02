package coupon

import "errors"

var (
	errMissingCurrencyID              = errors.New("missing currency id")
	errConfigCouponTypeWithCurrencyID = errors.New("config coupon type with currency id")
	errInvalidCouponType              = errors.New("invalid coupon type")
	errInvalidExpiredAt               = errors.New("invalid expired at")
)
