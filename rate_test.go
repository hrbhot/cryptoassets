package main

import (
	"context"
	"github.com/piquette/finance-go"
	"testing"
	"time"
)

func TestGetBTCUSDCNY(t *testing.T) {
	finance.NewBackends(newHttpClient("socks5://127.0.0.1:1080"))
	t.Log(GetBTCUSDCNYFromYahoo())
}

func TestGetUSDTUSD(t *testing.T) {
	t.Log(GetUSDTUSDFromCoinMarketCap())
}

func TestStartFetchRate(t *testing.T) {
	StartFetchRate(context.Background())
	time.Sleep(time.Second * 5)
}

func TestGetUSDCNYFromYahoo(t *testing.T) {
	t.Log(GetUSDCNYFromYahoo())
}

func TestGetUSDTUSDFromCoinMarketCap(t *testing.T) {
	t.Log(GetUSDTUSDFromCoinMarketCap())
}

func TestGetBTCUSDFromCoinMarketCap(t *testing.T) {
	t.Log(GetBTCUSDFromCoinMarketCap())
}
