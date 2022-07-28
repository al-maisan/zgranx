package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/alphabot-fi/T-801/internal/proto/exa"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const defaultAddr = "localhost:50051"

var (
	bts, rev, version string
	log               = logrus.New()
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Error("Error loading .env file")
	}
	version = fmt.Sprintf("%s::%s", bts, rev)
	log.Info("version = ", version)

	var exchange, saddr string
	app := &cli.App{
		Name:  "etc",
		Usage: "exa test client",
		Commands: []*cli.Command{
			{
				Name:    "get-balances",
				Aliases: []string{"gb"},
				Usage:   "get the balances for the given exchange",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "exchange",
						Usage:       "exchange to use (e.g. \"huobi\")",
						Required:    true,
						Destination: &exchange,
					},
					&cli.StringFlag{
						Name:        "server-address",
						Usage:       "server address (e.g. \"localhost:50051\")",
						Required:    false,
						Destination: &saddr,
					},
				},
				Action: func(c *cli.Context) error {
					if saddr == "" {
						saddr = defaultAddr
					}
					log.Info("get-balances, saddr = ", saddr)
					conn, err := grpc.Dial(saddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
					if err != nil {
						log.Fatalf("did not connect: %v", err)
					}
					defer conn.Close()
					getBalances(conn, exchange)
					return nil
				},
			},
			{
				Name:    "get-open-orders",
				Aliases: []string{"goo"},
				Usage:   "get the open orders for the given exchange",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "exchange",
						Usage:       "exchange to use (e.g. \"huobi\")",
						Required:    true,
						Destination: &exchange,
					},
					&cli.StringFlag{
						Name:        "server-address",
						Usage:       "server address (e.g. \"localhost:50051\")",
						Required:    false,
						Destination: &saddr,
					},
				},
				Action: func(c *cli.Context) error {
					if saddr == "" {
						saddr = defaultAddr
					}
					log.Info("get-balances, saddr = ", saddr)
					conn, err := grpc.Dial(saddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
					if err != nil {
						log.Fatalf("did not connect: %v", err)
					}
					defer conn.Close()
					getOrders(conn, exchange)
					return nil
				},
			},
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func getBalances(c *grpc.ClientConn, exchange string) {
	aks := apiKeys(exchange, "ro")
	cexa := exa.NewEXAClient(c)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	req := exa.GetBalancesRequest{
		RequestTime: timestamppb.Now(),
		RequestId:   fmt.Sprintf("%d", time.Now().UTC().Unix()),
		Exchange:    "huobi",
		ApiKey:      aks[0],
		ApiSecret:   aks[1],
	}
	log.Infof("get balances for %s at %v (%s)", exchange, time.Now().UTC(), req.GetRequestId())
	res, err := cexa.GetBalances(ctx, &req)
	if err != nil {
		log.Errorf("failed to get balances for %s, %v", exchange, err)
		return
	}
	log.Infof("got balances for %s at %v", exchange, res.ResponseTime.AsTime().UTC())
	for i, b := range res.Balances {
		if b == nil {
			log.Errorf("nil balance for index %d", i)
			continue
		}
		fmt.Printf("[%2d] %10s %6s %20s\n", i, b.GetAccount(), b.GetAsset(), b.GetBalance())
	}
}

func getOrders(c *grpc.ClientConn, exchange string) {
	aks := apiKeys(exchange, "ro")
	cexa := exa.NewEXAClient(c)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	req := exa.GetOpenOrdersRequest{
		RequestTime: timestamppb.Now(),
		RequestId:   fmt.Sprintf("%d", time.Now().UTC().Unix()),
		Exchange:    "huobi",
		ApiKey:      aks[0],
		ApiSecret:   aks[1],
	}
	log.Infof("get orders for %s at %v (%s)", exchange, time.Now().UTC(), req.GetRequestId())
	res, err := cexa.GetOpenOrders(ctx, &req)
	if err != nil {
		log.Errorf("failed to get orders for %s, %v", exchange, err)
		return
	}
	log.Infof("got orders for %s at %v", exchange, res.ResponseTime.AsTime().UTC())
	for i, b := range res.Orders {
		if b == nil {
			log.Errorf("nil order for index %d", i)
			continue
		}
		fmt.Printf("[%2d] %10s %8s %6\n", i, b.AccountId, b.Symbol, b.Price)
	}
}

func apiKeys(exchange, access string) [2]string {
	apiKey := os.Getenv(fmt.Sprintf("%s_%s_ACCESS_KEY", strings.ToUpper(exchange), strings.ToUpper(access)))
	secretKey := os.Getenv(fmt.Sprintf("%s_%s_SECRET_KEY", strings.ToUpper(exchange), strings.ToUpper(access)))
	return [2]string{apiKey, secretKey}
}
