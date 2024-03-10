package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"gitlab.com/bfelipe/go-grpc/simple-server/pb"
	"gitlab.com/bfelipe/go-grpc/simple-server/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

func main() {
	addr := ":3000"
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen %s", err)
	}

	server := newGrpcServer()
	log.Printf("info: listening on %s", addr)
	if err := server.Serve(l); err != nil {
		log.Fatalf("Failed to serve %s", err)
	}
}

func newGrpcServer() *grpc.Server {
	server := grpc.NewServer(grpc.UnaryInterceptor(requestIdInterceptor))
	var helloService service.HelloService
	pb.RegisterHelloServer(server, &helloService)
	reflection.Register(server)
	return server
}

func requestIdInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "no metadata")
	}

	if len(md.Get("Request-Id")) == 0 {
		newMd := md.Copy()
		newMd.Append("Request-Id", "123")
		ctx = metadata.NewIncomingContext(ctx, newMd)
		defer fmt.Println("New Request-Id generated")
	}
	return handler(ctx, req)
}
