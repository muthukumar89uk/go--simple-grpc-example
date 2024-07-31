package main

import (
	"fmt"
	"log"
	"time"

	pb "example.com/greet-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8000"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect : %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NamesList{
		Names: []string{"Akhil", "Alice", "Bob"},
	}

	// Unary type (Client -> Single request , Server -> Single response)
	callSayHello(client)
	time.Sleep(2 * time.Second)
	fmt.Println()

	// Server-Streaming type (Client -> Single request , Server -> streaming response)
	callSayHelloServerStream(client, names)
	time.Sleep(2 * time.Second)
	fmt.Println()

	// Client-Streaming type (Client -> streaming request , Server -> Single response)
	callSayHelloClientStream(client, names)
	time.Sleep(2 * time.Second)
	fmt.Println()

	// Bidirectional-Streaming type (Client -> streaming request , Server -> streaming response)
	callHelloBidirectionalStream(client, names)
	time.Sleep(2 * time.Second)
	fmt.Println()
}
