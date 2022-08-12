package ma

import (
	"errors"
	"fmt"

	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

// E computes the exponential moving average (EMA) for the prices passed.
// The last price in the array is the most recent one.
func EMA(prices []string) (string, error) {
	if len(prices) == 0 {
		return "", errors.New("empty price array")
	}
	var (
		// c = current price
		// p = previous period's EMA (a SMA is used for the first period's
		//     calculations)
		// k = exponential smoothing constant
		//
		// EMA = (k * (c - p)) + p
		zero, k, p, ema decimal.Decimal
		sma             string
		err             error
	)

	sma, err = SMA(prices)
	if err != nil {
		return "", err
	}
	p, err = decimal.NewFromString(sma)
	if err != nil {
		err = fmt.Errorf("invalid SMA value: '%s', %v", sma, err)
		log.Error(err)
		return "", err
	}
	ema = p

	zero = decimal.NewFromFloat(0.0)
	k = decimal.NewFromFloat(2.0).Div(decimal.NewFromInt(int64(len(prices) + 1)))

	for pi := len(prices) - 2; pi >= 0; pi-- {
		c, err := decimal.NewFromString(prices[pi])
		if err != nil {
			err = fmt.Errorf("invalid price value: '%s', %v", prices[pi], err)
			log.Error(err)
			return "", err
		}
		if c.LessThanOrEqual(zero) {
			err = fmt.Errorf("price value out of range: '%s', %v", prices[pi], err)
			log.Error(err)
			return "", err
		}
		ema = c.Sub(p).Mul(k).Add(p)
		p = ema
	}

	return ema.StringFixed(6), nil
}
