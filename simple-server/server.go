package main

import (
	"fmt"

	"gitlab.com/bfelipe/go-grpc/simple-server/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	app := pb.Application{
		Name:      "my-app",
		Id:        1,
		IsBackend: true,
		Type:      pb.ApplicationType_API,
		Team: &pb.Team{
			Name:    "Awesome team",
			Members: []string{"Erick", "Sonia", "Matt"},
		},
		CreatedAt: timestamppb.Now(),
	}

	fmt.Printf("%s", app.String())
}

func grpcServer() {

}
