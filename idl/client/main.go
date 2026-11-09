package main

import (
	"context"
	"fmt"
	grpc_service "go-grpc/idl/service"
	"log"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

func timer(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	begin := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	fmt.Printf("Time Used %d ms\n", time.Since(begin).Milliseconds())
	return err
}

func counter(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	fmt.Printf("add once %s\n", method)
	err := invoker(ctx, method, req, reply, cc, opts...)
	return err
}

func devKey(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	ctx = metadata.AppendToOutgoingContext(ctx, "dev_key", "123456")
	err := invoker(ctx, method, req, reply, cc, opts...)
	return err
}

func main() {
	creds, err := credentials.NewClientTLSFromFile("data/server.cert", "")
	if err != nil {
		log.Fatalf("Failed to generated credentials %v", err)
	}
	conn, err := grpc.NewClient(
		"localhost:5678",
		// grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(1024), // Default SendMsg  limit is MaxInt32, Default RecvMsg limist is 4M, i adjust to 1KB
		),
		// grpc.WithUnaryInterceptor(timer),
		grpc.WithChainUnaryInterceptor(timer, counter, devKey),
		grpc.WithTransportCredentials(creds),
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
