package couponbooking

import "errors"

var (
	errCouponNotActive          = errors.New("coupon is not active")
	errCouponUsageLimitExceeded = errors.New("coupon usage limit exceeded")
	errServiceIdsNotAccepted    = errors.New("service IDs are not valid or accepted")
	errCustomerIdNotMatch       = errors.New("customer ID does not match the user ID")
)
