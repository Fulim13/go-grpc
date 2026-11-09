package main

import (
	grpc_service "go-grpc/idl/service"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:5678")
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	// Register Multiple Server
	grpc_service.RegisterStudentServer(server, Student{})
	err = server.Serve(lis)
	if err != nil {
		panic(err)
	}
}

// go run ./idl/server
