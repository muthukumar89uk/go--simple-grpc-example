# Go With Simple gRPC 

This repository demonstrates how to set up a Go project using gRPC to create a service that supports various types of greet messages.

## Features

- gRPC service implementation in Go
- Support for unary, server streaming, client streaming, and bidirectional streaming RPCs
- Protocol Buffers for defining the gRPC service and messages

## Requirements

- Go 1.15 or higher
- Protocol Buffers compiler (protoc)
- gRPC and Protocol Buffers Go plugins

## Getting Started

### Installation

1. **Clone the repository:**

    ```sh
    git clone https://github.com/muthukumar89uk/go-simple-grpc-example.git
    ```
Click here to directly [download it](https://github.com/muthukumar89uk/go-simple-grpc-example/zipball/master).

2. **Install Protocol Buffers compiler:**

    Follow the instructions for your operating system on the [official Protocol Buffers website](https://grpc.io/docs/protoc-installation/).

3. **Install gRPC and Protocol Buffers Go plugins:**

    ```sh
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
    ```
    Update your PATH so that the protoc compiler can find the plugins:

    ```
      $ export PATH="$PATH:$(go env GOPATH)/bin"
    ```
4. **Install Go dependencies:**

    ```sh
    go mod tidy
    ```

### Define the gRPC Service

1. **Create the proto file:**

    Create a `proto` directory and a `greet.proto` file inside it with the following content:

    ```proto
       syntax="proto3";

       package greet_service;
       
       option go_package="example.com/greet-grpc;greet_grpc";
       
       service GreetService{
           rpc SayHello (NoParam) returns (HelloResponse);
           rpc SayHelloServerStreaming (NamesList) returns (stream HelloResponse);
           rpc SayHelloClientStreaming (stream HelloRequest) returns ( MessagesList);
           rpc SayHelloBidirectionalStreaming (stream HelloRequest) returns (stream HelloResponse);
       }
       
       message NoParam{};
       
       message HelloRequest{
           string name = 1;
       };
       
       message HelloResponse{
           string message = 1;
       };
       
       message NamesList{
          repeated string names = 1;
       };
       
       message MessagesList{
           repeated string messages = 1;
       };
    ```

2. **Generate Go code from the proto file:**

    ```sh
     protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative <foldername>/<filename>.proto
    ```
### Run the Application

1. **Run the gRPC Server:**

    ```sh
    go run server/main.go
    ```
2. **Run the gRPC Client:**

    ```sh
    go run client/main.go
    ```


