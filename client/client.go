package client

import (
	"context"
	"fmt"
	"os"

	pb "github.com/marsredskies/go-grpc-resizer/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func main() {
	// dialing option
	options := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	// message field number
	arg := os.Args
	// connection tartget
	connection, err := grpc.Dial("127.0.0.1:5300", options...)

	if err != nil {
		grpclog.Fatalf("failed to dial: %v", err)
	}

	defer connection.Close()
	// initialising client connection
	client := pb.NewGreetingClient(connection)
	// request to server
	request := &pb.Request{
		Message: arg[1],
	}
	//
	response, err := client.Do(context.Background(), request)

	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}

	fmt.Println(response.Message)
}
