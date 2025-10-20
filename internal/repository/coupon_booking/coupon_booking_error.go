package couponbooking

import "errors"

var (
	errCouponNotActive          = errors.New("coupon is not active")
	errCouponUsageLimitExceeded = errors.New("coupon usage limit exceeded")
)
