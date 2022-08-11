package main

import (
	"context"
	"flag"
	"fmt"
	"net"

	ma "github.com/alphabot-fi/T-801/internal/proto/ma"
	monitor "github.com/alphabot-fi/T-801/internal/proto/monitor"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	log               = logrus.New()
	port              = flag.Int("port", 50051, "server port")
	bts, rev, version string
)

type server struct {
	monitor.UnimplementedMonitorServer
	ma.UnimplementedMAServer
}

func (s *server) Ping(ctx context.Context, in *monitor.PingRequest) (*monitor.PingResponse, error) {
	log.Printf("monitor request: %v", in.GetRequestTime().AsTime())
	resp := monitor.PingResponse{
		ResponseTime: timestamppb.Now(),
		Version:      version,
	}
	return &resp, nil
}

func (s *server) S(ctx context.Context, in *ma.MARequest) (*ma.MAResponse, error) {
	log.Printf("SMA request: %v -- %v", in.GetRequestId(), in.GetRequestTime().AsTime())
	log.Printf("SMA request: period: %d", in.GetPeriod())
	period := in.GetPeriod()
	if period == 0 {
		err := status.Errorf(codes.InvalidArgument, "invalid period: %d", period)
		return nil, err
	}
	resp := ma.MAResponse{
		ResponseTime: timestamppb.Now(),
		RequestId:    in.GetRequestId(),
	}
	return &resp, nil
}

func main() {
	flag.Parse()
	version = fmt.Sprintf("%s::%s", bts, rev)
	fmt.Printf("moving average service: %s\n", version)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	ss := server{}
	monitor.RegisterMonitorServer(s, &ss)
	ma.RegisterMAServer(s, &ss)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
