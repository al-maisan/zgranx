package huobi

import "github.com/shopspring/decimal"

type Balance struct {
	Asset   string
	Balance decimal.Decimal
}

func GetBalance(apiKey, apiSecret, asset string) ([]Balance, error) {
	return nil, nil
}
