package reward

import "errors"

var (
	errMissingCurrencyID      = errors.New("missing currency id")
	errInvalidCouponType      = errors.New("invalid coupon type")
	errInvalidExpiredDuration = errors.New("invalid expired duration")
	errInvalidCouponValue     = errors.New("invalid coupon value")
	errInvaliedUsageLimit     = errors.New("invalid usage limit")
)
