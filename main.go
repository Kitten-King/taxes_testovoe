package main

import (
	"log"
	"net"

	"tax_service/internal/db"
	"tax_service/internal/repository"
	"tax_service/internal/service"

	"github.com/Kitten-King/tax_sdk/pb"
	"google.golang.org/grpc"
)

func main() {
	log.Println("Starting Tax Service...")

	database := db.Connect()
	defer database.Close()

	repo := repository.NewTaxRepository(database)
	taxServer := service.NewTaxServer(repo)

	grpcServer := grpc.NewServer()

	pb.RegisterTaxServiceServer(grpcServer, taxServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	log.Println("Tax Service is listening on gRPC port :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC: %v", err)
	}
}
