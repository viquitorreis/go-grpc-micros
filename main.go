package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"gitlab.com/victorreisprog/go-grpc-micros/client"
	"gitlab.com/victorreisprog/go-grpc-micros/proto"
)

func main() {
	var (
		jsonAddr = flag.String("json", ":3030", "server transporter JSON está rodando na porta 3030")
		grpcAddr = flag.String("grpc", ":4040", "server transporter GRPC está rodando na porta 4040")
		svc      = NewLoggingService(NewMetricService(&priceService{}))
		ctx      = context.Background()
	)
	flag.Parse()

	grpcClient, err := client.NewGRPCClient(":4040")
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			time.Sleep(3 * time.Second)
			resp, err := grpcClient.FetchPrice(ctx, &proto.PriceRequest{Ticker: "BTC"})
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("%+v\n", resp)
		}
	}()

	go makeGRPCServerAndRun(*grpcAddr, svc)

	jsonServer := NewJSONApiServer(*jsonAddr, svc)
	jsonServer.Run()

}
