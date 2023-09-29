package main

import (
	"context"
	"encoding/json"
	"math/rand"
	"net/http"

	"gitlab.com/victorreisprog/go-grpc-micros/types"
)

type APIFunc func(context.Context, http.ResponseWriter, *http.Request) error

type JSONApiServer struct {
	listenAddr string
	svc        PriceFetcher
}

func NewJSONApiServer(listenAddr string, svc PriceFetcher) *JSONApiServer {
	return &JSONApiServer{
		listenAddr: listenAddr,
		svc:        svc,
	}
}

func (s *JSONApiServer) Run() {
	http.HandleFunc("/", makeHTTPHandlerFunc(s.handleFetchPrice))

	http.ListenAndServe(s.listenAddr, nil)
}

func makeHTTPHandlerFunc(apiFn APIFunc) http.HandlerFunc {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "requestID", rand.Intn(100000000))

	return func(w http.ResponseWriter, r *http.Request) {
		if err := apiFn(context.Background(), w, r); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		}

	}
}

func (s *JSONApiServer) handleFetchPrice(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	ticker := r.URL.Query().Get("ticker")

	price, err := s.svc.FetchPrice(ctx, ticker)
	if err != nil {
		return err
	}

	priceResp := types.PriceResponse{
		Price:  price,
		Ticker: ticker,
	}

	return writeJSON(w, http.StatusOK, &priceResp)

}

func writeJSON(w http.ResponseWriter, s int, v any) error {
	w.WriteHeader(s) // retornando o status da req
	return json.NewEncoder(w).Encode(v)

}
