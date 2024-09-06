package main

// this file to ability to call the server from our client send a request with the names, get back the response as a stream, and print it out in the client

import (
	"context"
	"io"
	"log"

	pb "github.com/fadhillatmojo/grpc-project-first/proto"
)

func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Streaming started")
	stream, err := client.SayHelloServerStreaming(context.Background(), names)

	if err != nil {
		log.Fatalf("Could not send names: %v", err)
	}

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while streaming %v", err)
		}
		log.Println(message)
	}
	log.Printf("Streaming finished")
}
