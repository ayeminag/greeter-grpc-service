package main

import (
	"context"
	"time"
	"fmt"
	"google.golang.org/grpc"
	pb "github.com/ayeminag/greeter"
)
func main() {
	serverAddress := "localhost:50051"
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Failed to connect: %v\n", err)
	}

	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetGreeting(ctx, &pb.GreeterRequest{Name: "Mr. Client"})

	if err != nil {
		fmt.Printf("Error calling rpc: %v\n", err)
	}

	fmt.Printf("Greeting: %s\n", r.GetMessage())
}