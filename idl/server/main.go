package main

import (
	"context"
	"fmt"
	grpc_service "go-grpc/idl/service"
	"net"
	"time"

	"google.golang.org/grpc"
)

func timer(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	begin := time.Now()
	resp, err = handler(ctx, req)
	fmt.Printf("Time Used %d ms\n", time.Since(begin).Milliseconds())
	return
}

func counter(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	fmt.Printf("%s add once\n", info.FullMethod)
	resp, err = handler(ctx, req)
	return
}

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:5678")
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer(
		// grpc.UnaryInterceptor(timer),
		grpc.ChainUnaryInterceptor(timer, counter),
	)
	// Register Multiple Server
	grpc_service.RegisterStudentServer(server, Student{})
	err = server.Serve(lis)
	if err != nil {
		panic(err)
	}
}

// go run ./idl/server
