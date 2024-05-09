package main

import (
	"context"
	"log"
	"os"

	pb "github.com/muhammadsaefulr/nyoba-grpc/proto/protoexec" // Import generated protobuf package
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("tidak dapat terhubung: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("tidak dapat memberi salam: %v", err)
	}
	log.Printf("messageProto: %s", r.GetMessage())
}
