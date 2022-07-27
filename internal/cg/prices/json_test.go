package prices

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestFindFiles(t *testing.T) {
	expected := []string{
		"test_data/a/23/2022-07-23T14:10:02/prices.json",
		"test_data/a/23/2022-07-23T14:15:01/prices.json",
		"test_data/a/24/2022-07-24T15:20:02/prices.json",
		"test_data/a/24/2022-07-24T15:25:01/prices.json",
		"test_data/a/25/2022-07-25T16:30:01/prices.json",
		"test_data/a/25/2022-07-25T16:35:01/prices.json",
	}
	actual, err := find("test_data/a")
	assert.Nil(t, err)
	assert.NotNil(t, actual)
	assert.Equal(t, 6, len(actual))
	assert.Equal(t, expected, actual)
}

func TestParseFile(t *testing.T) {
	expected := []Multi{
		{
			Base: "cardano",
			TS:   1658585374,
			Data: []Single{
				{
					Quote:      "usd",
					Price:      decimal.NewFromFloat(0.483638),
					QVol:       decimal.NewFromFloat(709567314.465705),
					QVolChange: decimal.NewFromFloat(-3.0629958920522284),
				},
				{
					Quote:      "eur",
					Price:      decimal.NewFromFloat(0.473551),
					QVol:       decimal.NewFromFloat(694768578.5552071),
					QVolChange: decimal.NewFromFloat(-2.8900286983096284),
				},
			},
		},
		{
			Base: "ethereum",
			TS:   1658585335,
			Data: []Single{
				{
					Quote:      "jpy",
					Price:      decimal.NewFromFloat(208926),
					QVol:       decimal.NewFromFloat(2092031509387.3545),
					QVolChange: decimal.NewFromFloat(-4.038405807993545),
				},
				{
					Quote:      "chf",
					Price:      decimal.NewFromFloat(1475.41),
					QVol:       decimal.NewFromFloat(14773695529.076485),
					QVolChange: decimal.NewFromFloat(-4.168956490757027),
				},
			},
		},
		{
			Base: "litecoin",
			TS:   1658585350,
			Data: []Single{
				{
					Quote:      "usd",
					Price:      decimal.NewFromFloat(55.7),
					QVol:       decimal.NewFromFloat(383496778.0165852),
					QVolChange: decimal.NewFromFloat(-4.519129782073681),
				},
				{
					Quote:      "krw",
					Price:      decimal.NewFromFloat(73011),
					QVol:       decimal.NewFromFloat(502683741656.3606),
					QVolChange: decimal.NewFromFloat(-4.252791198878055),
				},
			},
		},
	}
	actual, err := parse("test_data/a/23/2022-07-23T14:10:02/prices.json")
	assert.Nil(t, err)
	assert.NotNil(t, actual)
	assert.Equal(t, 3, len(actual))
	assert.True(t, assert.ObjectsAreEqualValues(expected, actual))
}
