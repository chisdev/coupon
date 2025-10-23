package coupon

import (
	"context"

	coupon "github.com/chisdev/coupon/api"
	"go.uber.org/zap"
)

func (s *couponServer) ListCouponForCms(ctx context.Context, request *coupon.ListCouponForCmsRequest) (*coupon.ListCouponForCmsResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	resp, err := s.service.CouponService.ListCouponForCms(ctx, request)
	if err != nil {
		s.logger.Error("failed to list coupons for cms", zap.Error(err))
		return nil, err
	}

	return resp, nil
}
