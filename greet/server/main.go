package main

import (
	"log"
	"net"

	pb "github.com/machingclee/gogrpc/greet/proto"
	"google.golang.org/grpc"
)

var addr string = ":50051"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen due to: %v\n", err)
	}

	log.Printf("Listening on %v\n", addr)

	s := grpc.NewServer()

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve due to: %v\n", err)
	}
}
