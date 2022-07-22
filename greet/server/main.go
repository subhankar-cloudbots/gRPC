package main

import (
	"log"
	"net"

	pb "github.com/subhankar-cloudbots/grpc-basic/greet/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("failed")
	}

	log.Printf("Listening on 0.0.0.0:50051\n")

	s := grpc.NewServer()

	if err = s.Serve(lis); err != nil {
		log.Fatal("Failed to serve")
	}
}
