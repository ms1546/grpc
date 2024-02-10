package main

import (
	"context"
	pb "main/protos"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

type GreeterServiceMock struct {
	mock.Mock
}

func (m *GreeterServiceMock) SayHello(ctx context.Context, in *pb.HelloRequest, opts ...grpc.CallOption) (*pb.HelloReply, error) {
	args := m.Called(ctx, in, opts)
	return args.Get(0).(*pb.HelloReply), args.Error(1)
}

func TestGreet(t *testing.T) {
	client := new(GreeterServiceMock)
	expectedResponse := &pb.HelloReply{Message: "Hello test"}

	client.On("SayHello", mock.Anything, &pb.HelloRequest{Name: "test"}, mock.Anything).Return(expectedResponse, nil)

	message := greet(client, "test")

	assert.Equal(t, "Hello test", message)
	client.AssertExpectations(t)
}
