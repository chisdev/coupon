package coupon

import (
	coupon "github.com/chisdev/coupon/api"
	"github.com/chisdev/coupon/internal/services"
	"github.com/chisdev/coupon/internal/utiils/extractor"
	"github.com/chisdev/coupon/internal/utiils/redis"
	"go.uber.org/zap"
)

func NewServer(service *services.Services, logger *zap.Logger, redis redis.Redis) coupon.CouponServer {
	return &couponServer{
		logger:    logger,
		service:   service,
		redis:     redis,
		extractor: extractor.New(),
	}
}

type couponServer struct {
	coupon.UnimplementedCouponServer
	service   *services.Services
	logger    *zap.Logger
	redis     redis.Redis
	extractor extractor.Extractor
}
