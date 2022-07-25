package ohlc

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

const (
	day     = "2022-07-17"
	chain   = "ethereum"
	slug    = "regulars"
	dsource = "opensea"
)

func TestFindFiles(t *testing.T) {
	expected := []string{
		"test_data/ohlc/a/23/bitcoin_usd.ohlc",
		"test_data/ohlc/a/23/cardano_usd.ohlc",
		"test_data/ohlc/a/23/ethereum_usd.ohlc",
		"test_data/ohlc/a/24/litecoin_eur.ohlc",
		"test_data/ohlc/a/24/polkadot_eur.ohlc",
		"test_data/ohlc/a/24/solana_eur.ohlc",
		"test_data/ohlc/a/25/bitcoin_krw.ohlc",
		"test_data/ohlc/a/25/cardano_krw.ohlc",
		"test_data/ohlc/a/25/ethereum_krw.ohlc",
		"test_data/ohlc/a/25/litecoin_krw.ohlc",
		"test_data/ohlc/a/25/polkadot_krw.ohlc",
		"test_data/ohlc/a/25/solana_krw.ohlc"}
	actual, err := find("test_data/ohlc/a")
	assert.Nil(t, err)
	assert.NotNil(t, actual)
	assert.Equal(t, 12, len(actual))
	assert.Equal(t, expected, actual)
}

