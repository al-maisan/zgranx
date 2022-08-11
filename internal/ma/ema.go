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
		zero, wm, pema, ema decimal.Decimal
		sma                 string
		err                 error
	)

	sma, err = SMA(prices)
	if err != nil {
		return "", err
	}
	pema, err = decimal.NewFromString(sma)
	if err != nil {
		return "", fmt.Errorf("invalid SMA value: '%s'", sma)
	}
	ema = pema

	zero = decimal.NewFromFloat(0.0)
	wm = decimal.NewFromFloat(2.0).Div(decimal.NewFromInt(int64(len(prices) + 1)))
	log.Infof("1>>   wm = %s", wm.String())

	for pi := len(prices) - 2; pi >= 0; pi-- {
		log.Infof("2>> pema = %s", pema.String())
		pv, err := decimal.NewFromString(prices[pi])
		if err != nil {
			return "", fmt.Errorf("invalid price value: '%s'", prices[pi])
		}
		if pv.LessThanOrEqual(zero) {
			return "", fmt.Errorf("price value out of range: '%s'", prices[pi])
		}
		log.Infof("3>> cval = %s", pv.String())
		ema = pv.Sub(pema).Mul(wm).Add(pema)
		log.Infof("4>>  ema = %s", ema.String())
		pema = ema
	}

	return ema.StringFixed(6), nil
}
