package main

import (
	"context"
)

type metricsService struct {
	next PriceFetcher
}

func NewMetricService(next PriceFetcher) PriceFetcher {
	return &metricsService{
		next: next,
	}
}

func (s *metricsService) FetchPrice(ctx context.Context, ticker string) (float64, error) {
	// fmt.Println("Pushing metrics to prometheus")
	// lógica das métricas -- push  no prometheus
	return s.next.FetchPrice(ctx, ticker)
}
