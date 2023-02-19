package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/zzuun/grpc-go/proto"
)

const (
	Address = "127.0.0.1:50052"
)

type server struct {
	pb.UnimplementedHelloServiceServer
}

func (s server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	message := fmt.Sprintf("Hello, %s!", req.GetName())
	return &pb.HelloResponse{Message: message}, nil
}

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterHelloServiceServer(s, server{})

	fmt.Println("Listen on " + Address)

	if err := s.Serve(listen); err != nil {
		return
	}
}
