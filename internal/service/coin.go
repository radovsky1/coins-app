package service

import (
	"coins-app/internal/core"
	"coins-app/internal/service/webapi"
)

type CoinService struct {
	webapi *webapi.BinanceWebAPI
}

func NewCoinService(webapi *webapi.BinanceWebAPI) *CoinService {
	return &CoinService{
		webapi: webapi,
	}
}

func (c *CoinService) GetCoinPrices(symbol string) ([]core.SymbolPrice, error) {
	return c.webapi.GetPrices(symbol)
}
