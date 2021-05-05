package main

import (
	"context"
	"fmt"
	"net"

	pb "github.com/marsredskies/go-grpc-resizer/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func main() {
	listener, err := net.Listen("tcp", ":5300")

	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}
	options := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(options...)

	pb.RegisterGreetingServer(grpcServer, &server{})
	grpcServer.Serve(listener)
}

type server struct{}

func (s *server) Do(c context.Context, request *pb.Request) (response *pb.Response, err error) {
	greeting := fmt.Sprintf("Hello, %v", request.Message)

	response = &pb.Response{
		Message: greeting,
	}

	return response, nil
}
