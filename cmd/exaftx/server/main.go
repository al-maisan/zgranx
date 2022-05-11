package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"strings"

	exa "github.com/al-maisan/zgranx/internal/proto/exa"
	monitor "github.com/al-maisan/zgranx/internal/proto/monitor"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	port              = flag.Int("port", 50051, "server port")
	bts, rev, version string
	exchange          = flag.String("exchange", "", "exchange to operate on")
)

type server struct {
	monitor.UnimplementedMonitorServer
	exa.UnimplementedEXAServer
}

func exchangeIsValid(exchange string) bool {
	exchanges := []string{"binance", "ftx", "kraken", "kucoin", "gate_io"}
	for _, e := range exchanges {
		if exchange == e {
			return true
		}
	}
	return false
}

func (s *server) Ping(ctx context.Context, in *monitor.PingRequest) (*monitor.PingResponse, error) {
	log.Printf("monitor request: %v", in.GetRequestTime().AsTime())
	resp := monitor.PingResponse{
		ResponseTime: timestamppb.Now(),
		Version:      version,
	}
	return &resp, nil
}

func (s *server) GetBalance(ctx context.Context, in *exa.GetBalanceRequest) (*exa.GetBalanceResponse, error) {
	log.Printf("exa GetBalance request: %v -- %v", in.GetRequestId(), in.GetRequestTime().AsTime())
	asset := in.GetAsset()
	if asset != "" {
		log.Printf("exa GetBalance request: exchange: %s -- asset: %v", in.GetExchange(), asset)
	} else {
		log.Printf("exa GetBalance request: exchange: %s", in.GetExchange())
	}
	uid := in.GetUserId()
	if uid < 1 {
		st := status.New(codes.InvalidArgument, "invalid user id")
		return nil, st.Err()
	}
	re := in.GetExchange().String()
	if strings.ToLower(re) != strings.ToLower(*exchange) {
		st := status.New(codes.InvalidArgument, fmt.Sprintf("wrong exchange: '%s'", re))
		return nil, st.Err()
	}
	resp := exa.GetBalanceResponse{
		ResponseTime: timestamppb.Now(),
		RequestId:    in.GetRequestId(),
	}
	return &resp, nil
}

func main() {
	flag.Parse()
	version = fmt.Sprintf("%s::%s", bts, rev)
	fmt.Printf("exaftx: %s\n", version)
	if *exchange == "" {
		log.Fatalf("exchange parameter not specified")
	}
	if !exchangeIsValid(*exchange) {
		log.Fatalf("invalid exchange: %s", *exchange)
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	ss := server{}
	monitor.RegisterMonitorServer(s, &ss)
	exa.RegisterEXAServer(s, &ss)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
