package main

import (
	"flag"
)

func main() {
	// client := client.New("http://localhost:4000")

	// price, err := client.FetchPrice(context.Background(), "ET")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("%+v\n", price)
	listenAddr := flag.String("listenaddr", ":3030", "server está rodando na porta 3030")
	flag.Parse()

	svc := NewLoggingService(NewMetricService(&priceFetcher{}))

	server := NewJSONApiServer(*listenAddr, svc)
	server.Run()

}
