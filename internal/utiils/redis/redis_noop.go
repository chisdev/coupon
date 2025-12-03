package redis

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/protobuf/proto"
)

type noop struct {
	Redis
}

func Noop() Redis {
	return &noop{}
}

func (d *noop) Set(ctx context.Context, key string, value proto.Message, expireTime time.Duration) (bool, error) {
	return false, fmt.Errorf("redis is not enabled")
}

func (d *noop) SetV2(ctx context.Context, key string, value interface{}, expireTime time.Duration) (bool, error) {
	return false, fmt.Errorf("redis is not enabled")
}

func (d *noop) Get(ctx context.Context, key string) ([]byte, error) {
	return nil, fmt.Errorf("redis is not enabled")
}

func (d *noop) Delete(ctx context.Context, key string) (bool, error) {
	return false, fmt.Errorf("redis is not enabled")
}
