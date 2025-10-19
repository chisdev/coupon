package milestone

import "errors"

var (
	errMissingCurrencyID    = errors.New("Missing currency id")
	errInvalidCouponValue   = errors.New("Invalid coupon value")
	errMissingThreshold     = errors.New("Missing milestone threshold")
	errMissingStep          = errors.New("Missing milstone step")
	errIvalidExpiredTime    = errors.New("Invalid expired time")
	errInvalidMilestoneType = errors.New("Invalid milestone type")
)
