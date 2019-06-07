package main

import (
	"context"
	"fmt"
	"net"

	pb "github.com/ayeminag/greeter"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Printf("Couldn't listen on tcp port: %v\n", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err = s.Serve(lis); err != nil {
		fmt.Printf("Cloudn't start grpc server: %v\n", err)
	}
}

type server struct{}

func (s *server) GetGreeting(ctx context.Context, gr *pb.GreeterRequest) (*pb.GreeterResponse, error) {
	fmt.Printf("Recieved GetGreeting rpc Call from %v\n", gr.GetName())
	return &pb.GreeterResponse{Message: "Hello, " + gr.GetName()}, nil
}
