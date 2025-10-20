package couponcms

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/chisdev/coupon/api"
)

func (s *couponCmsServer) ConfirmCouponUsage(ctx context.Context, request *coupon.ConfirmCouponUsageRequest) (*emptypb.Empty, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
