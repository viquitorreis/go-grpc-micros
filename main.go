package main

import (
	"context"
	"fmt"
	"log"
)

func main() {
	svg := NewLoggingService(&priceFetcher{})
	price, err := svg.FetchPrice(context.Background(), "ETH")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(price)

}
