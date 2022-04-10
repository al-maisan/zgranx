package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/al-maisan/zgranx/internal/monitor"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	port              = flag.Int("port", 50051, "The server port")
	bts, rev, version string
)

type server struct {
	pb.UnimplementedMonitorServer
}

func (s *server) Ping(ctx context.Context, in *pb.PingRequest) (*pb.PingResponse, error) {
	log.Printf("Received: %v", in.GetRequestTime().AsTime())
	resp := pb.PingResponse{
		ResponseTime: timestamppb.Now(),
		Version:      version,
	}
	return &resp, nil
}

func main() {
	flag.Parse()
	version = fmt.Sprintf("%s::%s", bts, rev)
	fmt.Printf("exaftx: %s\n", version)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMonitorServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
