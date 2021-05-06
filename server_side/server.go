package main

import (
	"context"
	"fmt"
	"github.com/marsredskies/go-grpc-resizer"
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

func (s *server) Do(c context.Context, request *greeting.Request) (response *greeting.Response, err error) {
	greet := fmt.Sprintf("Hello, %v", request.Message)

	response = &greeting.Response{
		Message: greet,
	}

	return response, nil
}
