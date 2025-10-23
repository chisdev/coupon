package coupon

import (
	"context"

	coupon "github.com/chisdev/coupon/api"
)

func (s *couponServer) ListProgress(ctx context.Context, request *coupon.ListProgressRequest) (*coupon.ListProgressResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	return &coupon.ListProgressResponse{}, nil
}
