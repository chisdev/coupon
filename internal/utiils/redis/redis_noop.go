package redis

import (
	"context"
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
	return false, nil
}
func (d *noop) Get(ctx context.Context, key string) ([]byte, error) {
	return nil, nil
}

func (d *noop) Delete(ctx context.Context, key string) (bool, error) {
	return false, nil
}