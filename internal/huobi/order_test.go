package huobi

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

const otd = `
{
    "status": "ok",
    "data": [
        {
            "symbol": "apnusdt",
            "source": "web",
            "price": "1.555550000000000000",
            "created-at": 1630633835224,
            "amount": "572.330000000000000000",
            "account-id": 13496526,
            "filled-cash-amount": "0.0",
            "client-order-id": "abc-123-xyz",
            "filled-amount": "0.0",
            "filled-fees": "0.0",
            "id": 357630527817871,
            "state": "submitted",
            "type": "sell-limit"
        }
    ]
}
`

func TestParseOpenOrders(t *testing.T) {
	expected := []Order{
		{
			Symbol:        "apnusdt",
			Source:        "web",
			Price:         decimal.New(1555550000000000000, -18),
			CreatedAt:     1630633835224,
			Amount:        decimal.New(57233, -2),
			AccountId:     13496526,
			ClientOrderId: "abc-123-xyz",
			FilledAmount:  decimal.NewFromFloat(0.0),
			FilledFees:    decimal.NewFromFloat(0.0),
			Id:            357630527817871,
			State:         "submitted",
			Type:          "sell-limit",
		},
	}

	actual, err := parseOpenOrders([]byte(otd))
	assert.Nil(t, err)
	assert.NotNil(t, actual)
	ad := actual[0]
	ed := expected[0]
	assert.Equal(t, ed.Symbol, ad.Symbol)
	assert.Equal(t, ed.Source, ad.Source)
	assert.True(t, ad.Price.Equal(ed.Price))
	assert.Equal(t, ed.CreatedAt, ad.CreatedAt)
	assert.True(t, ad.Amount.Equal(ed.Amount))
	assert.Equal(t, ed.AccountId, ad.AccountId)
	assert.Equal(t, ed.ClientOrderId, ad.ClientOrderId)
	assert.True(t, ad.FilledAmount.Equal(ed.FilledAmount))
	assert.True(t, ad.FilledFees.Equal(ed.FilledFees))
	assert.Equal(t, ed.Id, ad.Id)
	assert.Equal(t, ed.State, ad.State)
	assert.Equal(t, ed.Type, ad.Type)
}
