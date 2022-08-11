package ma

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func SMA(prices []string) (string, error) {
	var pvs []decimal.Decimal
	zero := decimal.NewFromFloat(0.0)

	for _, p := range prices {
		pv, err := decimal.NewFromString(p)
		if err != nil {
			return "", fmt.Errorf("invalid price value: '%s'", p)
		}
		if pv.LessThanOrEqual(zero) {
			return "", fmt.Errorf("price value out of range: '%s'", p)
		}
		pvs = append(pvs, pv)
	}
	avg := decimal.Avg(pvs[0], pvs[1:]...)
	return avg.StringFixed(6), nil
}
