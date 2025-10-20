package couponinternal

import (
	"github.com/chisdev/coupon/api"
)

func NewServer() coupon.CouponInternalServer {
	return &couponInternalServer{}
}

type couponInternalServer struct {
	coupon.UnimplementedCouponInternalServer
}
