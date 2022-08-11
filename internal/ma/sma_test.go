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

func TestSMAInvalidPrice(t *testing.T) {
	expected := ""
	actual, err := SMA([]string{"1.0", "not-a-price", "3.0"})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "invalid price value: 'not-a-price'")
	assert.Equal(t, expected, actual)
}

func TestSMAOutOfRange(t *testing.T) {
	expected := ""
	actual, err := SMA([]string{"1.0", "-3.14159", "3.0"})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "price value out of range: '-3.14159'")
	assert.Equal(t, expected, actual)
}
