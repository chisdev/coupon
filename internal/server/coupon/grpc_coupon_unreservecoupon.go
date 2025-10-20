package coupon

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/chisdev/coupon/api"
)

func (s *couponServer) UnReserveCoupon(ctx context.Context, request *coupon.UnReserveCouponRequest) (*emptypb.Empty, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
