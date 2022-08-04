package main

import (
	"context"
	"gRPC-and-Protobuf-API/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type server struct {
}

func (s *server) mustEmbedUnimplementedAddServiceServer() {
	//TODO implement me
	panic("implement me")
}

func main() {
	//Starting server on tcp network on port 4040
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	//Starting grpc server
	srv := grpc.NewServer()
	proto.RegisterAddServiceServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(err)
	}
}

//Add function for adding a and b from request message (used in rpc service in protobuf)
func (s *server) Add(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a + b

	return &proto.Response{Result: result}, nil
}

//Multiply function for multiplying a and b from request message (used in rpc service in protobuf)
func (s *server) Multiply(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a * b

	return &proto.Response{Result: result}, nil
}
