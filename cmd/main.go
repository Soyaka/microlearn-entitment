package main

import (
	"log"
	"net"
	"os"

	entitment "github.com/Soyaka/microlearn-entitment/api/gen"
	"github.com/Soyaka/microlearn-entitment/internal/handlers"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("GRPC_SERVER_PORT")
	listener, err := net.Listen("tcp", "127.0.0.1:"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	GlobalService := handlers.NewGlobalHandler()
	go GlobalService.MessageReader()

	server := grpc.NewServer()

	RegisterServerServices(server, GlobalService)
	if err = server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func RegisterServerServices(server *grpc.Server, service *handlers.GlobalHandler) {
	entitment.RegisterBookmarkServiceServer(server, service)
	entitment.RegisterInterestServiceServer(server, service)
	entitment.RegisterProgressServiceServer(server, service)
	entitment.RegisterSubscriptionServiceServer(server, service)
}
