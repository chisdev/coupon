package couponcms

import (
	"github.com/chisdev/coupon/api"
)

func NewServer() coupon.CouponCmsServer {
	return &couponCmsServer{}
}

type couponCmsServer struct {
	coupon.UnimplementedCouponCmsServer
}
