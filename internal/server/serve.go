package server

import (
	"context"
	"net/http"
	"strings"

	dbe "github.com/ChisTrun/database/pkg/ent"
	mykit "github.com/ChisTrun/trunkit/pkg/api"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"

	pb0 "github.com/chisdev/coupon/api"
	"github.com/chisdev/coupon/internal/repository"
	"github.com/chisdev/coupon/internal/server/coupon"
	"github.com/chisdev/coupon/internal/server/couponcms"
	"github.com/chisdev/coupon/internal/server/couponinternal"
	"github.com/chisdev/coupon/internal/services"
	"github.com/chisdev/coupon/internal/utiils/extractor"
	config "github.com/chisdev/coupon/pkg/config"
	"github.com/chisdev/coupon/pkg/ent"
	"github.com/chisdev/coupon/pkg/ent/migrate"
)

func customMetadataAnnotator(ctx context.Context, req *http.Request) metadata.MD {
	md := metadata.MD{}

	// Map tất cả các header có prefix "x-" vào metadata
	for name, values := range req.Header {
		lowerName := strings.ToLower(name)
		if strings.HasPrefix(lowerName, "x-") {
			md.Append(lowerName, values...)
		}
	}

	return md
}

// Serve ...
func Serve(cfg *config.Config) {
	service := newService(cfg, []mykit.Option{}...)

	logger := service.Logger()

	drv, err := dbe.Open("bookie_coupon", cfg.GetDatabase())
	ent := ent.NewClient(ent.Driver(drv))
	defer func() {
		if err := ent.Close(); err != nil {
			logger.Fatal("can not close ent client", zap.Error(err))
		}
	}()
	if err != nil {
		logger.Fatal("can not open ent client", zap.Error(err))
	}
	if err = ent.Schema.Create(context.Background(), migrate.WithDropIndex(true)); err != nil {
		logger.Fatal("can not init my database", zap.Error(err))
	}

	server := service.Server()

	repo := repository.New(ent)

	extractor := extractor.New()

	services := services.New(repo, extractor)

	couponServie := coupon.NewServer(services, logger)
	couponCmsService := couponcms.NewServer(services, logger)
	couponInternalService := couponinternal.NewServer(services, logger)

	grpcGatewayMux := runtime.NewServeMux(
		runtime.WithMetadata(customMetadataAnnotator),
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseProtoNames:   true,
				EmitUnpopulated: true,
				UseEnumNumbers:  false,
			},
		}),
	)
	service.HttpServeMux().Handle("/coupon/", grpcGatewayMux)
	service.HttpServeMux().Handle("/couponcms/", grpcGatewayMux)
	service.HttpServeMux().Handle("/couponint/", grpcGatewayMux)

	err = pb0.RegisterCouponHandlerServer(context.Background(), grpcGatewayMux, couponServie)
	if err != nil {
		logger.Fatal("can not register http coupon server", zap.Error(err))
	}
	err = pb0.RegisterCouponCmsHandlerServer(context.Background(), grpcGatewayMux, couponCmsService)
	if err != nil {
		logger.Fatal("can not register http coupon cms server", zap.Error(err))
	}
	err = pb0.RegisterCouponInternalHandlerServer(context.Background(), grpcGatewayMux, couponInternalService)
	if err != nil {
		logger.Fatal("can not register http coupon internal server", zap.Error(err))
	}

	pb0.RegisterCouponServer(server, couponServie)
	pb0.RegisterCouponCmsServer(server, couponCmsService)
	// Register reflection service on gRPC server.
	// Please remove if you it's not necessary for your service
	reflection.Register(server)

	service.Serve()
}
