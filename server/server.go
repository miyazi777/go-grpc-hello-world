package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"test/hello"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listenPort, err := net.Listen("tcp", ":19003")
	if err != nil {
		log.Fatal(err)
	}
	authInterceptor := grpc.UnaryInterceptor(grpc_auth.UnaryServerInterceptor(authenticate))
	server := grpc.NewServer(authInterceptor)
	hello.RegisterHelloServer(server, &Hello{})
	reflection.Register(server)
	server.Serve(listenPort)
}

type Hello struct{}

func (h *Hello) Hello(cts context.Context, message *hello.HelloMessage) (*hello.HelloResponse, error) {
	res := hello.HelloResponse{Msg: fmt.Sprintf("hello %s\n", message.Name)}
	return &res, nil
}

func authenticate(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "Bearer")
	if err != nil {
		return nil, err
	}
	fmt.Println(token)
	if token != "testtoken" {
		return nil, errors.New("unauthorized")
	}
	return ctx, nil
}
