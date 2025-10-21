package coupon

import (
	coupon "github.com/chisdev/coupon/api"
	"github.com/chisdev/coupon/internal/services"
	"go.uber.org/zap"
)

func NewServer(service *services.Services, logger *zap.Logger) coupon.CouponServer {
	return &couponServer{
		logger:  logger,
		service: service,
	}
}

type couponServer struct {
	coupon.UnimplementedCouponServer
	service *services.Services
	logger  *zap.Logger
}
