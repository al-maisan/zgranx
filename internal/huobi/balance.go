package huobi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

const domain = "https://api.huobi.pro"

var log = logrus.New()

type BalanceData struct {
	Account  int       `json:"id"`
	Type     string    `json:"type"`
	State    string    `json:"state"`
	Balances []Balance `json:"list"`
}

type Balance struct {
	Asset   string          `json:"currency"`
	Type    string          `json:"type"`
	Balance decimal.Decimal `json:"balance"`
}

func GetBalances(apiKey, apiSecret string) ([]BalanceData, error) {
	var bdr []BalanceData
	log.Info("getting accounts..")
	as, err := getAccounts(apiKey, apiSecret)
	if err != nil {
		log.Error("failed to get accounts, ", err)
		return nil, err
	}

	for _, a := range as {
		log.Info("getting balances for account ", a.ID)
		ap := fmt.Sprintf("/v1/account/accounts/%s/balance", a.ID)
		body, err := doReq(apiKey, apiSecret, http.MethodGet, domain, ap)
		if err != nil {
			log.Errorf("failed to get balances for account %d, %v", a.ID, err)
			return nil, err
		}
		bd, err := parseBalance(body)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		bdr = append(bdr, *bd)
	}

	return bdr, nil
}

func parseBalance(body []byte) (*BalanceData, error) {
	type br struct {
		Status string      `json:"status"`
		Data   BalanceData `json:"data"`
	}
	bd := br{}
	err := json.Unmarshal(body, &bd)
	if err != nil {
		log.Error("failed to parse: ", string(body))
		return nil, err
	}
	if bd.Status != "ok" {
		err = fmt.Errorf("get balances from huobi failed, '%s'", string(body))
		log.Error(err)
		return nil, err
	}

	return &bd.Data, nil
}
