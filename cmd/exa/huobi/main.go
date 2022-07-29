package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/alphabot-fi/T-801/internal/huobi"
	exa "github.com/alphabot-fi/T-801/internal/proto/exa"
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

func assetOfInterest(symbol string) bool {
	aoi := map[string]bool{
		"btc":   true,
		"eth":   true,
		"bnb":   true,
		"ada":   true,
		"sol":   true,
		"dot":   true,
		"avax":  true,
		"matic": true,
		"ltc":   true,
		"usdt":  true,
		"usd":   true,
		"eur":   true,
		"jpy":   true,
		"chf":   true,
		"cad":   true,
		"krw":   true,
		"usdc":  true,
	}
	_, ok := aoi[symbol]
	return ok
}

func (s *server) GetBalances(ctx context.Context, in *exa.GetBalancesRequest) (*exa.GetBalancesResponse, error) {
	log.Printf("exa GetBalances request: %v -- %v", in.GetRequestId(), in.GetRequestTime().AsTime())
	log.Printf("exa GetBalances request: exchange: %s", in.GetExchange())
	re := in.GetExchange()
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

	bdata, err := huobi.GetBalances(apiKey, apiSecret)
	if err != nil {
		err := status.Error(codes.Internal, err.Error())
		return nil, err
	}

	resp := exa.GetBalancesResponse{
		ResponseTime: timestamppb.Now(),
		RequestId:    in.GetRequestId(),
	}
	for _, bd := range bdata {
		account := fmt.Sprintf("%d", bd.Account)
		for _, b := range bd.Balances {
			if b.Type != "trade" {
				continue
			}
			if !assetOfInterest(b.Asset) {
				err := fmt.Sprintf("ignored asset: '%s'", b.Asset)
				log.Warn(err)
			} else {
				brr := exa.Balance{
					Asset:   b.Asset,
					Balance: b.Balance.String(),
					Account: &account,
				}
				resp.Balances = append(resp.Balances, &brr)
			}
		}
	}
	return &resp, nil
}

func main() {
	flag.Parse()
	version = fmt.Sprintf("%s::%s", bts, rev)
	fmt.Printf("huobi: %s\n", version)
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

func (s *server) GetOpenOrders(ctx context.Context, in *exa.GetOpenOrdersRequest) (*exa.GetOpenOrdersResponse, error) {
	log.Printf("exa GetOpenOrders request: %v -- %v", in.GetRequestId(), in.GetRequestTime().AsTime())
	log.Printf("exa GetOpenOrders request: exchange: %s", in.GetExchange())
	re := in.GetExchange()
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

	odata, err := huobi.GetOpenOrders(apiKey, apiSecret)
	if err != nil {
		err := status.Error(codes.Internal, err.Error())
		return nil, err
	}

	resp := exa.GetOpenOrdersResponse{
		ResponseTime: timestamppb.Now(),
		RequestId:    in.GetRequestId(),
	}
	for _, o := range odata {
		oo := exa.Order{
			Symbol:        o.Symbol,
			Source:        o.Source,
			Price:         o.Price.String(),
			Amount:        o.Amount.String(),
			AccountId:     strconv.Itoa(int(o.AccountId)),
			ClientOrderId: o.ClientOrderId,
			FilledAmount:  o.FilledAmount.String(),
			FilledFees:    o.FilledFees.String(),
			Id:            strconv.Itoa(int(o.Id)),
			State:         o.State,
			Type:          o.Type,
		}
		oo.CreatedAt = timestamppb.New(time.Unix(int64(o.CreatedAt/1e3), 0))
		resp.Orders = append(resp.Orders, &oo)
	}
	return &resp, nil
}

func (s *server) CancelOrders(ctx context.Context, in *exa.CancelOrdersRequest) (*exa.CancelOrdersResponse, error) {
	log.Printf("exa CancelOrders request: %v -- %v", in.GetRequestId(), in.GetRequestTime().AsTime())
	log.Printf("exa CancelOrders request: exchange: %s", in.GetExchange())
	re := in.GetExchange()
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

	cd, err := huobi.CancelOrders(apiKey, apiSecret, in.OrderIds)
	if err != nil {
		err := status.Error(codes.Internal, err.Error())
		return nil, err
	}

	resp := exa.CancelOrdersResponse{
		ResponseTime: timestamppb.Now(),
		RequestId:    in.GetRequestId(),
		Succeeded:    cd.Succeeded,
	}
	if len(cd.Failed) > 0 {
		var cfs []*exa.CancelFailure
		for _, f := range cd.Failed {
			oid := f.OrderId
			if oid == "" && f.ClientOrderId != "" {
				oid = f.ClientOrderId
			}
			ecf := exa.CancelFailure{
				OrderId:      oid,
				ErrorMessage: f.ErrorMessage,
			}
			cfs = append(cfs, &ecf)
		}
		resp.Failed = cfs
	}
	return &resp, nil
}

func (s *server) PlaceOrder(ctx context.Context, in *exa.PlaceOrderRequest) (*exa.PlaceOrderResponse, error) {
	log.Printf("exa PlaceOrder request: %v -- %v", in.GetRequestId(), in.GetRequestTime().AsTime())
	log.Printf("exa PlaceOrder request: exchange: %s", in.GetExchange())
	re := in.GetExchange()
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

	oid, err := huobi.PlaceOrder(apiKey, apiSecret, in.AccountId, in.Symbol, in.Type, in.Amount, in.Price, in.ClientOrderId)
	if err != nil {
		err := status.Error(codes.Internal, err.Error())
		return nil, err
	}

	resp := exa.PlaceOrderResponse{
		ResponseTime: timestamppb.Now(),
		RequestId:    in.GetRequestId(),
		OrderId:      oid,
	}
	return &resp, nil
}
