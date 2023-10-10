package main

import (
	"context"
	"fmt"
)

var priceMocks = map[string]float64{
	"BTC": 20_000.0,
	"ETH": 200.0,
	"VRS": 200.0,
}

// PriceFetcher é uma interface que pode buscar um preço
type PriceService interface {
	FetchPrice(context.Context, string) (float64, error)
}

// priceFetcher está implementando a **interface** PriceFetcher através da função FetchPrice
type priceService struct{}

func (s *priceService) FetchPrice(ctx context.Context, ticker string) (float64, error) {
	return MockPriceTicker(ctx, ticker)
}

func MockPriceTicker(_ context.Context, ticker string) (float64, error) {
	// resp := http.Get("link api para fazer fetch do preço")	 get request de uma api
	price, ok := priceMocks[ticker]
	if !ok {
		return 0.0, fmt.Errorf("O ticker dado %s não é suportado", ticker)
	}

	return price, nil
}
