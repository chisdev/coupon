package coupon

import (
	"context"

	"github.com/chisdev/coupon/api"
)

func (s *couponServer) HealCheck(ctx context.Context, request *coupon.HealCheckRequest) (*coupon.HealCheckResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	return &coupon.HealCheckResponse{}, nil
}
