package main

import (
	"context"
	"flag"
	"log"
	"time"

	exa "github.com/alphabot-fi/T-801/internal/proto/exa"
	monitor "github.com/alphabot-fi/T-801/internal/proto/monitor"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	addr = flag.String("addr", "localhost:50051", "address to connect to")
	tef  = flag.Bool("tef", false, "test EXA functions [false]")
)

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	testMonitor(conn)
	if tef != nil && *tef == true {
		testEXA(conn)
	}
}

func testMonitor(c *grpc.ClientConn) {
	cmon := monitor.NewMonitorClient(c)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r1, err := cmon.Ping(ctx, &monitor.PingRequest{RequestTime: timestamppb.Now()})
	if err != nil {
		log.Printf("could not ping: %v\n", err)
	}
	log.Printf("response time: %v\n", r1.GetResponseTime().AsTime())
	log.Printf("version: %s\n", r1.GetVersion())
}

func testEXA(c *grpc.ClientConn) {
	cexa := exa.NewEXAClient(c)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r2, err := cexa.GetBalance(ctx, &exa.GetBalanceRequest{RequestTime: timestamppb.Now(), RequestId: "1", Exchange: exa.ExchangeType_KRAKEN, ApiKey: "key", ApiSecret: "secret"})
	if err != nil {
		log.Printf("failed to get balance 1: %v\n", err)
	} else {
		log.Printf("balance 1: response time: %v\n", r2.GetResponseTime().AsTime())
	}

	r3, err := cexa.GetBalance(ctx, &exa.GetBalanceRequest{RequestTime: timestamppb.Now(), RequestId: "2", ApiKey: "key", ApiSecret: "secret", Exchange: exa.ExchangeType_KUCOIN})
	if err != nil {
		log.Printf("failed to get balance 2: %v\n", err)
	} else {
		log.Printf("balance 2: response time: %v\n", r3.GetResponseTime().AsTime())
	}

	r4, err := cexa.GetBalance(ctx, &exa.GetBalanceRequest{RequestTime: timestamppb.Now(), RequestId: "3", ApiKey: "key", ApiSecret: "secret", Exchange: exa.ExchangeType_FTX})
	if err != nil {
		log.Printf("failed to get balance 3: %v\n", err)
	} else {
		log.Printf("balance 3: response time: %v\n", r4.GetResponseTime().AsTime())
	}
}
