package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	hello "test/hello"
)

func main() {
	con, err := grpc.Dial("127.0.0.1:19003", grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer con.Close()

	md := metadata.New(map[string]string{"authorization": "Bearer testtoken"})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	client := hello.NewHelloClient(con)
	message := &hello.HelloMessage{Name: "world"}
	res, err := client.Hello(ctx, message)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("%s\n", res.Msg)
}
