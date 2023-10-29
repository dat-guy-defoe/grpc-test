package main

import (
	"google.golang.org/grpc"
	"grpcTest/pkg/adder"
	api2 "grpcTest/pkg/api"
	"log"
	"net"
)

func main() {
	s := grpc.NewServer()
	srv := &adder.GRPCServer{}
	api2.RegisterAdderServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
