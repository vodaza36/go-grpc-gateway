package main

import (
	"context"
	"log"

	"github.com/vodaza36/go-grpc-gateway/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Authentication holds the login/password
type Authentication struct {
	Login    string
	Password string
}

// GetRequestMetadata gets the current request metadata
func (a *Authentication) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{
		"login":    a.Login,
		"password": a.Password,
	}, nil
}

// RequireTransportSecurity indicates whether the credentials requires transport security
func (a *Authentication) RequireTransportSecurity() bool {
	return true
}

func main() {
	// Create the client TLS credentials
	creds, err := credentials.NewClientTLSFromFile("cert/server.crt", "")
	if err != nil {
		log.Fatalf("could not load tls cert: %s", err)
	}

	// Setup the login/pass
	auth := Authentication{
		Login:    "john",
		Password: "doe",
	}

	conn, err := grpc.Dial("go-grpc-gateway.hochbichler.at:7777", grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(&auth))

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
