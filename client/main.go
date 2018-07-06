package main

import (
	"context"
	"log"

	"github.com/vodaza36/go-grpc-gateway/api"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":7777", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Error creating conn: %v", err)
	}

	defer conn.Close()

	client := api.NewPingClient(conn)

	resp, err := client.SayHello(context.Background(), &api.PingMessage{Greeting: "Thomas"})

	if err != nil {
		log.Fatalf("Error invoking service: %v", err)
	}

	log.Printf("Response: %s", resp.Greeting)
}
