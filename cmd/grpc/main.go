package main

import (
	"log"
	"net"
	"os"
	"strconv"

	pb "github.com/lessbutter/alloff-api/api/grpcServer"
	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/internal/storage/mongo"
	"github.com/lessbutter/alloff-api/internal/storage/postgres"
	"google.golang.org/grpc"
)

func main() {
	conf := config.GetConfiguration()
	log.Println(conf)

	conn := mongo.NewMongoDB(conf)
	conn.RegisterRepos()

	pgconn := postgres.NewPostgresDB(conf)
	pgconn.RegisterRepos()

	// (TODO) Be Refactored
	config.InitIamPort(conf)
	config.InitSlack(conf)

	port := os.Getenv("GRPC_PORT")
	if port == "" {
		port = strconv.Itoa(conf.GRPC_PORT)
	}

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterProductServer(grpcServer, &pb.ProductService{})
	pb.RegisterProductGroupServer(grpcServer, &pb.ProductGroupService{})

	log.Printf("start gRPC server on %s port", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
