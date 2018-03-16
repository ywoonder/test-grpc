package main

import (
	"context"
	"log"

	"github.com/ywoonder/grpc/crack"

	"google.golang.org/grpc"
)

const (
	address = "localhost:8080"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := crack.NewCrackClient(conn)

	// Make payment to the payment gateway and print out its response.
	request := &crack.CrackRequest{
		UserID:  1,
		TokenID: 12323,
	}

	r, err := c.CrackToken(context.Background(), request)
	if err != nil {
		log.Fatalf("could not crack Token: %v", err)
	}

	log.Printf("crack token response: %s", r.String())
}
