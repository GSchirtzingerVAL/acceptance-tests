package grpcserver

import (
	"context"

	"github.com/quii/go-specs-greet/domain/interactions"
	"google.golang.org/grpc"
)

type GreetServer struct {
	UnimplementedGreeterServer
}

func (g GreetServer) Curse(_ context.Context, request *CurseRequest) (*CurseReply, error) {
	return &CurseReply{Message: interactions.Curse(request.Name)}, nil
}
func (g GreetServer) Greet(_ context.Context, request *GreetRequest) (*GreetReply, error) {
	return &GreetReply{Message: interactions.Greet(request.Name)}, nil
}

func Server() *grpc.Server {
	server := grpc.NewServer()
	RegisterGreeterServer(server, &GreetServer{})
	return server
}
