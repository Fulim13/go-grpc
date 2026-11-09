package main

import (
	"context"
	"fmt"
	grpc_model "go-grpc/idl/model"
	grpc_service "go-grpc/idl/service"
)

type Student struct {
	grpc_service.UnimplementedStudentServer
}

func (s Student) QueryStudent(ctx context.Context, query *grpc_service.QueryStudentRequest) (resp *grpc_service.QueryStudentResponse, err error) {
	fmt.Printf("Request: %+v\n", query)
	resp = &grpc_service.QueryStudentResponse{
		Students: []*grpc_model.Student{
			{
				Id:   123,
				Age:  24,
				Name: "Fu Lim",
			},
		},
	}
	return
}
