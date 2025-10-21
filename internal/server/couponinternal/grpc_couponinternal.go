package couponinternal

import (
	coupon "github.com/chisdev/coupon/api"
	"github.com/chisdev/coupon/internal/services"
	"go.uber.org/zap"
)

func NewServer(service *services.Services, logger *zap.Logger) coupon.CouponInternalServer {
	return &couponInternalServer{
		logger:  logger,
		service: service,
	}
}

type couponInternalServer struct {
	coupon.UnimplementedCouponInternalServer
	service *services.Services
	logger  *zap.Logger
}
