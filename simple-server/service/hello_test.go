package service

import (
	"context"
	"testing"

	"gitlab.com/bfelipe/go-grpc/simple-server/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func TestHelloServiceSuccess(t *testing.T) {
	message := "Hello"
	subject := "Mr. Hands"
	req := pb.Request{
		Message: message,
		Subject: subject,
	}

	var svc HelloService
	ctx := context.Background()
	md := metadata.MD{}
	md.Append("api_key", "123")
	ctx = metadata.NewIncomingContext(ctx, md)

	res, err := svc.Greeting(ctx, &req)
	if err != nil {
		t.Fatal(err)
		return
	}

	expected := message + " " + subject
	if res.Message != expected {
		t.Fatalf("Failed Response Message: got %v, expected %v", res.Message, expected)
	}
}

func TestHelloServiceWithoutApiKeyMetadata(t *testing.T) {
	message := "Hello"
	subject := "Mr. Hands"
	req := pb.Request{
		Message: message,
		Subject: subject,
	}

	var svc HelloService
	ctx := context.Background()

	_, err := svc.Greeting(ctx, &req)
	expectedErr := status.Errorf(codes.Unauthenticated, "no metadata")
	if err.Error() != expectedErr.Error() {
		t.Fatalf("Fail Unauthenticated: got %v, expected %v", err, expectedErr)
		return
	}

}
