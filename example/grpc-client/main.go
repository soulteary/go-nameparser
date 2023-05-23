package main

import (
	"context"
	"log"
	"time"

	pb "grpc-client/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	TestName = "Dr. Juan Q. Xavier de la Vega III (Doc Vega)"
	GrpcAddr = "localhost:8081"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(GrpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewConverterClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.HumanName(ctx, &pb.ConvertRequest{Name: TestName})
	if err != nil {
		log.Fatalf("could not convert: %v", err)
	}
	log.Printf("Result: %s", r.GetMessage())
}
