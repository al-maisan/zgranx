package huobi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const atd = `
{
    "status": "ok",
    "data": [
        {
            "id": 10000001,
            "type": "spot",
            "subtype": "",
            "state": "working"
        },
        {
            "id": 10000002,
            "type": "otc",
            "subtype": "",
            "state": "working"
        },
        {
            "id": 10000003,
            "type": "point",
            "subtype": "",
            "state": "working"
        }
    ]
}
`

func TestParseAccounts(t *testing.T) {
	expected := []Account{
		{
			ID:      10000001,
			Type:    "spot",
			SubType: "",
			State:   "working",
		},
		{
			ID:      10000002,
			Type:    "otc",
			SubType: "",
			State:   "working",
		},
		{
			ID:      10000003,
			Type:    "point",
			SubType: "",
			State:   "working",
		},
	}

	actual, err := parseAccounts([]byte(atd))
	assert.Nil(t, err)
	assert.NotNil(t, actual)
	assert.Equal(t, expected, actual)
}
