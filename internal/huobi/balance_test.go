package huobi

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

const btd = `
{
    "status": "ok",
    "data": {
        "id": 1000001,
        "type": "spot",
        "state": "working",
        "list": [
            {
                "currency": "usdt",
                "type": "trade",
                "balance": "91.850043797676510303",
                "seq-num": "477"
            },
            {
                "currency": "usdt",
                "type": "frozen",
                "balance": "5.160000000000000015",
                "seq-num": "477"
            },
            {
                "currency": "poly",
                "type": "trade",
                "balance": "147.928994082840236",
                "seq-num": "2"
            }
        ]
    }
}
`

func TestParseBalances(t *testing.T) {
	b1, _ := decimal.NewFromString("91.850043797676510303")
	b2, _ := decimal.NewFromString("5.160000000000000015")
	b3, _ := decimal.NewFromString("147.928994082840236")
	expected := BalanceData{
		Account: 1000001,
		Type:    "spot",
		State:   "working",
		Balances: []Balance{
			{
				Asset:   "usdt",
				Type:    "trade",
				Balance: b1,
			},
			{
				Asset:   "usdt",
				Type:    "frozen",
				Balance: b2,
			},
			{
				Asset:   "poly",
				Type:    "trade",
				Balance: b3,
			},
		},
	}

	actual, err := parseBalances([]byte(btd))
	assert.Nil(t, err)
	assert.NotNil(t, actual)
	assert.Equal(t, expected, *actual)
}
