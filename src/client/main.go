package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "main/protos"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GreeterServiceClient interface {
	SayHello(ctx context.Context, in *pb.HelloRequest, opts ...grpc.CallOption) (*pb.HelloReply, error)
}

func greet(c GreeterServiceClient, name string) string {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return r.GetMessage()
}

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	name := "world"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	message := greet(c, name)
	log.Printf("Greeting: %v", message)
}
