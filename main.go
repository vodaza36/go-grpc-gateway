package main

import (
	"log"
	"net"

	"github.com/vodaza36/go-grpc-gateway/api"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":7777")

	if err != nil {
		log.Fatalf("Error creating listening port: %v", err)
	}

	s := api.Server{}

	grpcServer := grpc.NewServer()

	api.RegisterPingServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Error serving gRPC API: %v", err)
	}

}
