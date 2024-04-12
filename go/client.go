package main

import (
	"context"
	"grpc-go/grpc-go/hellopb"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	defer conn.Close()

	c := hellopb.NewGreeterClient(conn)
	name := "world"
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &hellopb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
