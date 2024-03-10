package service

import (
	"errors"
	"fmt"
	"io"

	"gitlab.com/bfelipe/go-grpc/simple-streamer/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type StreamerService struct {
	pb.UnimplementedScannerServer
}

func (ss *StreamerService) Scan(stream pb.Scanner_ScanServer) error {

	scans := []int64{}

	for {
		req, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return status.Errorf(codes.Internal, "Can't read stream")
		}

		scans = append(scans, req.ScanningData)
	}

	var count0 int64
	var count1 int64

	for _, v := range scans {
		if v == 0 {
			count0 += 1
		} else {
			count1 += 1
		}
	}

	fmt.Println(scans)
	fmt.Println(count0)
	fmt.Println(count1)

	res := pb.ScanningResponse{}
	if count0 < count1 || count0 == count1 {
		res.Result = pb.Result_ANOMALY
	} else if count0 > count1 {
		res.Result = pb.Result_CLEAR
	} else {
		res.Result = pb.Result_UNKOWN
	}
	fmt.Println(&res.Result)
	return stream.SendAndClose(&res)
}
