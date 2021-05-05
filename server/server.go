package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net"
)

func main() {
	// listen at port 5300
	listener, err := net.Listen("tcp", ":5300")

	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}
	options := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(options...)

	greeting.RegisterGreetingServer(grpcServer, &server{})
	grpcServer.Serve(listener)
}

type server struct{}

func (s *server) Do(c context.Context, request *pb.Request) (response *greeting.Response, err error) {
	greeting := fmt.Sprintf("Hello, %v", request.Message)

	response = &greeting.Response{
		Message: greeting,
	}

	return response, nil
}
