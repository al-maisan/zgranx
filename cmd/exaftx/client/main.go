package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/al-maisan/zgranx/internal/proto/monitor"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMonitorClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Ping(ctx, &pb.PingRequest{RequestTime: timestamppb.Now()})
	if err != nil {
		log.Fatalf("could not ping: %v", err)
	}
	log.Printf("response time: %v\n", r.GetResponseTime().AsTime())
	log.Printf("version: %s\n", r.GetVersion())
}
