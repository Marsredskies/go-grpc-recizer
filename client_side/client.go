package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	greeting "github.com/marsredskies/go-grpc-resizer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func main() {
	// dialing option
	options := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("ur name, pussy:")
	msg, _ := reader.ReadString('\n')
	// connection tartget
	connection, err := grpc.Dial("127.0.0.1:5300", options...)

	if err != nil {
		grpclog.Fatalf("failed to dial: %v", err)
	}

	defer connection.Close()
	// initialising client connection
	client := greeting.NewGreetingClient(connection)
	// request to server
	request := &greeting.Request{
		Message: msg,
	}
	//
	response, err := client.Do(context.Background(), request)

	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}

	fmt.Println(response.Message)
}
