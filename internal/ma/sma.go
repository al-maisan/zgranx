package ma

import (
	"errors"
	"fmt"

	"github.com/shopspring/decimal"
)

func SMA(prices []string) (string, error) {
	var err error
	if len(prices) == 0 {
		err = errors.New("empty price array")
		log.Error(err)
		return "", err
	}
	var pvs []decimal.Decimal
	zero := decimal.NewFromFloat(0.0)

	for _, p := range prices {
		pv, err := decimal.NewFromString(p)
		if err != nil {
			err = fmt.Errorf("invalid price value: '%s', %v", p, err)
			log.Error(err)
			return "", err
		}
		if pv.LessThanOrEqual(zero) {
			err = fmt.Errorf("price value out of range: '%s'", p)
			log.Error(err)
			return "", err
		}
		pvs = append(pvs, pv)
	}
	avg := decimal.Avg(pvs[0], pvs[1:]...)
	return avg.StringFixed(6), nil
}
