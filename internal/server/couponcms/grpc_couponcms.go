package couponcms

import (
	coupon "github.com/chisdev/coupon/api"
	"github.com/chisdev/coupon/internal/services"
	"go.uber.org/zap"
)

func NewServer(service *services.Services, logger *zap.Logger) coupon.CouponCmsServer {
	return &couponCmsServer{
		service: service,
		logger:  logger,
	}
}

type couponCmsServer struct {
	coupon.UnimplementedCouponCmsServer
	logger  *zap.Logger
	service *services.Services
}
