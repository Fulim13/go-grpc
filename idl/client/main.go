package main

import (
	"context"
	"fmt"
	grpc_service "go-grpc/idl/service"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient(
		"localhost:5678",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(1024), // Default SendMsg  limit is MaxInt32, Default RecvMsg limist is 4M, i adjust to 1KB
		),
	)

	if err != nil {
		fmt.Printf("dial failed %v\n", err)
		return
	}

	// MultiPlexing - 1 *grpc.ClientConn can be reused by many client
	wg := sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		wg.Go(func() {
			sendRequest(conn)
		})
	}
	wg.Wait()
}

func sendRequest(conn *grpc.ClientConn) {
	client := grpc_service.NewStudentClient(conn)
	request := grpc_service.QueryStudentRequest{
		Id:   123,
		Name: "Fu Lim",
	}

	resp, err := client.QueryStudent(context.Background(), &request, grpc.MaxCallSendMsgSize(1024))
	if err != nil {
		fmt.Printf("rpc error: %v\n", err)
	} else {
		fmt.Printf("response: %v\n", resp)
	}
}

// go run ./idl/client
