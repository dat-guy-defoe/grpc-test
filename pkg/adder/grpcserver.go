package adder

import (
	"context"
	api2 "grpcTest/pkg/api"
)

type GRPCServer struct {
	api2.UnimplementedAdderServer
}

func (s *GRPCServer) Add(ctx context.Context, req *api2.AddRequest) (*api2.AddResponse, error) {
	return &api2.AddResponse{Result: req.GetX() + req.GetY()}, nil
}
