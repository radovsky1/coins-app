package util

const (
	BTC  = "BTC"
	ETH  = "ETH"
	BNB  = "BNB"
	LTC  = "LTC"
	USDT = "USDT"
)

func IsSupportedCoin(coin string) bool {
	switch coin {
	case BTC, ETH, BNB, LTC, USDT:
		return true
	default:
		return false
	}
}
