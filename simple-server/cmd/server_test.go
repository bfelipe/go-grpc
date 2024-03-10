package main

import (
	"context"
	"fmt"
	"net"
	"testing"
	"time"

	"gitlab.com/bfelipe/go-grpc/simple-server/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func TestServerE2ESuccess(t *testing.T) {
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatalf("Fail to listen %s", err)
		return
	}

	server := newGrpcServer()
	go server.Serve(l)

	port := l.Addr().(*net.TCPAddr).Port
	addr := fmt.Sprintf(":%d", port)
	creds := insecure.NewCredentials()
	conn, err := grpc.DialContext(
		context.Background(),
		addr,
		grpc.WithTransportCredentials(creds),
		grpc.WithBlock(),
	)
	if err != nil {
		t.Fatalf("Fail to create conn %s", err)
		return
	}

	client := pb.NewHelloClient(conn)
	message := "Hello"
	subject := "Mr. Hands"
	req := pb.Request{
		Message: message,
		Subject: subject,
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	ctx = metadata.AppendToOutgoingContext(ctx, "api_key", "123")

	res, err := client.Greeting(ctx, &req)
	if err != nil {
		t.Fatal(err)
		return
	}
	expected := message + " " + subject
	if res.Message != expected {
		t.Fatalf("Failed Response Message: got %v, expected %v", res.Message, expected)
	}
}

func TestServerE2EWithoutApiKeyMetadata(t *testing.T) {
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatalf("Fail to listen %s", err)
		return
	}

	server := newGrpcServer()
	go server.Serve(l)

	port := l.Addr().(*net.TCPAddr).Port
	addr := fmt.Sprintf(":%d", port)
	creds := insecure.NewCredentials()
	conn, err := grpc.DialContext(
		context.Background(),
		addr,
		grpc.WithTransportCredentials(creds),
		grpc.WithBlock(),
	)
	if err != nil {
		t.Fatalf("Fail to create conn %s", err)
		return
	}

	client := pb.NewHelloClient(conn)
	message := "Hello"
	subject := "Mr. Hands"
	req := pb.Request{
		Message: message,
		Subject: subject,
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	_, err = client.Greeting(ctx, &req)
	expectedErr := status.Errorf(codes.Unauthenticated, "no metadata")
	if err.Error() != expectedErr.Error() {
		t.Fatalf("Fail Unauthenticated: got %v, expected %v", err, expectedErr)
		return
	}

}
