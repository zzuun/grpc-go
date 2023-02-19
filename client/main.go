package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/credentials/insecure"
	"log"

	"google.golang.org/grpc"

	pb "github.com/zzuun/grpc-go/proto"
)

const (
	Address = "127.0.0.1:50052"
)

func main() {
	conn, err := grpc.Dial(Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}

	defer conn.Close()

	c := pb.NewHelloServiceClient(conn)

	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: "Zunnorain"})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(r.Message)
}
