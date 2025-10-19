package milestone

import "errors"

var (
	errStoreIdNotFound        = errors.New("store id not found")
	errMissingNames           = errors.New("missing milestone names")
	errInvalidMilestoneConfig = errors.New("invalid milestone configuration")
)
