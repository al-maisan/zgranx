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

func TestParseFileBTC(t *testing.T) {
	expected := []OHLC{
		{1658503800, decimal.NewFromFloat(23591.87), decimal.NewFromFloat(23591.87), decimal.NewFromFloat(23396.74), decimal.NewFromFloat(23396.74)},
		{1658505600, decimal.NewFromFloat(23381.84), decimal.NewFromFloat(23381.84), decimal.NewFromFloat(23287.97), decimal.NewFromFloat(23299.11)},
		{1658507400, decimal.NewFromFloat(23293.79), decimal.NewFromFloat(23304.54), decimal.NewFromFloat(23155.67), decimal.NewFromFloat(23166.81)},
		{1658509200, decimal.NewFromFloat(23163.1), decimal.NewFromFloat(23181.54), decimal.NewFromFloat(23114.55), decimal.NewFromFloat(23128.31)},
		{1658511000, decimal.NewFromFloat(23130.04), decimal.NewFromFloat(23161.48), decimal.NewFromFloat(23044.85), decimal.NewFromFloat(23044.85)},
		{1658512800, decimal.NewFromFloat(23035.04), decimal.NewFromFloat(23035.04), decimal.NewFromFloat(22956.03), decimal.NewFromFloat(22971.73)},
		{1658514600, decimal.NewFromFloat(22984.02), decimal.NewFromFloat(22991.03), decimal.NewFromFloat(22959.49), decimal.NewFromFloat(22991.03)},
		{1658516400, decimal.NewFromFloat(23015.52), decimal.NewFromFloat(23031.5), decimal.NewFromFloat(23012.61), decimal.NewFromFloat(23031.5)},
		{1658518200, decimal.NewFromFloat(23041.24), decimal.NewFromFloat(23126.98), decimal.NewFromFloat(23041.24), decimal.NewFromFloat(23126.98)},
		{1658520000, decimal.NewFromFloat(23142.42), decimal.NewFromFloat(23142.42), decimal.NewFromFloat(22648.7), decimal.NewFromFloat(22648.7)},
		{1658521800, decimal.NewFromFloat(22634.8), decimal.NewFromFloat(22693.12), decimal.NewFromFloat(22634.8), decimal.NewFromFloat(22693.12)},
		{1658523600, decimal.NewFromFloat(22734.39), decimal.NewFromFloat(22734.39), decimal.NewFromFloat(22645.69), decimal.NewFromFloat(22660.54)},
		{1658525400, decimal.NewFromFloat(22620.1), decimal.NewFromFloat(22720.64), decimal.NewFromFloat(22620.1), decimal.NewFromFloat(22720.64)},
		{1658527200, decimal.NewFromFloat(22709.24), decimal.NewFromFloat(22759.06), decimal.NewFromFloat(22709.24), decimal.NewFromFloat(22759.06)},
		{1658529000, decimal.NewFromFloat(22750.71), decimal.NewFromFloat(22760.46), decimal.NewFromFloat(22747.02), decimal.NewFromFloat(22755.41)},
		{1658530800, decimal.NewFromFloat(22679.54), decimal.NewFromFloat(22737.28), decimal.NewFromFloat(22679.54), decimal.NewFromFloat(22737.28)},
		{1658532600, decimal.NewFromFloat(22742.12), decimal.NewFromFloat(22747.95), decimal.NewFromFloat(22729.6), decimal.NewFromFloat(22739.41)},
		{1658534400, decimal.NewFromFloat(22741.04), decimal.NewFromFloat(22769.17), decimal.NewFromFloat(22693.25), decimal.NewFromFloat(22696.9)},
		{1658536200, decimal.NewFromFloat(22700.1), decimal.NewFromFloat(22700.1), decimal.NewFromFloat(22663.78), decimal.NewFromFloat(22690.25)},
		{1658538000, decimal.NewFromFloat(22705.06), decimal.NewFromFloat(22789.97), decimal.NewFromFloat(22705.06), decimal.NewFromFloat(22781.41)},
		{1658539800, decimal.NewFromFloat(22780.25), decimal.NewFromFloat(22798.09), decimal.NewFromFloat(22779.12), decimal.NewFromFloat(22794.48)},
		{1658541600, decimal.NewFromFloat(22806.84), decimal.NewFromFloat(22825.2), decimal.NewFromFloat(22789.72), decimal.NewFromFloat(22810.34)},
		{1658543400, decimal.NewFromFloat(22800.59), decimal.NewFromFloat(22820.98), decimal.NewFromFloat(22800.59), decimal.NewFromFloat(22818.73)},
		{1658545200, decimal.NewFromFloat(22818.44), decimal.NewFromFloat(22924.71), decimal.NewFromFloat(22818.44), decimal.NewFromFloat(22869.87)},
		{1658547000, decimal.NewFromFloat(22868.48), decimal.NewFromFloat(22868.48), decimal.NewFromFloat(22832.91), decimal.NewFromFloat(22862.05)},
		{1658548800, decimal.NewFromFloat(22865.64), decimal.NewFromFloat(22886.45), decimal.NewFromFloat(22862.33), decimal.NewFromFloat(22862.33)},
		{1658550600, decimal.NewFromFloat(22856.8), decimal.NewFromFloat(22856.8), decimal.NewFromFloat(22775.15), decimal.NewFromFloat(22778.12)},
		{1658552400, decimal.NewFromFloat(22777.02), decimal.NewFromFloat(22894.35), decimal.NewFromFloat(22777.02), decimal.NewFromFloat(22894.35)},
		{1658554200, decimal.NewFromFloat(22974.56), decimal.NewFromFloat(22976.25), decimal.NewFromFloat(22960.34), decimal.NewFromFloat(22964.73)},
		{1658556000, decimal.NewFromFloat(22940.6), decimal.NewFromFloat(22944.76), decimal.NewFromFloat(22911.12), decimal.NewFromFloat(22928.33)},
		{1658557800, decimal.NewFromFloat(22913.85), decimal.NewFromFloat(22913.85), decimal.NewFromFloat(22892.93), decimal.NewFromFloat(22892.93)},
		{1658559600, decimal.NewFromFloat(22888.56), decimal.NewFromFloat(22906.37), decimal.NewFromFloat(22886.27), decimal.NewFromFloat(22892.28)},
		{1658561400, decimal.NewFromFloat(22860.91), decimal.NewFromFloat(22862.65), decimal.NewFromFloat(22768.08), decimal.NewFromFloat(22801.67)},
		{1658563200, decimal.NewFromFloat(22795.42), decimal.NewFromFloat(22795.42), decimal.NewFromFloat(22734.07), decimal.NewFromFloat(22734.07)},
		{1658565000, decimal.NewFromFloat(22766.38), decimal.NewFromFloat(22822.54), decimal.NewFromFloat(22766.38), decimal.NewFromFloat(22818.89)},
		{1658566800, decimal.NewFromFloat(22871.87), decimal.NewFromFloat(22871.87), decimal.NewFromFloat(22706.93), decimal.NewFromFloat(22706.93)},
		{1658568600, decimal.NewFromFloat(22689.79), decimal.NewFromFloat(22689.79), decimal.NewFromFloat(22619.31), decimal.NewFromFloat(22619.31)},
		{1658570400, decimal.NewFromFloat(22672.29), decimal.NewFromFloat(22700.56), decimal.NewFromFloat(22639.75), decimal.NewFromFloat(22655.92)},
		{1658572200, decimal.NewFromFloat(22648.23), decimal.NewFromFloat(22710.13), decimal.NewFromFloat(22641.15), decimal.NewFromFloat(22692.94)},
		{1658574000, decimal.NewFromFloat(22637.02), decimal.NewFromFloat(22637.02), decimal.NewFromFloat(22524.29), decimal.NewFromFloat(22541.53)},
		{1658575800, decimal.NewFromFloat(22494.86), decimal.NewFromFloat(22524.95), decimal.NewFromFloat(22444.95), decimal.NewFromFloat(22472.35)},
		{1658577600, decimal.NewFromFloat(22489.26), decimal.NewFromFloat(22489.26), decimal.NewFromFloat(22255.3), decimal.NewFromFloat(22255.3)},
		{1658579400, decimal.NewFromFloat(22255.15), decimal.NewFromFloat(22293.47), decimal.New(222230, -1), decimal.NewFromFloat(22293.47)},
		{1658581200, decimal.NewFromFloat(22313.14), decimal.NewFromFloat(22318.9), decimal.NewFromFloat(22309.43), decimal.NewFromFloat(22309.43)},
		{1658583000, decimal.NewFromFloat(22291.46), decimal.NewFromFloat(22326.43), decimal.NewFromFloat(22291.46), decimal.NewFromFloat(22326.43)},
		{1658584800, decimal.NewFromFloat(22313.97), decimal.NewFromFloat(22344.6), decimal.NewFromFloat(22305.95), decimal.NewFromFloat(22344.6)},
		{1658586600, decimal.NewFromFloat(22353.48), decimal.NewFromFloat(22353.48), decimal.NewFromFloat(22330.41), decimal.NewFromFloat(22336.16)},
		{1658588400, decimal.NewFromFloat(22317.55), decimal.NewFromFloat(22317.55), decimal.NewFromFloat(22240.96), decimal.NewFromFloat(22240.96)},
		{1658590200, decimal.NewFromFloat(22230.87), decimal.NewFromFloat(22289.08), decimal.NewFromFloat(22230.87), decimal.NewFromFloat(22289.08)},
	}
	actual, err := parse("test_data/ohlc/a/23/bitcoin_usd.ohlc")
	assert.Nil(t, err)
	assert.NotNil(t, actual)
	assert.Equal(t, 49, len(actual))
	assert.Equal(t, expected, actual)
}