func TestParseFile(t *testing.T) {
	expected := []Data{
		{1658503800000, decimal.NewFromFloat(23591.87), decimal.NewFromFloat(23591.87), decimal.NewFromFloat(23396.74), decimal.NewFromFloat(23396.74)},
		{1658505600000, decimal.NewFromFloat(23381.84), decimal.NewFromFloat(23381.84), decimal.NewFromFloat(23287.97), decimal.NewFromFloat(23299.11)},
		{1658507400000, decimal.NewFromFloat(23293.79), decimal.NewFromFloat(23304.54), decimal.NewFromFloat(23155.67), decimal.NewFromFloat(23166.81)},
		{1658509200000, decimal.NewFromFloat(23163.1), decimal.NewFromFloat(23181.54), decimal.NewFromFloat(23114.55), decimal.NewFromFloat(23128.31)},
		{1658511000000, decimal.NewFromFloat(23130.04), decimal.NewFromFloat(23161.48), decimal.NewFromFloat(23044.85), decimal.NewFromFloat(23044.85)},
		{1658512800000, decimal.NewFromFloat(23035.04), decimal.NewFromFloat(23035.04), decimal.NewFromFloat(22956.03), decimal.NewFromFloat(22971.73)},
		{1658514600000, decimal.NewFromFloat(22984.02), decimal.NewFromFloat(22991.03), decimal.NewFromFloat(22959.49), decimal.NewFromFloat(22991.03)},
		{1658516400000, decimal.NewFromFloat(23015.52), decimal.NewFromFloat(23031.5), decimal.NewFromFloat(23012.61), decimal.NewFromFloat(23031.5)},
		{1658518200000, decimal.NewFromFloat(23041.24), decimal.NewFromFloat(23126.98), decimal.NewFromFloat(23041.24), decimal.NewFromFloat(23126.98)},
		{1658520000000, decimal.NewFromFloat(23142.42), decimal.NewFromFloat(23142.42), decimal.NewFromFloat(22648.7), decimal.NewFromFloat(22648.7)},
		{1658521800000, decimal.NewFromFloat(22634.8), decimal.NewFromFloat(22693.12), decimal.NewFromFloat(22634.8), decimal.NewFromFloat(22693.12)},
		{1658523600000, decimal.NewFromFloat(22734.39), decimal.NewFromFloat(22734.39), decimal.NewFromFloat(22645.69), decimal.NewFromFloat(22660.54)},
		{1658525400000, decimal.NewFromFloat(22620.1), decimal.NewFromFloat(22720.64), decimal.NewFromFloat(22620.1), decimal.NewFromFloat(22720.64)},
		{1658527200000, decimal.NewFromFloat(22709.24), decimal.NewFromFloat(22759.06), decimal.NewFromFloat(22709.24), decimal.NewFromFloat(22759.06)},
		{1658529000000, decimal.NewFromFloat(22750.71), decimal.NewFromFloat(22760.46), decimal.NewFromFloat(22747.02), decimal.NewFromFloat(22755.41)},
		{1658530800000, decimal.NewFromFloat(22679.54), decimal.NewFromFloat(22737.28), decimal.NewFromFloat(22679.54), decimal.NewFromFloat(22737.28)},
		{1658532600000, decimal.NewFromFloat(22742.12), decimal.NewFromFloat(22747.95), decimal.NewFromFloat(22729.6), decimal.NewFromFloat(22739.41)},
		{1658534400000, decimal.NewFromFloat(22741.04), decimal.NewFromFloat(22769.17), decimal.NewFromFloat(22693.25), decimal.NewFromFloat(22696.9)},
		{1658536200000, decimal.NewFromFloat(22700.1), decimal.NewFromFloat(22700.1), decimal.NewFromFloat(22663.78), decimal.NewFromFloat(22690.25)},
		{1658538000000, decimal.NewFromFloat(22705.06), decimal.NewFromFloat(22789.97), decimal.NewFromFloat(22705.06), decimal.NewFromFloat(22781.41)},
		{1658539800000, decimal.NewFromFloat(22780.25), decimal.NewFromFloat(22798.09), decimal.NewFromFloat(22779.12), decimal.NewFromFloat(22794.48)},
		{1658541600000, decimal.NewFromFloat(22806.84), decimal.NewFromFloat(22825.2), decimal.NewFromFloat(22789.72), decimal.NewFromFloat(22810.34)},
		{1658543400000, decimal.NewFromFloat(22800.59), decimal.NewFromFloat(22820.98), decimal.NewFromFloat(22800.59), decimal.NewFromFloat(22818.73)},
		{1658545200000, decimal.NewFromFloat(22818.44), decimal.NewFromFloat(22924.71), decimal.NewFromFloat(22818.44), decimal.NewFromFloat(22869.87)},
		{1658547000000, decimal.NewFromFloat(22868.48), decimal.NewFromFloat(22868.48), decimal.NewFromFloat(22832.91), decimal.NewFromFloat(22862.05)},
		{1658548800000, decimal.NewFromFloat(22865.64), decimal.NewFromFloat(22886.45), decimal.NewFromFloat(22862.33), decimal.NewFromFloat(22862.33)},
		{1658550600000, decimal.NewFromFloat(22856.8), decimal.NewFromFloat(22856.8), decimal.NewFromFloat(22775.15), decimal.NewFromFloat(22778.12)},
		{1658552400000, decimal.NewFromFloat(22777.02), decimal.NewFromFloat(22894.35), decimal.NewFromFloat(22777.02), decimal.NewFromFloat(22894.35)},
		{1658554200000, decimal.NewFromFloat(22974.56), decimal.NewFromFloat(22976.25), decimal.NewFromFloat(22960.34), decimal.NewFromFloat(22964.73)},
		{1658556000000, decimal.NewFromFloat(22940.6), decimal.NewFromFloat(22944.76), decimal.NewFromFloat(22911.12), decimal.NewFromFloat(22928.33)},
		{1658557800000, decimal.NewFromFloat(22913.85), decimal.NewFromFloat(22913.85), decimal.NewFromFloat(22892.93), decimal.NewFromFloat(22892.93)},
		{1658559600000, decimal.NewFromFloat(22888.56), decimal.NewFromFloat(22906.37), decimal.NewFromFloat(22886.27), decimal.NewFromFloat(22892.28)},
		{1658561400000, decimal.NewFromFloat(22860.91), decimal.NewFromFloat(22862.65), decimal.NewFromFloat(22768.08), decimal.NewFromFloat(22801.67)},
		{1658563200000, decimal.NewFromFloat(22795.42), decimal.NewFromFloat(22795.42), decimal.NewFromFloat(22734.07), decimal.NewFromFloat(22734.07)},
		{1658565000000, decimal.NewFromFloat(22766.38), decimal.NewFromFloat(22822.54), decimal.NewFromFloat(22766.38), decimal.NewFromFloat(22818.89)},
		{1658566800000, decimal.NewFromFloat(22871.87), decimal.NewFromFloat(22871.87), decimal.NewFromFloat(22706.93), decimal.NewFromFloat(22706.93)},
		{1658568600000, decimal.NewFromFloat(22689.79), decimal.NewFromFloat(22689.79), decimal.NewFromFloat(22619.31), decimal.NewFromFloat(22619.31)},
		{1658570400000, decimal.NewFromFloat(22672.29), decimal.NewFromFloat(22700.56), decimal.NewFromFloat(22639.75), decimal.NewFromFloat(22655.92)},
		{1658572200000, decimal.NewFromFloat(22648.23), decimal.NewFromFloat(22710.13), decimal.NewFromFloat(22641.15), decimal.NewFromFloat(22692.94)},
		{1658574000000, decimal.NewFromFloat(22637.02), decimal.NewFromFloat(22637.02), decimal.NewFromFloat(22524.29), decimal.NewFromFloat(22541.53)},
		{1658575800000, decimal.NewFromFloat(22494.86), decimal.NewFromFloat(22524.95), decimal.NewFromFloat(22444.95), decimal.NewFromFloat(22472.35)},
		{1658577600000, decimal.NewFromFloat(22489.26), decimal.NewFromFloat(22489.26), decimal.NewFromFloat(22255.3), decimal.NewFromFloat(22255.3)},
		{1658579400000, decimal.NewFromFloat(22255.15), decimal.NewFromFloat(22293.47), decimal.New(222230, -1), decimal.NewFromFloat(22293.47)},
		{1658581200000, decimal.NewFromFloat(22313.14), decimal.NewFromFloat(22318.9), decimal.NewFromFloat(22309.43), decimal.NewFromFloat(22309.43)},
		{1658583000000, decimal.NewFromFloat(22291.46), decimal.NewFromFloat(22326.43), decimal.NewFromFloat(22291.46), decimal.NewFromFloat(22326.43)},
		{1658584800000, decimal.NewFromFloat(22313.97), decimal.NewFromFloat(22344.6), decimal.NewFromFloat(22305.95), decimal.NewFromFloat(22344.6)},
		{1658586600000, decimal.NewFromFloat(22353.48), decimal.NewFromFloat(22353.48), decimal.NewFromFloat(22330.41), decimal.NewFromFloat(22336.16)},
		{1658588400000, decimal.NewFromFloat(22317.55), decimal.NewFromFloat(22317.55), decimal.NewFromFloat(22240.96), decimal.NewFromFloat(22240.96)},
		{1658590200000, decimal.NewFromFloat(22230.87), decimal.NewFromFloat(22289.08), decimal.NewFromFloat(22230.87), decimal.NewFromFloat(22289.08)},
	}
	actual, err := parse("test_data/ohlc/a/23/bitcoin_usd.ohlc")
	assert.Nil(t, err)
	assert.NotNil(t, actual)
	assert.Equal(t, 49, len(actual))
	assert.Equal(t, expected, actual)
}
