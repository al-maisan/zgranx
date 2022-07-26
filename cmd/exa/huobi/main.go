package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/alphabot-fi/T-801/internal/huobi"
	exa "github.com/alphabot-fi/T-801/internal/proto/exa"
	monitor "github.com/alphabot-fi/T-801/internal/proto/monitor"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	port              = flag.Int("port", 50051, "server port")
	bts, rev, version string
)

type server struct {
	monitor.UnimplementedMonitorServer
	exa.UnimplementedEXAServer
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
	if asset != exa.Asset_UNKNOWN_ASSET {
		log.Printf("exa GetBalance request: exchange: %s -- asset: %v", in.GetExchange(), asset)
	} else {
		log.Printf("exa GetBalance request: exchange: %s", in.GetExchange())
	}
	re := in.GetExchange().String()
	if strings.ToLower(re) != "huobi" {
		err := status.Errorf(codes.InvalidArgument, "wrong exchange: '%s'", re)
		return nil, err
	}
	apiKey := in.GetApiKey()
	if apiKey == "" {
		err := status.Error(codes.InvalidArgument, "no API key")
		return nil, err
	}
	apiSecret := in.GetApiSecret()
	if apiSecret == "" {
		err := status.Error(codes.InvalidArgument, "no API secret")
		return nil, err
	}

	bs, err := huobi.GetBalance(apiKey, apiSecret, asset.String())
	if err != nil {
		err := status.Error(codes.Internal, err.Error())
		return nil, err
	}

	resp := exa.GetBalanceResponse{
		ResponseTime: timestamppb.Now(),
		RequestId:    in.GetRequestId(),
	}
	for _, b := range bs {
		brr := exa.GetBalanceResponse_Result{
			Asset:   exa.Asset(exa.Asset_value[b.Asset]),
			Balance: b.Balance.String(),
		}
		resp.Balances = append(resp.Balances, &brr)
	}
	return &resp, nil
}

func main() {
	flag.Parse()
	version = fmt.Sprintf("%s::%s", bts, rev)
	fmt.Printf("exakkn: %s\n", version)
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
