package main

import (
	"context"
	"fmt"
)

// PriceFetcher é uma interface que pode buscar um preço
type PriceFetcher interface {
	FetchPrice(context.Context, string) (float64, error)
}

// priceFetcher está implementando a **interface** PriceFetcher através da função FetchPrice
type priceFetcher struct{}

func (s *priceFetcher) FetchPrice(ctx context.Context, ticker string) (float64, error) {
	return MockPriceTicker(ctx, ticker)
}

var priceMocks = map[string]float64{
	"BTC": 20_000.0,
	"ETH": 200.0,
	"VRS": 200.0,
}

func MockPriceTicker(ctx context.Context, ticker string) (float64, error) {
	price, ok := priceMocks[ticker]
	if !ok {
		return price, fmt.Errorf("O ticker dado (%s) não é suportado")
	}

	return price, nil
}
