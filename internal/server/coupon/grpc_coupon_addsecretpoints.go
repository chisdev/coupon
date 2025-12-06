package coupon

import (
	"context"
	"fmt"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"

	coupon "github.com/chisdev/coupon/api"
)

func (s *couponServer) AddSecretPoints(ctx context.Context, request *coupon.AddSecretPointsRequest) (*emptypb.Empty, error) {

	redisCodeBytes, err := s.redis.Get(ctx, request.StoreId)
	if err != nil {
		s.logger.Error("failed to get secret code from redis", zap.Error(err))
		return nil, err
	}

	redisCode := string(redisCodeBytes)
	if redisCode != request.SecretCode {
		s.logger.Error("invalid secret code", zap.String("expected", redisCode), zap.String("got", request.SecretCode))
		return nil, fmt.Errorf("invalid secret code")
	}

	if request.Points < 0 {
		s.logger.Error("points to add cannot be negative", zap.Int32("points", request.Points))
		return nil, fmt.Errorf("points to add cannot be negative")
	}

	if err := s.service.ProgressService.AddSecretPoints(ctx, request); err != nil {
		s.logger.Error("error while calling AddSecretPoints", zap.Error(err))
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
