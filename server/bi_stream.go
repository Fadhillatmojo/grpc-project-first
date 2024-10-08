package main

import (
	"io"
	"log"

	pb "github.com/fadhillatmojo/grpc-project-first/proto"
)

func (s *helloServer) SayHelloBiderectionalStreaming(stream pb.GreetService_SayHelloBiderectionalStreamingServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Got request with name: %v", req.Name)
		res := &pb.HelloResponse{
			Message: "Hello " + req.Name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}

}
