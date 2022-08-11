package ma

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEMA(t *testing.T) {
	expected := "12.80000"
	actual, err := EMA([]string{"14", "13", "14", "12", "13"})
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestEMASingleValue(t *testing.T) {
	expected := "3.444568"
	actual, err := EMA([]string{"3.444567890"})
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestEMAInvalidPrice(t *testing.T) {
	expected := ""
	actual, err := EMA([]string{"1.0", "not-a-price", "3.0"})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "invalid price value: 'not-a-price'")
	assert.Equal(t, expected, actual)
}

func TestEMAOutOfRange(t *testing.T) {
	expected := ""
	actual, err := EMA([]string{"1.0", "-3.14159", "3.0"})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "price value out of range: '-3.14159'")
	assert.Equal(t, expected, actual)
}

func TestEMAEmptyArray(t *testing.T) {
	expected := ""
	actual, err := EMA([]string{})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "empty price array")
	assert.Equal(t, expected, actual)
}
