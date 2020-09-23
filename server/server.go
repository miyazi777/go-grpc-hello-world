package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"test/hello"

	"google.golang.org/grpc"
)

func main() {
	listenPort, err := net.Listen("tcp", ":19003")
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer(
		grpc.UnaryInterceptor(CustomInterceptor()),
	)
	hello.RegisterHelloServer(server, &Hello{})
	server.Serve(listenPort)
}

type Hello struct{}

func (h *Hello) Hello(cts context.Context, message *hello.HelloMessage) (*hello.HelloResponse, error) {
	res := hello.HelloResponse{Msg: fmt.Sprintf("hello %s", message.Name)}
	return &res, nil
}

func CustomInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, request interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		fmt.Printf("before process")
		helloMessage, _ := request.(*hello.HelloMessage)
		helloMessage.Name = fmt.Sprintf("<<< %s >>>", helloMessage.Name)

		res, err := handler(ctx, request)
		if err != nil {
			return nil, err
		}

		fmt.Println("after process")
		helloResponse, _ := res.(*hello.HelloResponse)
		helloResponse.Msg = fmt.Sprintf("*** %s ***", helloResponse.Msg)
		return res, nil
	}
}
