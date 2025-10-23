package coupon

import "errors"

var (
	errMissingCustomerID = errors.New("missing customer id")
	errMissingStoreID    = errors.New("missing store id")
)
