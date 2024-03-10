package service

import (
	"context"
	"log"

	"gitlab.com/bfelipe/go-grpc/simple-server/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type HelloService struct {
	pb.UnimplementedHelloServer
}

func (hs *HelloService) Greeting(ctx context.Context, req *pb.Request) (*pb.Response, error) {

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(md.Get("api_key")) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "no metadata")
	}

	log.Printf("Metadata %v", md)

	res := pb.Response{
		Message: req.Message + " " + req.Subject,
	}
	return &res, nil
}
