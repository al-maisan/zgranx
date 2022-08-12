package ma

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSMA(t *testing.T) {
	expected := "2.000000"
	actual, err := SMA([]string{"1.0", "2.0", "3.0"})
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestSMASingleValue(t *testing.T) {
	expected := "3.444568"
	actual, err := SMA([]string{"3.444567890"})
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestSMAInvalidPrice(t *testing.T) {
	expected := ""
	actual, err := SMA([]string{"1.0", "not-a-price", "3.0"})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "invalid price value: 'not-a-price', can't convert not-a-price to decimal: exponent is not numeric")
	assert.Equal(t, expected, actual)
}

func TestSMAOutOfRange(t *testing.T) {
	expected := ""
	actual, err := SMA([]string{"1.0", "-3.14159", "3.0"})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "price value out of range: '-3.14159'")
	assert.Equal(t, expected, actual)
}

func TestSMAEmptyArray(t *testing.T) {
	expected := ""
	actual, err := SMA([]string{})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "empty price array")
	assert.Equal(t, expected, actual)
}
