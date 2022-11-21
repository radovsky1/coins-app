package webapi

import (
	"coins-app/internal/core"
	"context"
	"github.com/adshao/go-binance/v2"
	"strings"
)

type BinanceWebAPIConfig struct {
	APIKey    string
	APISecret string
}

type BinanceWebAPI struct {
	config BinanceWebAPIConfig
	client *binance.Client
}

func NewBinanceWebAPI(config BinanceWebAPIConfig) *BinanceWebAPI {
	return &BinanceWebAPI{
		config: config,
		client: binance.NewClient(config.APIKey, config.APISecret),
	}
}

func (b *BinanceWebAPI) GetPrices(coin string) ([]core.SymbolPrice, error) {
	prices, err := b.client.NewListPricesService().Do(context.Background())
	if err != nil {
		return nil, err
	}
	coinPrices := make([]core.SymbolPrice, 0)
	for _, price := range prices {
		if strings.HasPrefix(price.Symbol, coin) {
			coinPrices = append(coinPrices, core.SymbolPrice{
				Symbol: price.Symbol,
				Price:  price.Price,
			})
		}
	}
	return coinPrices, nil
}