func TestParseFileETH(t *testing.T) {
	expected := []OHLC{
		{1658503800, decimal.NewFromFloat(1604.49), decimal.NewFromFloat(1604.49), decimal.NewFromFloat(1588.17), decimal.NewFromFloat(1588.17)},
		{1658505600, decimal.NewFromFloat(1580.76), decimal.NewFromFloat(1580.76), decimal.NewFromFloat(1575.37), decimal.NewFromFloat(1576.38)},
		{1658507400, decimal.NewFromFloat(1575.56), decimal.NewFromFloat(1575.76), decimal.NewFromFloat(1566.92), decimal.NewFromFloat(1575.76)},
		{1658509200, decimal.NewFromFloat(1577.02), decimal.NewFromFloat(1580.27), decimal.NewFromFloat(1577.02), decimal.NewFromFloat(1580.27)},
		{1658511000, decimal.NewFromFloat(1581.09), decimal.NewFromFloat(1586.54), decimal.NewFromFloat(1578.77), decimal.NewFromFloat(1578.77)},
		{1658512800, decimal.NewFromFloat(1575.91), decimal.NewFromFloat(1577.23), decimal.NewFromFloat(1564.48), decimal.NewFromFloat(1564.48)},
		{1658514600, decimal.NewFromFloat(1568.44), decimal.NewFromFloat(1568.44), decimal.NewFromFloat(1563.11), decimal.NewFromFloat(1564.81)},
		{1658516400, decimal.NewFromFloat(1567.17), decimal.NewFromFloat(1574.23), decimal.NewFromFloat(1567.17), decimal.NewFromFloat(1574.23)},
		{1658518200, decimal.NewFromFloat(1576.02), decimal.NewFromFloat(1580.64), decimal.NewFromFloat(1574.5), decimal.NewFromFloat(1580.64)},
		{1658520000, decimal.NewFromFloat(1579.14), decimal.NewFromFloat(1579.14), decimal.NewFromFloat(1532.15), decimal.NewFromFloat(1532.15)},
		{1658521800, decimal.NewFromFloat(1530.48), decimal.NewFromFloat(1531.76), decimal.New(15280, -1), decimal.NewFromFloat(1531.76)},
		{1658523600, decimal.NewFromFloat(1531.17), decimal.NewFromFloat(1534.62), decimal.NewFromFloat(1528.18), decimal.NewFromFloat(1528.18)},
		{1658525400, decimal.NewFromFloat(1526.75), decimal.NewFromFloat(1538.57), decimal.NewFromFloat(1526.75), decimal.NewFromFloat(1538.57)},
		{1658527200, decimal.NewFromFloat(1537.18), decimal.NewFromFloat(1541.49), decimal.NewFromFloat(1537.18), decimal.NewFromFloat(1541.49)},
		{1658529000, decimal.NewFromFloat(1540.72), decimal.NewFromFloat(1540.72), decimal.NewFromFloat(1537.44), decimal.NewFromFloat(1537.44)},
		{1658530800, decimal.NewFromFloat(1533.29), decimal.NewFromFloat(1542.11), decimal.NewFromFloat(1533.29), decimal.NewFromFloat(1542.11)},
		{1658532600, decimal.NewFromFloat(1541.05), decimal.NewFromFloat(1542.1), decimal.NewFromFloat(1538.17), decimal.NewFromFloat(1541.35)},
		{1658534400, decimal.NewFromFloat(1540.76), decimal.NewFromFloat(1540.76), decimal.NewFromFloat(1535.34), decimal.NewFromFloat(1536.12)},
		{1658536200, decimal.NewFromFloat(1537.65), decimal.NewFromFloat(1538.08), decimal.NewFromFloat(1532.04), decimal.NewFromFloat(1538.08)},
		{1658538000, decimal.NewFromFloat(1546.91), decimal.NewFromFloat(1547.09), decimal.NewFromFloat(1544.03), decimal.NewFromFloat(1545.23)},
		{1658539800, decimal.NewFromFloat(1545.6), decimal.NewFromFloat(1548.03), decimal.NewFromFloat(1544.63), decimal.NewFromFloat(1547.27)},
		{1658541600, decimal.NewFromFloat(1549.89), decimal.NewFromFloat(1553.61), decimal.NewFromFloat(1549.61), decimal.NewFromFloat(1553.61)},
		{1658543400, decimal.NewFromFloat(1553.04), decimal.NewFromFloat(1561.84), decimal.NewFromFloat(1553.04), decimal.NewFromFloat(1561.84)},
		{1658545200, decimal.NewFromFloat(1564.6), decimal.NewFromFloat(1573.54), decimal.NewFromFloat(1564.6), decimal.NewFromFloat(1569.81)},
		{1658547000, decimal.NewFromFloat(1569.1), decimal.NewFromFloat(1570.97), decimal.NewFromFloat(1565.43), decimal.NewFromFloat(1570.97)},
		{1658548800, decimal.NewFromFloat(1571.71), decimal.NewFromFloat(1577.24), decimal.NewFromFloat(1571.71), decimal.NewFromFloat(1575.13)},
		{1658550600, decimal.NewFromFloat(1574.25), decimal.NewFromFloat(1575.45), decimal.NewFromFloat(1568.79), decimal.NewFromFloat(1568.79)},
		{1658552400, decimal.NewFromFloat(1571.61), decimal.NewFromFloat(1586.54), decimal.NewFromFloat(1571.61), decimal.NewFromFloat(1586.54)},
		{1658554200, decimal.NewFromFloat(1591.37), decimal.NewFromFloat(1591.93), decimal.NewFromFloat(1588.75), decimal.NewFromFloat(1589.91)},
		{1658556000, decimal.NewFromFloat(1588.33), decimal.NewFromFloat(1588.33), decimal.NewFromFloat(1581.69), decimal.NewFromFloat(1584.05)},
		{1658557800, decimal.NewFromFloat(1582.26), decimal.NewFromFloat(1582.86), decimal.NewFromFloat(1581.04), decimal.NewFromFloat(1581.04)},
		{1658559600, decimal.NewFromFloat(1581.4), decimal.NewFromFloat(1588.48), decimal.NewFromFloat(1581.4), decimal.NewFromFloat(1586.19)},
		{1658561400, decimal.NewFromFloat(1583.46), decimal.NewFromFloat(1583.46), decimal.NewFromFloat(1576.88), decimal.NewFromFloat(1583.23)},
		{1658563200, decimal.NewFromFloat(1582.94), decimal.NewFromFloat(1582.94), decimal.NewFromFloat(1575.13), decimal.NewFromFloat(1575.13)},
		{1658565000, decimal.NewFromFloat(1579.49), decimal.NewFromFloat(1584.7), decimal.NewFromFloat(1579.49), decimal.NewFromFloat(1581.67)},
		{1658566800, decimal.NewFromFloat(1586.3), decimal.NewFromFloat(1586.3), decimal.NewFromFloat(1562.45), decimal.NewFromFloat(1562.45)},
		{1658568600, decimal.NewFromFloat(1557.26), decimal.NewFromFloat(1559.37), decimal.NewFromFloat(1555.72), decimal.NewFromFloat(1556.57)},
		{1658570400, decimal.NewFromFloat(1561.04), decimal.NewFromFloat(1566.01), decimal.NewFromFloat(1559.52), decimal.NewFromFloat(1559.52)},
		{1658572200, decimal.NewFromFloat(1559.74), decimal.NewFromFloat(1563.92), decimal.NewFromFloat(1558.27), decimal.NewFromFloat(1560.49)},
		{1658574000, decimal.NewFromFloat(1556.97), decimal.NewFromFloat(1556.97), decimal.NewFromFloat(1544.08), decimal.NewFromFloat(1544.08)},
		{1658575800, decimal.NewFromFloat(1541.76), decimal.NewFromFloat(1541.76), decimal.NewFromFloat(1533.03), decimal.NewFromFloat(1537.85)},
		{1658577600, decimal.NewFromFloat(1539.63), decimal.NewFromFloat(1539.63), decimal.NewFromFloat(1523.69), decimal.NewFromFloat(1523.69)},
		{1658579400, decimal.NewFromFloat(1524.69), decimal.NewFromFloat(1524.69), decimal.NewFromFloat(1516.99), decimal.NewFromFloat(1524.05)},
		{1658581200, decimal.NewFromFloat(1522.79), decimal.NewFromFloat(1524.44), decimal.NewFromFloat(1521.92), decimal.NewFromFloat(1521.92)},
		{1658583000, decimal.NewFromFloat(1521.52), decimal.NewFromFloat(1526.18), decimal.NewFromFloat(1521.52), decimal.NewFromFloat(1526.18)},
		{1658584800, decimal.NewFromFloat(1523.86), decimal.NewFromFloat(1533.35), decimal.NewFromFloat(1523.86), decimal.NewFromFloat(1533.35)},
		{1658586600, decimal.NewFromFloat(1532.4), decimal.NewFromFloat(1535.02), decimal.NewFromFloat(1531.7), decimal.NewFromFloat(1531.7)},
		{1658588400, decimal.New(15290, -1), decimal.NewFromFloat(1530.14), decimal.NewFromFloat(1522.44), decimal.NewFromFloat(1522.44)},
		{1658590200, decimal.NewFromFloat(1530.99), decimal.NewFromFloat(1530.99), decimal.NewFromFloat(1530.36), decimal.NewFromFloat(1530.95)},
	}
	actual, err := parse("test_data/ohlc/a/23/ethereum_usd.ohlc")
	assert.Nil(t, err)
	assert.NotNil(t, actual)
	assert.Equal(t, 49, len(actual))
	assert.Equal(t, expected, actual)
}

