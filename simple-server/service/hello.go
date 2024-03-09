package service

import (
	"context"

	"gitlab.com/bfelipe/go-grpc/simple-server/pb"
)

type HelloService struct {
	pb.UnimplementedHelloServer
}

func (hs *HelloService) Greeting(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	res := pb.Response{
		Message: req.Message + " " + req.Subject,
	}
	return &res, nil
}
