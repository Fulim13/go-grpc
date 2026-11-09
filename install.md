# Install Protocol Buffers compiler

- read and parse .proto files and check syntax,
- but do not geenrated go code, only hands the parsed definitions to language plugin)

```sh
brew install protobuf
```

# Install plugin for generate Go code

```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

- protobuf is to generate the message code to xxx.pb.go
- grpc is to generate the services (RPC methods - client interface stub, and server interface to implement) to xxx_grpc.pb.go
