package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"

	"github.com/lessbutter/alloff-api/api/grpcServer/services"
	"github.com/lessbutter/alloff-api/cmd"
	pb "github.com/lessbutter/alloff-grpc-protos/gen/goalloff"
	"google.golang.org/grpc"
)

var (
	GitInfo   = "no info"
	BuildTime = "no datetime"
	Env       = "dev"
)

func main() {
	fmt.Println("Git commit information: ", GitInfo)
	fmt.Println("Build date, time: ", BuildTime)

	conf := cmd.SetBaseConfig(Env)

	port := os.Getenv("GRPC_PORT")
	if port == "" {
		port = strconv.Itoa(conf.GRPC_PORT)
	}

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
	pb.RegisterHomeTabItemServer(grpcServer, &services.HomeTabService{})
	pb.RegisterExhibitionServer(grpcServer, &services.ExhibitionService{})
	pb.RegisterTopBannerServer(grpcServer, &services.TopBannerService{})
	pb.RegisterAlloffCategoryServer(grpcServer, &services.AlloffCategoryService{})

	log.Printf("start gRPC server on %s port", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
