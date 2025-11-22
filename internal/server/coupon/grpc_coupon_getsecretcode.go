package coupon

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	coupon "github.com/chisdev/coupon/api"
)

func (s *couponServer) GetSecretCode(ctx context.Context, request *emptypb.Empty) (*coupon.GetSecretCodeResponse, error) {

	return &coupon.GetSecretCodeResponse{}, nil
}
