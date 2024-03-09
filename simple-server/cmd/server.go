package main

import (
	"log"
	"net"

	"gitlab.com/bfelipe/go-grpc/simple-server/pb"
	"gitlab.com/bfelipe/go-grpc/simple-server/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	addr := ":3000"
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen %s", err)
	}

	server := grpc.NewServer()
	var helloService service.HelloService
	pb.RegisterHelloServer(server, &helloService)
	reflection.Register(server)

	log.Printf("info: listening on %s", addr)
	if err := server.Serve(l); err != nil {
		log.Fatalf("Failed to serve %s", err)
	}
}
