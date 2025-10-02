package server

import (
	mykit "github.com/ChisTrun/trunkit/pkg/api"
	"google.golang.org/grpc/reflection"

	pb0 "github.com/chisdev/coupon/api"
	"github.com/chisdev/coupon/internal/server/coupon"
	config "github.com/chisdev/coupon/pkg/config"
)

// Serve ...
func Serve(cfg *config.Config) {
	service := newService(cfg, []mykit.Option{}...)

	server := service.Server()
	pb0.RegisterCouponServer(server, coupon.NewServer())

	// Register reflection service on gRPC server.
	// Please remove if you it's not necessary for your service
	reflection.Register(server)

	service.Serve()
}
