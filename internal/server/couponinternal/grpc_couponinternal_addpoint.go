package couponinternal

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	coupon "github.com/chisdev/coupon/api"
)

func (s *couponInternalServer) AddPoint(ctx context.Context, request *coupon.AddPointRequest) (*emptypb.Empty, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
