package rpc

import (
	"context"
	"log"
	"net"

	"github.com/soulteary/go-nameparser/internal/bridge"
	"github.com/soulteary/go-nameparser/internal/define"
	pb "github.com/soulteary/go-nameparser/pkg/pb"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedConverterServer
}

func (s *server) HumanName(ctx context.Context, in *pb.ConvertRequest) (*pb.ConvertReply, error) {
	return &pb.ConvertReply{Message: bridge.Convert(in.GetName())}, nil
}

func Launch() {
	lis, err := net.Listen("tcp", define.GRPC_PORT)
	if err != nil {
		log.Fatalf("GRPC server failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterConverterServer(s, &server{})
	log.Printf("GRPC server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("GRPC server failed to serve: %v", err)
	}
}
