// the client will stream something to the server, and the server is just to send the response
package main

import (
	"io"
	"log"

	pb "github.com/fadhillatmojo/grpc-project-first/proto"
)

func (s *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
	var messages []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessagesList{Messages: messages})
		}

		if err != nil {
			return err
		}

		log.Printf("Got Request with name: %v", req.Name)
		messages = append(messages, "Hello", req.Name)

	}
}
