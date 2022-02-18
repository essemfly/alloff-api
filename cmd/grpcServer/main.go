package main

import (
	"log"
	"net"
	"os"
	"strconv"

	pb "github.com/lessbutter/alloff-api/api/grpcServer"
	"github.com/lessbutter/alloff-api/api/grpcServer/services"
	"github.com/lessbutter/alloff-api/cmd"
	"google.golang.org/grpc"
)

func main() {
	conf := cmd.SetBaseConfig()

	port := os.Getenv("GRPC_PORT")
	if port == "" {
		port = strconv.Itoa(conf.GRPC_PORT)
	}

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
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
