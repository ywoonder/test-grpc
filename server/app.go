package main

import (
	"context"
	"log"
	"net"

	"github.com/ywoonder/grpc/crack"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func (s *server) CrackToken(ctx context.Context, req *crack.CrackRequest) (*crack.CrackResponse, error) {
	log.Printf("REQUEST  %+v", *req)
	return &crack.CrackResponse{Success: true, Message: "success"}, nil
}

var port = ":8080"

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("fail to start server ", err.Error())
	}

	s := grpc.NewServer()
	crack.RegisterCrackServer(s, &server{})

	reflection.Register(s)

	log.Println("server is listening at port", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
