package huobi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Account struct {
	ID      int    `json:"id"`
	Type    string `json:"type"`
	SubType string `json:"subtype"`
	State   string `json:"state"`
}

func getAccounts(apiKey, apiSecret string) ([]Account, error) {
	body, err := doReq(apiKey, apiSecret, http.MethodGet, domain, "/v1/account/accounts")

	if err != nil {
		log.Error("failed to get accounts, ", err)
		return nil, err
	}

	type ar struct {
		Status string    `json:"status"`
		Data   []Account `json:"data"`
	}
	return parseAccounts(body)
}

func parseAccounts(body []byte) ([]Account, error) {
	type ar struct {
		Status string    `json:"status"`
		Data   []Account `json:"data"`
	}
	jd := ar{}
	err := json.Unmarshal(body, &jd)
	if err != nil {
		log.Error("failed to parse: ", string(body))
		return nil, err
	}
	if jd.Status != "ok" {
		err = fmt.Errorf("failed to get accounts from huobi, '%s'", string(body))
		log.Error(err)
		return nil, err
	}

	return jd.Data, nil
}