func TestTradingPair(t *testing.T) {
	b, q := tradingPair("test_data/ohlc/a/23/bitcoin_usd.ohlc")
	assert.Equal(t, "bitcoin", b)
	assert.Equal(t, "usd", q)
}

func TestTradingPairNoExt(t *testing.T) {
	b, q := tradingPair("test_data/ohlc/a/23/bitcoin_usd")
	assert.Equal(t, "bitcoin", b)
	assert.Equal(t, "usd", q)
}

func TestTradingPairNoSep(t *testing.T) {
	b, q := tradingPair("test_data/ohlc/a/23/bitcoin:usd")
	assert.Equal(t, "", b)
	assert.Equal(t, "", q)
}

func TestProcess(t *testing.T) {
	btc := []OHLC{
		{1658503800, decimal.NewFromFloat(23591.87), decimal.NewFromFloat(23591.87), decimal.NewFromFloat(23396.74), decimal.NewFromFloat(23396.74)},
		{1658505600, decimal.NewFromFloat(23381.84), decimal.NewFromFloat(23381.84), decimal.NewFromFloat(23287.97), decimal.NewFromFloat(23299.11)},
		{1658507400, decimal.NewFromFloat(23293.79), decimal.NewFromFloat(23304.54), decimal.NewFromFloat(23155.67), decimal.NewFromFloat(23166.81)},
		{1658509200, decimal.NewFromFloat(23163.1), decimal.NewFromFloat(23181.54), decimal.NewFromFloat(23114.55), decimal.NewFromFloat(23128.31)},
		{1658511000, decimal.NewFromFloat(23130.04), decimal.NewFromFloat(23161.48), decimal.NewFromFloat(23044.85), decimal.NewFromFloat(23044.85)},
		{1658512800, decimal.NewFromFloat(23035.04), decimal.NewFromFloat(23035.04), decimal.NewFromFloat(22956.03), decimal.NewFromFloat(22971.73)},
		{1658514600, decimal.NewFromFloat(22984.02), decimal.NewFromFloat(22991.03), decimal.NewFromFloat(22959.49), decimal.NewFromFloat(22991.03)},
		{1658516400, decimal.NewFromFloat(23015.52), decimal.NewFromFloat(23031.5), decimal.NewFromFloat(23012.61), decimal.NewFromFloat(23031.5)},
		{1658518200, decimal.NewFromFloat(23041.24), decimal.NewFromFloat(23126.98), decimal.NewFromFloat(23041.24), decimal.NewFromFloat(23126.98)},
		{1658520000, decimal.NewFromFloat(23142.42), decimal.NewFromFloat(23142.42), decimal.NewFromFloat(22648.7), decimal.NewFromFloat(22648.7)},
		{1658521800, decimal.NewFromFloat(22634.8), decimal.NewFromFloat(22693.12), decimal.NewFromFloat(22634.8), decimal.NewFromFloat(22693.12)},
		{1658523600, decimal.NewFromFloat(22734.39), decimal.NewFromFloat(22734.39), decimal.NewFromFloat(22645.69), decimal.NewFromFloat(22660.54)},
		{1658525400, decimal.NewFromFloat(22620.1), decimal.NewFromFloat(22720.64), decimal.NewFromFloat(22620.1), decimal.NewFromFloat(22720.64)},
		{1658527200, decimal.NewFromFloat(22709.24), decimal.NewFromFloat(22759.06), decimal.NewFromFloat(22709.24), decimal.NewFromFloat(22759.06)},
		{1658529000, decimal.NewFromFloat(22750.71), decimal.NewFromFloat(22760.46), decimal.NewFromFloat(22747.02), decimal.NewFromFloat(22755.41)},
		{1658530800, decimal.NewFromFloat(22679.54), decimal.NewFromFloat(22737.28), decimal.NewFromFloat(22679.54), decimal.NewFromFloat(22737.28)},
		{1658532600, decimal.NewFromFloat(22742.12), decimal.NewFromFloat(22747.95), decimal.NewFromFloat(22729.6), decimal.NewFromFloat(22739.41)},
		{1658534400, decimal.NewFromFloat(22741.04), decimal.NewFromFloat(22769.17), decimal.NewFromFloat(22693.25), decimal.NewFromFloat(22696.9)},
		{1658536200, decimal.NewFromFloat(22700.1), decimal.NewFromFloat(22700.1), decimal.NewFromFloat(22663.78), decimal.NewFromFloat(22690.25)},
		{1658538000, decimal.NewFromFloat(22705.06), decimal.NewFromFloat(22789.97), decimal.NewFromFloat(22705.06), decimal.NewFromFloat(22781.41)},
		{1658539800, decimal.NewFromFloat(22780.25), decimal.NewFromFloat(22798.09), decimal.NewFromFloat(22779.12), decimal.NewFromFloat(22794.48)},
		{1658541600, decimal.NewFromFloat(22806.84), decimal.NewFromFloat(22825.2), decimal.NewFromFloat(22789.72), decimal.NewFromFloat(22810.34)},
		{1658543400, decimal.NewFromFloat(22800.59), decimal.NewFromFloat(22820.98), decimal.NewFromFloat(22800.59), decimal.NewFromFloat(22818.73)},
		{1658545200, decimal.NewFromFloat(22818.44), decimal.NewFromFloat(22924.71), decimal.NewFromFloat(22818.44), decimal.NewFromFloat(22869.87)},
		{1658547000, decimal.NewFromFloat(22868.48), decimal.NewFromFloat(22868.48), decimal.NewFromFloat(22832.91), decimal.NewFromFloat(22862.05)},
		{1658548800, decimal.NewFromFloat(22865.64), decimal.NewFromFloat(22886.45), decimal.NewFromFloat(22862.33), decimal.NewFromFloat(22862.33)},
		{1658550600, decimal.NewFromFloat(22856.8), decimal.NewFromFloat(22856.8), decimal.NewFromFloat(22775.15), decimal.NewFromFloat(22778.12)},
		{1658552400, decimal.NewFromFloat(22777.02), decimal.NewFromFloat(22894.35), decimal.NewFromFloat(22777.02), decimal.NewFromFloat(22894.35)},
		{1658554200, decimal.NewFromFloat(22974.56), decimal.NewFromFloat(22976.25), decimal.NewFromFloat(22960.34), decimal.NewFromFloat(22964.73)},
		{1658556000, decimal.NewFromFloat(22940.6), decimal.NewFromFloat(22944.76), decimal.NewFromFloat(22911.12), decimal.NewFromFloat(22928.33)},
		{1658557800, decimal.NewFromFloat(22913.85), decimal.NewFromFloat(22913.85), decimal.NewFromFloat(22892.93), decimal.NewFromFloat(22892.93)},
		{1658559600, decimal.NewFromFloat(22888.56), decimal.NewFromFloat(22906.37), decimal.NewFromFloat(22886.27), decimal.NewFromFloat(22892.28)},
		{1658561400, decimal.NewFromFloat(22860.91), decimal.NewFromFloat(22862.65), decimal.NewFromFloat(22768.08), decimal.NewFromFloat(22801.67)},
		{1658563200, decimal.NewFromFloat(22795.42), decimal.NewFromFloat(22795.42), decimal.NewFromFloat(22734.07), decimal.NewFromFloat(22734.07)},
		{1658565000, decimal.NewFromFloat(22766.38), decimal.NewFromFloat(22822.54), decimal.NewFromFloat(22766.38), decimal.NewFromFloat(22818.89)},
		{1658566800, decimal.NewFromFloat(22871.87), decimal.NewFromFloat(22871.87), decimal.NewFromFloat(22706.93), decimal.NewFromFloat(22706.93)},
		{1658568600, decimal.NewFromFloat(22689.79), decimal.NewFromFloat(22689.79), decimal.NewFromFloat(22619.31), decimal.NewFromFloat(22619.31)},
		{1658570400, decimal.NewFromFloat(22672.29), decimal.NewFromFloat(22700.56), decimal.NewFromFloat(22639.75), decimal.NewFromFloat(22655.92)},
		{1658572200, decimal.NewFromFloat(22648.23), decimal.NewFromFloat(22710.13), decimal.NewFromFloat(22641.15), decimal.NewFromFloat(22692.94)},
		{1658574000, decimal.NewFromFloat(22637.02), decimal.NewFromFloat(22637.02), decimal.NewFromFloat(22524.29), decimal.NewFromFloat(22541.53)},
		{1658575800, decimal.NewFromFloat(22494.86), decimal.NewFromFloat(22524.95), decimal.NewFromFloat(22444.95), decimal.NewFromFloat(22472.35)},
		{1658577600, decimal.NewFromFloat(22489.26), decimal.NewFromFloat(22489.26), decimal.NewFromFloat(22255.3), decimal.NewFromFloat(22255.3)},
		{1658579400, decimal.NewFromFloat(22255.15), decimal.NewFromFloat(22293.47), decimal.New(222230, -1), decimal.NewFromFloat(22293.47)},
		{1658581200, decimal.NewFromFloat(22313.14), decimal.NewFromFloat(22318.9), decimal.NewFromFloat(22309.43), decimal.NewFromFloat(22309.43)},
		{1658583000, decimal.NewFromFloat(22291.46), decimal.NewFromFloat(22326.43), decimal.NewFromFloat(22291.46), decimal.NewFromFloat(22326.43)},
		{1658584800, decimal.NewFromFloat(22313.97), decimal.NewFromFloat(22344.6), decimal.NewFromFloat(22305.95), decimal.NewFromFloat(22344.6)},
		{1658586600, decimal.NewFromFloat(22353.48), decimal.NewFromFloat(22353.48), decimal.NewFromFloat(22330.41), decimal.NewFromFloat(22336.16)},
		{1658588400, decimal.NewFromFloat(22317.55), decimal.NewFromFloat(22317.55), decimal.NewFromFloat(22240.96), decimal.NewFromFloat(22240.96)},
		{1658590200, decimal.NewFromFloat(22230.87), decimal.NewFromFloat(22289.08), decimal.NewFromFloat(22230.87), decimal.NewFromFloat(22289.08)},
	}
	eth := []OHLC{
		{1658503800, decimal.NewFromFloat(1604.49), decimal.NewFromFloat(1604.49), decimal.NewFromFloat(1588.17), decimal.NewFromFloat(1588.17)},
		{1658505600, decimal.NewFromFloat(1580.76), decimal.NewFromFloat(1580.76), decimal.NewFromFloat(1575.37), decimal.NewFromFloat(1576.38)},
		{1658507400, decimal.NewFromFloat(1575.56), decimal.NewFromFloat(1575.76), decimal.NewFromFloat(1566.92), decimal.NewFromFloat(1575.76)},
		{1658509200, decimal.NewFromFloat(1577.02), decimal.NewFromFloat(1580.27), decimal.NewFromFloat(1577.02), decimal.NewFromFloat(1580.27)},
		{1658511000, decimal.NewFromFloat(1581.09), decimal.NewFromFloat(1586.54), decimal.NewFromFloat(1578.77), decimal.NewFromFloat(1578.77)},
		{1658512800, decimal.NewFromFloat(1575.91), decimal.NewFromFloat(1577.23), decimal.NewFromFloat(1564.48), decimal.NewFromFloat(1564.48)},
		{1658514600, decimal.NewFromFloat(1568.44), decimal.NewFromFloat(1568.44), decimal.NewFromFloat(1563.11), decimal.NewFromFloat(1564.81)},
		{1658516400, decimal.NewFromFloat(1567.17), decimal.NewFromFloat(1574.23), decimal.NewFromFloat(1567.17), decimal.NewFromFloat(1574.23)},
		{1658518200, decimal.NewFromFloat(1576.02), decimal.NewFromFloat(1580.64), decimal.NewFromFloat(1574.5), decimal.NewFromFloat(1580.64)},
		{1658520000, decimal.NewFromFloat(1579.14), decimal.NewFromFloat(1579.14), decimal.NewFromFloat(1532.15), decimal.NewFromFloat(1532.15)},
		{1658521800, decimal.NewFromFloat(1530.48), decimal.NewFromFloat(1531.76), decimal.New(15280, -1), decimal.NewFromFloat(1531.76)},
		{1658523600, decimal.NewFromFloat(1531.17), decimal.NewFromFloat(1534.62), decimal.NewFromFloat(1528.18), decimal.NewFromFloat(1528.18)},
		{1658525400, decimal.NewFromFloat(1526.75), decimal.NewFromFloat(1538.57), decimal.NewFromFloat(1526.75), decimal.NewFromFloat(1538.57)},
		{1658527200, decimal.NewFromFloat(1537.18), decimal.NewFromFloat(1541.49), decimal.NewFromFloat(1537.18), decimal.NewFromFloat(1541.49)},
		{1658529000, decimal.NewFromFloat(1540.72), decimal.NewFromFloat(1540.72), decimal.NewFromFloat(1537.44), decimal.NewFromFloat(1537.44)},
		{1658530800, decimal.NewFromFloat(1533.29), decimal.NewFromFloat(1542.11), decimal.NewFromFloat(1533.29), decimal.NewFromFloat(1542.11)},
		{1658532600, decimal.NewFromFloat(1541.05), decimal.NewFromFloat(1542.1), decimal.NewFromFloat(1538.17), decimal.NewFromFloat(1541.35)},
		{1658534400, decimal.NewFromFloat(1540.76), decimal.NewFromFloat(1540.76), decimal.NewFromFloat(1535.34), decimal.NewFromFloat(1536.12)},
		{1658536200, decimal.NewFromFloat(1537.65), decimal.NewFromFloat(1538.08), decimal.NewFromFloat(1532.04), decimal.NewFromFloat(1538.08)},
		{1658538000, decimal.NewFromFloat(1546.91), decimal.NewFromFloat(1547.09), decimal.NewFromFloat(1544.03), decimal.NewFromFloat(1545.23)},
		{1658539800, decimal.NewFromFloat(1545.6), decimal.NewFromFloat(1548.03), decimal.NewFromFloat(1544.63), decimal.NewFromFloat(1547.27)},
		{1658541600, decimal.NewFromFloat(1549.89), decimal.NewFromFloat(1553.61), decimal.NewFromFloat(1549.61), decimal.NewFromFloat(1553.61)},
		{1658543400, decimal.NewFromFloat(1553.04), decimal.NewFromFloat(1561.84), decimal.NewFromFloat(1553.04), decimal.NewFromFloat(1561.84)},
		{1658545200, decimal.NewFromFloat(1564.6), decimal.NewFromFloat(1573.54), decimal.NewFromFloat(1564.6), decimal.NewFromFloat(1569.81)},
		{1658547000, decimal.NewFromFloat(1569.1), decimal.NewFromFloat(1570.97), decimal.NewFromFloat(1565.43), decimal.NewFromFloat(1570.97)},
		{1658548800, decimal.NewFromFloat(1571.71), decimal.NewFromFloat(1577.24), decimal.NewFromFloat(1571.71), decimal.NewFromFloat(1575.13)},
		{1658550600, decimal.NewFromFloat(1574.25), decimal.NewFromFloat(1575.45), decimal.NewFromFloat(1568.79), decimal.NewFromFloat(1568.79)},
		{1658552400, decimal.NewFromFloat(1571.61), decimal.NewFromFloat(1586.54), decimal.NewFromFloat(1571.61), decimal.NewFromFloat(1586.54)},
		{1658554200, decimal.NewFromFloat(1591.37), decimal.NewFromFloat(1591.93), decimal.NewFromFloat(1588.75), decimal.NewFromFloat(1589.91)},
		{1658556000, decimal.NewFromFloat(1588.33), decimal.NewFromFloat(1588.33), decimal.NewFromFloat(1581.69), decimal.NewFromFloat(1584.05)},
		{1658557800, decimal.NewFromFloat(1582.26), decimal.NewFromFloat(1582.86), decimal.NewFromFloat(1581.04), decimal.NewFromFloat(1581.04)},
		{1658559600, decimal.NewFromFloat(1581.4), decimal.NewFromFloat(1588.48), decimal.NewFromFloat(1581.4), decimal.NewFromFloat(1586.19)},
		{1658561400, decimal.NewFromFloat(1583.46), decimal.NewFromFloat(1583.46), decimal.NewFromFloat(1576.88), decimal.NewFromFloat(1583.23)},
		{1658563200, decimal.NewFromFloat(1582.94), decimal.NewFromFloat(1582.94), decimal.NewFromFloat(1575.13), decimal.NewFromFloat(1575.13)},
		{1658565000, decimal.NewFromFloat(1579.49), decimal.NewFromFloat(1584.7), decimal.NewFromFloat(1579.49), decimal.NewFromFloat(1581.67)},
		{1658566800, decimal.NewFromFloat(1586.3), decimal.NewFromFloat(1586.3), decimal.NewFromFloat(1562.45), decimal.NewFromFloat(1562.45)},
		{1658568600, decimal.NewFromFloat(1557.26), decimal.NewFromFloat(1559.37), decimal.NewFromFloat(1555.72), decimal.NewFromFloat(1556.57)},
		{1658570400, decimal.NewFromFloat(1561.04), decimal.NewFromFloat(1566.01), decimal.NewFromFloat(1559.52), decimal.NewFromFloat(1559.52)},
		{1658572200, decimal.NewFromFloat(1559.74), decimal.NewFromFloat(1563.92), decimal.NewFromFloat(1558.27), decimal.NewFromFloat(1560.49)},
		{1658574000, decimal.NewFromFloat(1556.97), decimal.NewFromFloat(1556.97), decimal.NewFromFloat(1544.08), decimal.NewFromFloat(1544.08)},
		{1658575800, decimal.NewFromFloat(1541.76), decimal.NewFromFloat(1541.76), decimal.NewFromFloat(1533.03), decimal.NewFromFloat(1537.85)},
		{1658577600, decimal.NewFromFloat(1539.63), decimal.NewFromFloat(1539.63), decimal.NewFromFloat(1523.69), decimal.NewFromFloat(1523.69)},
		{1658579400, decimal.NewFromFloat(1524.69), decimal.NewFromFloat(1524.69), decimal.NewFromFloat(1516.99), decimal.NewFromFloat(1524.05)},
		{1658581200, decimal.NewFromFloat(1522.79), decimal.NewFromFloat(1524.44), decimal.NewFromFloat(1521.92), decimal.NewFromFloat(1521.92)},
		{1658583000, decimal.NewFromFloat(1521.52), decimal.NewFromFloat(1526.18), decimal.NewFromFloat(1521.52), decimal.NewFromFloat(1526.18)},
		{1658584800, decimal.NewFromFloat(1523.86), decimal.NewFromFloat(1533.35), decimal.NewFromFloat(1523.86), decimal.NewFromFloat(1533.35)},
		{1658586600, decimal.NewFromFloat(1532.4), decimal.NewFromFloat(1535.02), decimal.NewFromFloat(1531.7), decimal.NewFromFloat(1531.7)},
		{1658588400, decimal.New(15290, -1), decimal.NewFromFloat(1530.14), decimal.NewFromFloat(1522.44), decimal.NewFromFloat(1522.44)},
		{1658590200, decimal.NewFromFloat(1530.99), decimal.NewFromFloat(1530.99), decimal.NewFromFloat(1530.36), decimal.NewFromFloat(1530.95)},
	}
	expected := []Data{
		{
			Base:  "bitcoin",
			Quote: "usd",
			Data:  btc,
		},
		{
			Base:  "ethereum",
			Quote: "usd",
			Data:  eth,
		},
	}
	actual, err := Process("test_data/ohlc/b")
	assert.Nil(t, err)
	assert.NotNil(t, actual)
	assert.Equal(t, 2, len(actual))
	assert.Equal(t, expected, actual)
}
