package main

import (
	"log"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/spf13/viper"

	"github.com/lessbutter/alloff-api/api/grpcServer/services"
	"github.com/lessbutter/alloff-api/cmd"
	pb "github.com/lessbutter/alloff-grpc-protos/gen/goalloff"
	"google.golang.org/grpc"
)

func main() {
	cmd.SetBaseConfig()
	port := viper.GetString("GRPC_PORT")

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_recovery.UnaryServerInterceptor(),
		)),
	)
	pb.RegisterProductServer(grpcServer, &services.ProductService{})
	pb.RegisterProductGroupServer(grpcServer, &services.ProductGroupService{})
	pb.RegisterBrandServer(grpcServer, &services.BrandService{})
	pb.RegisterNotificationServer(grpcServer, &services.NotiService{})
	pb.RegisterExhibitionServer(grpcServer, &services.ExhibitionService{})
	pb.RegisterAlloffCategoryServer(grpcServer, &services.AlloffCategoryService{})
	pb.RegisterAlloffSizeServer(grpcServer, &services.AlloffSizeService{})
	pb.RegisterHealthServer(grpcServer, &services.HealthService{})
	pb.RegisterTopBannerServer(grpcServer, &services.TopBannerService{})

	log.Printf("start gRPC server on %s port", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
