package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"test/hello"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listenPort, err := net.Listen("tcp", ":19003")
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()
	hello.RegisterHelloServer(server, &Hello{})
	reflection.Register(server)
	server.Serve(listenPort)
}

type Hello struct{}

func (h *Hello) Hello(cts context.Context, message *hello.HelloMessage) (*hello.HelloResponse, error) {
	res := hello.HelloResponse{Msg: fmt.Sprintf("hello %s\n", message.Name)}
	return &res, nil
}
