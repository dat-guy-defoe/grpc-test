package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	gAPI "grpcTest/pkg/api"
	"log"
	"strconv"
)

func main() {
	flag.Parse()
	if flag.NArg() < 2 {
		log.Fatal("noe enough arguments")
	}

	x, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	y, err := strconv.Atoi(flag.Arg(1))
	if err != nil {
		log.Fatal(err)
	}

	conn, err := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	client := gAPI.NewAdderClient(conn)
	res, err := client.Add(context.Background(), &gAPI.AddRequest{X: int32(x), Y: int32(y)})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res)
}
