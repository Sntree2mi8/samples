package grpcserver

import "google.golang.org/grpc"

func New() *grpc.Server {
	s := grpc.NewServer()
	return s
}
