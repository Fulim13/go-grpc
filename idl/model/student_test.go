package grpc_model_test

import (
	"encoding/json"
	"fmt"
	grpc_model "go-grpc/idl/model"
	"testing"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
)

// go get google.golang.org/protobuf/encoding/protojson
// go get google.golang.org/protobuf/encoding/prototext
// go get google.golang.org/protobuf/proto

var stu = &grpc_model.Student{
	Id:        1,
	Name:      "Fu Lim",
	Locations: []string{"Malaysia", "Singapore"},
	Scores:    map[string]float32{"Math": 98.7},
}

func TestStudentSerialize(t *testing.T) {
	// Proto Serialization
	bb, _ := proto.Marshal(stu)
	fmt.Println("Proto Serialization: ", string(bb))

	//stu1 := new(grpc_model.Student) // already allocated memory for stu1 struct, not a nill pointer
	// var stud1 *grpc_model.Student // nil pointer, does not have allocated memory
	var stu1 grpc_model.Student
	proto.Unmarshal(bb, &stu1)
	fmt.Println("Proto Deserialization: ,", stu1.Id, stu1.Name, stu1.Locations, stu1.Age)

	// ProtoJSON Serialization
	bj, _ := protojson.Marshal(stu)
	fmt.Println("ProtoJSON Serialization: ", string(bj))
	var stu2 grpc_model.Student
	protojson.Unmarshal(bj, &stu2)
	fmt.Println("ProtoJSON Deserialization: ,", stu2.Id, stu2.Name, stu2.Locations, stu2.Age)

	// ProtoText Serialization
	bt, _ := prototext.Marshal(stu)
	fmt.Println("ProtoText Serialization: ", string(bj))
	var stu3 grpc_model.Student
	prototext.Unmarshal(bt, &stu3)
	fmt.Println("ProtoText Deserialization: ,", stu3.Id, stu3.Name, stu3.Locations, stu3.Age)
}

func BenchmarkJson(b *testing.B) {
	buffer := new(grpc_model.Student)
	for i := 0; i < b.N; i++ {
		bs, _ := json.Marshal(stu)
		json.Unmarshal(bs, buffer)
	}
}

func BenchmarkProto(b *testing.B) {
	buffer := new(grpc_model.Student)
	for i := 0; i < b.N; i++ {
		bs, _ := proto.Marshal(stu)
		proto.Unmarshal(bs, buffer)
	}
}

func BenchmarkProtoJson(b *testing.B) {
	buffer := new(grpc_model.Student)
	for i := 0; i < b.N; i++ {
		bs, _ := protojson.Marshal(stu)
		protojson.Unmarshal(bs, buffer)
	}
}

func BenchmarkProtoText(b *testing.B) {
	buffer := new(grpc_model.Student)
	for i := 0; i < b.N; i++ {
		bs, _ := prototext.Marshal(stu)
		prototext.Unmarshal(bs, buffer)
	}
}

// go test -v ./idl/model -run=^TestStudentSerialize -count=1
// go test ./idl/model -bench=^Benchmark -run=^ -count=1 -benchmem
