package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"gitlab.com/victorreisprog/go-grpc-micros/client"
)

func main() {
	client := client.New("http://localhost:4000")

	price, err := client.FetchPrice(context.Background(), "ET")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", price)
	return
	listenAddr := flag.String("listenaddr", ":4000", "server está rodando na porta 4000")
	flag.Parse()

	svc := NewLoggingService(NewMetricService(&priceFetcher{}))

	server := NewJSONApiServer(*listenAddr, svc)
	server.Run()

}
