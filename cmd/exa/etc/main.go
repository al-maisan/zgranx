package main

import (
	"context"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/alphabot-fi/T-801/internal/proto/exa"
	"github.com/alphabot-fi/T-801/internal/proto/monitor"
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

	var ids cli.StringSlice
	var exchange, saddr, accountId, symbol, otype, amount, price, clientOrderId string
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
					log.Info("get-open-orders, saddr = ", saddr)
					conn, err := grpc.Dial(saddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
					if err != nil {
						log.Fatalf("did not connect: %v", err)
					}
					defer conn.Close()
					getOrders(conn, exchange)
					return nil
				},
			},
			{
				Name:  "ping",
				Usage: "ping the exchange adapter for the given exchange",
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
					log.Info("ping, saddr = ", saddr)
					conn, err := grpc.Dial(saddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
					if err != nil {
						log.Fatalf("did not connect: %v", err)
					}
					defer conn.Close()
					ping(conn, exchange)
					return nil
				},
			},
			{
				Name:    "cancel-orders",
				Aliases: []string{"cos"},
				Usage:   "cancel the specified orders on the given exchange",
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
					&cli.StringSliceFlag{
						Name:        "ids",
						Usage:       "1+ order id",
						Aliases:     []string{"i"},
						Required:    true,
						Destination: &ids,
					},
				},
				Action: func(c *cli.Context) error {
					if saddr == "" {
						saddr = defaultAddr
					}
					log.Info("cancel-orders, saddr = ", saddr)
					conn, err := grpc.Dial(saddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
					if err != nil {
						log.Fatalf("did not connect: %v", err)
					}
					defer conn.Close()
					cancelOrders(conn, exchange, ids.Value())
					return nil
				},
			},
			{
				Name:    "place-order",
				Aliases: []string{"po"},
				Usage:   "place an order on the given exchange",
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
					&cli.StringFlag{
						Name:        "account-id",
						Usage:       "account identifier",
						Required:    true,
						Destination: &accountId,
					},
					&cli.StringFlag{
						Name:        "symbol",
						Usage:       "trading pair (e.g. \"btcusdt\")",
						Required:    true,
						Destination: &symbol,
					},
					&cli.StringFlag{
						Name:        "type",
						Usage:       "order type, either 'buy-limit' or 'sell-limit'",
						Required:    true,
						Destination: &otype,
					},
					&cli.StringFlag{
						Name:        "amount",
						Usage:       "how much of the base asset you wish to buy/sell",
						Required:    true,
						Destination: &amount,
					},
					&cli.StringFlag{
						Name:        "price",
						Usage:       "at what price (quote asset) you wish to buy/sell",
						Required:    true,
						Destination: &price,
					},
					&cli.StringFlag{
						Name:        "cid",
						Usage:       "the id you wish to assign to the new order",
						Required:    true,
						Destination: &clientOrderId,
					},
				},
				Action: func(c *cli.Context) error {
					if saddr == "" {
						saddr = defaultAddr
					}
					log.Info("place-order, saddr = ", saddr)
					if otype != "buy-limit" && otype != "sell-limit" {
						err := fmt.Errorf("invalid order type: '%s', must be either 'buy-limit' or 'sell-limit'", otype)
						return err
					}
					conn, err := grpc.Dial(saddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
					if err != nil {
						log.Fatalf("did not connect: %v", err)
					}
					defer conn.Close()
					placeOrder(conn, exchange, accountId, symbol, otype, amount, price, clientOrderId)
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

func pair2string(p *exa.Pair) string {
	if p == nil {
		return ""
	}
	return strings.ToLower(p.Base.String()) + "-" + strings.ToLower(p.Quote.String())
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
	for i, o := range res.Orders {
		if o == nil {
			log.Errorf("nil order for index %d", i)
			continue
		}
		var price string
		if o.Price != nil {
			price = fmt.Sprintf("@ %8s", *o.Price)
		}
		fmt.Printf("[%2d] %12s | %6s |%4s | %8s %11s | %16s | %8s | %10s | %12s | %s\n", i, pair2string(o.Pair), o.Type, o.Side, o.Amount, price, o.Id, o.AccountId, o.State, o.ClientOrderId, o.CreatedAt.AsTime().UTC())
	}
}

func apiKeys(exchange, access string) [2]string {
	apiKey := os.Getenv(fmt.Sprintf("%s_%s_ACCESS_KEY", strings.ToUpper(exchange), strings.ToUpper(access)))
	secretKey := os.Getenv(fmt.Sprintf("%s_%s_SECRET_KEY", strings.ToUpper(exchange), strings.ToUpper(access)))
	return [2]string{apiKey, secretKey}
}

func ping(c *grpc.ClientConn, exchange string) {
	cmon := monitor.NewMonitorClient(c)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	req := monitor.PingRequest{
		RequestTime: timestamppb.Now(),
	}
	log.Infof("ping for %s at %v", exchange, time.Now().UTC())
	res, err := cmon.Ping(ctx, &req)
	if err != nil {
		log.Errorf("failed to get ping EXA for %s, %v", exchange, err)
		return
	}
	log.Infof("ping response from %s EXA at %v, version('%s')", exchange, res.ResponseTime.AsTime().UTC(), res.Version)
}

func cancelOrders(c *grpc.ClientConn, exchange string, ids []string) {
	aks := apiKeys(exchange, "rw")
	cexa := exa.NewEXAClient(c)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	req := exa.CancelOrdersRequest{
		RequestTime: timestamppb.Now(),
		RequestId:   fmt.Sprintf("%d", time.Now().UTC().Unix()),
		Exchange:    "huobi",
		ApiKey:      aks[0],
		ApiSecret:   aks[1],
		OrderIds:    ids,
	}
	log.Infof("cancel orders on %s at %v (%s)", exchange, time.Now().UTC(), req.GetRequestId())
	res, err := cexa.CancelOrders(ctx, &req)
	if err != nil {
		log.Errorf("failed to cancel orders on %s, %v", exchange, err)
		return
	}
	log.Infof("canceled orders on %s at %v -- succeeded: %d, failed: %d", exchange, res.ResponseTime.AsTime().UTC(), len(res.Succeeded), len(res.Failed))
	if len(res.Succeeded) > 0 {
		sort.Strings(res.Succeeded)
		fmt.Println(" >> succeeded:")
		for _, s := range res.Succeeded {
			fmt.Printf("    - %s\n", s)
		}
	}
	if len(res.Failed) > 0 {
		fmt.Println(" ** failed:")
		for _, f := range res.Failed {
			fmt.Printf("    ! %s, %s\n", f.OrderId, f.ErrorMessage)
		}
	}
}

func getPair(sym string) (*exa.Pair, error) {
	ss := strings.Split(sym, "-")
	if len(ss) != 2 {
		err := fmt.Errorf("invalid pair: '%s'", sym)
		return nil, err
	}
	b, q := ss[0], ss[1]
	ba, ok := exa.Asset_value[b]
	if !ok {
		err := fmt.Errorf("unknown base asset: '%s'", b)
		return nil, err
	}
	qa, ok := exa.Asset_value[q]
	if !ok {
		err := fmt.Errorf("unknown quote asset: '%s'", q)
		return nil, err
	}
	return &exa.Pair{Base: exa.Asset(ba), Quote: exa.Asset(qa)}, nil
}

func getTypeAndSide(otype string) (exa.OrderType, exa.Side, error) {
	switch otype {
	case "sell-limit":
		return exa.OrderType_LIMIT, exa.Side_SELL, nil
	case "buy-limit":
		return exa.OrderType_LIMIT, exa.Side_BUY, nil
	}
	return 0, 0, fmt.Errorf("invalid order type: '%s'", otype)
}

func placeOrder(c *grpc.ClientConn, exchange, accountId, symbol, otype, amount, price, clientOrderId string) error {
	aks := apiKeys(exchange, "rw")
	cexa := exa.NewEXAClient(c)

	pair, err := getPair(symbol)
	if err != nil {
		log.Error(err)
		return err
	}
	ot, os, err := getTypeAndSide(otype)
	if err != nil {
		log.Error(err)
		return err
	}
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	req := exa.PlaceOrderRequest{
		RequestTime:   timestamppb.Now(),
		RequestId:     fmt.Sprintf("%d", time.Now().UTC().Unix()),
		Exchange:      "huobi",
		ApiKey:        aks[0],
		ApiSecret:     aks[1],
		AccountId:     accountId,
		Pair:          pair,
		Type:          ot,
		Side:          os,
		Amount:        amount,
		Price:         price,
		ClientOrderId: clientOrderId,
	}
	log.Infof("place order on %s at %v (%s)", exchange, time.Now().UTC(), req.GetRequestId())
	res, err := cexa.PlaceOrder(ctx, &req)
	if err != nil {
		log.Errorf("failed to place order on %s, %v", exchange, err)
		return err
	}
	log.Infof("placed order on %s at %v", exchange, res.ResponseTime.AsTime().UTC())
	log.Infof("new order id: %s", res.OrderId)
	return nil
}
