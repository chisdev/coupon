package coupon

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"

	coupon "github.com/chisdev/coupon/api"
	"github.com/chisdev/coupon/internal/utiils/generator"
)

func (s *couponServer) GetSecretCode(ctx context.Context, request *emptypb.Empty) (*coupon.GetSecretCodeResponse, error) {
	storeID := s.extractor.GetStoreID(ctx)
	if storeID == "" {
		s.logger.Error("store id is missing in context")
		return nil, fmt.Errorf("store id is missing in context")
	}

	code, err := generator.GenCodeV2("Z0123456789", 4)
	if err != nil {
		s.logger.Error("failed to generate code", zap.Error(err))
		return nil, fmt.Errorf("failed to generate code: %w", err)
	}

	if _, err := s.redis.SetV2(ctx, storeID, code, time.Hour*24); err != nil {
		s.logger.Error("failed to set secret code in redis", zap.Error(err))
		return nil, fmt.Errorf("failed to set secret code in redis: %w", err)
	}

	return &coupon.GetSecretCodeResponse{SecretCode: code}, nil
}
