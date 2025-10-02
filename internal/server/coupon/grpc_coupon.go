package coupon

import (
	"github.com/chisdev/coupon/api"
)

func NewServer() coupon.CouponServer {
	return &couponServer{}
}

type couponServer struct {
	coupon.UnimplementedCouponServer
}
