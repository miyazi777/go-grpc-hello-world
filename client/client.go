package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	hello "test/hello"
)

func main() {
	con, err := grpc.Dial("127.0.0.1:19003", grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer con.Close()
	client := hello.NewHelloClient(con)
	message := &hello.HelloMessage{Name: "world"}
	res, _ := client.Hello(context.TODO(), message)
	fmt.Printf("%s\n", res.Msg)
}
