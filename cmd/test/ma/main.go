package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/alphabot-fi/T-801/internal/proto/ma"
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

	var prices cli.StringSlice
	var saddr string
	app := &cli.App{
		Name:  "mtc",
		Usage: "ma test client",
		Commands: []*cli.Command{
			{
				Name:  "ping",
				Usage: "ping the moving average service",
				Flags: []cli.Flag{
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
					ping(conn)
					return nil
				},
			},
			{
				Name:  "ema",
				Usage: "calculate exponential moving average",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "server-address",
						Usage:       "server address (e.g. \"localhost:50051\")",
						Required:    false,
						Destination: &saddr,
					},
					&cli.StringSliceFlag{
						Name:        "prices",
						Usage:       "1+ price",
						Aliases:     []string{"p"},
						Required:    true,
						Destination: &prices,
					},
				},
				Action: func(c *cli.Context) error {
					if saddr == "" {
						saddr = defaultAddr
					}
					log.Info("EMA, saddr = ", saddr)
					conn, err := grpc.Dial(saddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
					if err != nil {
						log.Fatalf("did not connect: %v", err)
					}
					defer conn.Close()
					cma(conn, prices.Value(), false)
					return nil
				},
			},
			{
				Name:  "sma",
				Usage: "calculate simple moving average",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "server-address",
						Usage:       "server address (e.g. \"localhost:50051\")",
						Required:    false,
						Destination: &saddr,
					},
					&cli.StringSliceFlag{
						Name:        "prices",
						Usage:       "1+ price",
						Aliases:     []string{"p"},
						Required:    true,
						Destination: &prices,
					},
				},
				Action: func(c *cli.Context) error {
					if saddr == "" {
						saddr = defaultAddr
					}
					log.Info("SMA, saddr = ", saddr)
					conn, err := grpc.Dial(saddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
					if err != nil {
						log.Fatalf("did not connect: %v", err)
					}
					defer conn.Close()
					cma(conn, prices.Value(), true)
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

func ping(c *grpc.ClientConn) {
	cmon := monitor.NewMonitorClient(c)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	req := monitor.PingRequest{
		RequestTime: timestamppb.Now(),
	}
	log.Infof("ping at %v", time.Now().UTC())
	res, err := cmon.Ping(ctx, &req)
	if err != nil {
		log.Errorf("failed to ping moving average service, %v", err)
		return
	}
	log.Infof("ping response from moving average service at %v, version('%s')", res.ResponseTime.AsTime().UTC(), res.Version)
}

func cma(c *grpc.ClientConn, prices []string, simple bool) {
	client := ma.NewMAClient(c)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	req := ma.MARequest{
		RequestTime: timestamppb.Now(),
		RequestId:   fmt.Sprintf("%d", time.Now().UTC().Unix()),
		Period:      uint32(len(prices)),
		Prices:      prices,
	}
	var (
		res *ma.MAResponse
		err error
	)
	if simple {
		log.Infof("SMA at %v (%s)", time.Now().UTC(), req.GetRequestId())
		res, err = client.S(ctx, &req)
	} else {
		log.Infof("EMA at %v (%s)", time.Now().UTC(), req.GetRequestId())
		res, err = client.E(ctx, &req)
	}
	if err != nil {
		log.Errorf("failed to obtain SMA, %v", err)
		return
	}
	log.Infof("SMA value obtained at %v -- %s", res.ResponseTime.AsTime().UTC(), res.Result)
}
