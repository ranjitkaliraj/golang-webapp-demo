package main

import (
	"log"
	"net"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	pb "github.com/ranjitkaliraj/go-test-crm-service/cmd/grpc"
	"github.com/ranjitkaliraj/go-test-crm-service/cmd/service"
)

func main() {

	// Note - GRPC URL - localhost:50051, HTTP URL - localhost:8080

	// Set up gRPC server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Register GRPC endpoints
	grpcServer := grpc.NewServer()
	contactGrpcService := service.NewContactGrpcService()
	pb.RegisterContactServiceServer(grpcServer, contactGrpcService)

	// Start GRPC Server
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve gRPC server: %v", err)
		}
	}()

	// Register HTTP endpoints
	router := gin.Default()
	contactHttpService := service.NewContactHttpService()
	{
		router.POST("/contacts", contactHttpService.Create)
		router.GET("/contacts/:id", contactHttpService.Get)
		router.PUT("/contacts/:id", contactHttpService.Update)
		router.DELETE("/contacts/:id", contactHttpService.Delete)
		router.GET("/contacts", contactHttpService.List)
	}

	// Start HTTP Server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run gin server: %v", err)
	}
}
