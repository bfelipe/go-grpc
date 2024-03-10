package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"gitlab.com/bfelipe/go-grpc/simple-server/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func main() {
	addr := ":3000"
	creds := insecure.NewCredentials()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	conn, err := grpc.DialContext(
		ctx,
		addr,
		grpc.WithTransportCredentials(creds),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Printf("Fail to call %s", err)
	}
	defer conn.Close()

	log.Printf("Connected to %s", addr)
	client := pb.NewHelloClient(conn)

	req := pb.Request{
		Message: "Hello",
		Subject: "Mr. Hands",
	}

	ctx, cancel = context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	ctx = metadata.AppendToOutgoingContext(ctx, "api_key", "123")

	res, err := client.Greeting(ctx, &req)
	if err != nil {
		log.Fatalf("Error %s", err)
	}
	fmt.Println(res)
}
