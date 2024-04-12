package main

import (
	"context"
	"grpc-go/grpc-go/hellopb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	hellopb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, request *hellopb.HelloRequest) (*hellopb.HelloReply, error) {
	return &hellopb.HelloReply{Message: "hello " + request.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	hellopb.RegisterGreeterService(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
