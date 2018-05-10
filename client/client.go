// main_client.go
package main

import (
	"log"
	"os"

	"golang.org/x/net/context"

	pb "github.com/xuedev/sample-grpc/helloworld"

	"google.golang.org/grpc"
)

const (
	addr         = "localhost:8888"
	default_name = "world"
)

func main() {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Failed to connect to server")
	}

	defer conn.Close()
	name := default_name

	c := pb.NewGreeterClient(conn)
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ret, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatal("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", ret.Message)

}
