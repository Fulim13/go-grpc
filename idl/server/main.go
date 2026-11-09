package main

import (
	"context"
	"errors"
	"fmt"
	grpc_service "go-grpc/idl/service"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
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

func devKey(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("Unable to get the dev_key")
	}
	if value, exists := meta["dev_key"]; !exists {
		return nil, errors.New("Unable to get the dev_key")
	} else {
		if value[0] != "123456" {
			return nil, errors.New("dev_key is incorrect")
		} else {
			resp, err = handler(ctx, req)
			return
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:5678")
	if err != nil {
		panic(err)
	}

	creds, err := credentials.NewServerTLSFromFile("data/server.cert", "data/rsa_private_key.pem")
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer(
		// grpc.UnaryInterceptor(timer),
		grpc.ChainUnaryInterceptor(timer, counter, devKey),
		grpc.Creds(creds),
	)
	// Register Multiple Server
	grpc_service.RegisterStudentServer(server, Student{})
	err = server.Serve(lis)
	if err != nil {
		panic(err)
	}
}

// go run ./idl/server
// Generate a 1024-bit RSA private key:
// openssl genrsa -out data/rsa_private_key.pem 1024
//
// Generate the public key from the private key:
// openssl rsa -in data/rsa_private_key.pem -pubout -out data/rsa_public_key.pem
//
// PEM is a standard format that usually includes a header and a footer.
//
// Generate a self-signed certificate:
// openssl req -x509 -new -nodes -key data/rsa_private_key.pem -subj "/CN=localhost" -addext "subjectAltName=DNS:localhost" -days 3650 -out data/server.cert
