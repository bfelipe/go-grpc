package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"gitlab.com/bfelipe/go-grpc/simple-streamer/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	client := pb.NewScannerClient(conn)

	ctx, cancel = context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	stream, err := client.Scan(ctx)
	if err != nil {
		log.Fatalf("Error %s", err)
	}

	req := pb.ScanningRequest{
		RobotId: "wall-e",
	}

	// points := []int64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	// points := []int64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	// points := []int64{0, 0, 0, 0, 0, 1, 1, 1, 1, 1}
	points := []int64{}
	for _, v := range points {
		req.ScanningData = v
		if err := stream.Send(&req); err != nil {
			log.Fatalf("Error %s", err)
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error %s", err)
	}

	fmt.Println(res.Result)
}
