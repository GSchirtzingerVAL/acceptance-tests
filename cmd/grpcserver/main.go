package main

import (
	"log"
	"net"

	"github.com/quii/go-specs-greet/adapters/grpcserver"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	s := grpcserver.Server()

	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
