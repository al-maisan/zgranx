package huobi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/shopspring/decimal"
)

type Order struct {
	Symbol        string          `json:"symbol"`
	Source        string          `json:"source"`
	Price         decimal.Decimal `json:"price"`
	CreatedAt     int             `json:"created-at"`
	Amount        decimal.Decimal `json:"amount"`
	AccountId     int             `json:"account-id"`
	ClientOrderId string          `json:"client-order-id"`
	FilledAmount  decimal.Decimal `json:"filled-amount"`
	FilledFees    decimal.Decimal `json:"filled-fees"`
	Id            int             `json:"id"`
	State         string          `json:"state"`
	Type          string          `json:"type"`
}

type CancelFail struct {
	ErrorMessage  string `json:"err-msg"`
	OrderId       string `json:"order-id"`
	ClientOrderId string `json:"client-order-id"`
}
type CancelData struct {
	Succeeded []string     `json:"success"`
	Failed    []CancelFail `json:"failed"`
}

func GetOpenOrders(apiKey, apiSecret string) ([]Order, error) {
	var oor []Order
	log.Info("getting open orders from huobi")
	ap := "/v1/order/openOrders"
	body, err := doReq(apiKey, apiSecret, http.MethodGet, domain, ap, nil)
	if err != nil {
		log.Errorf("failed to get open orders, %v", err)
		return nil, err
	}
	od, err := parseOpenOrders(body)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	oor = append(oor, od...)

	return oor, nil
}

func parseOpenOrders(body []byte) ([]Order, error) {
	type or struct {
		Status string  `json:"status"`
		Data   []Order `json:"data"`
	}
	od := or{}
	err := json.Unmarshal(body, &od)
	if err != nil {
		log.Error("failed to parse: ", string(body))
		return nil, err
	}
	if od.Status != "ok" {
		err = fmt.Errorf("get open orders from huobi failed, '%s'", string(body))
		log.Error(err)
		return nil, err
	}
	log.Infof("found %d open orders", len(od.Data))
	return od.Data, nil
}

func CancelOrders(apiKey, apiSecret string, ids []string) (*CancelData, error) {
	type cmo struct {
		OIDs []string `json:"order-ids"`
	}
	log.Infof("canceling %d orders on huobi", len(ids))
	ap := "/v1/order/orders/batchcancel"
	input, err := json.Marshal(cmo{OIDs: ids})
	if err != nil {
		log.Errorf("failed to prepare order cancellation input data, %v", err)
		return nil, err
	}
	body, err := doReq(apiKey, apiSecret, http.MethodPost, domain, ap, input)
	if err != nil {
		log.Errorf("failed to cancel orders, %v", err)
		return nil, err
	}
	return parseCancelOrders(body)
}

func parseCancelOrders(body []byte) (*CancelData, error) {
	type or struct {
		Status string     `json:"status"`
		Data   CancelData `json:"data"`
	}
	od := or{}
	err := json.Unmarshal(body, &od)
	if err != nil {
		log.Error("failed to parse: ", string(body))
		return nil, err
	}
	if od.Status != "ok" {
		err = fmt.Errorf("canceling orders on huobi failed, '%s'", string(body))
		log.Error(err)
		return nil, err
	}
	log.Infof("order cancellation: %d succeeded, %d failed", len(od.Data.Succeeded), len(od.Data.Failed))
	return &od.Data, nil
}

func PlaceOrder(apiKey, apiSecret, accountId, symbol, otype, amount, price, clientOrderId string) (string, error) {
	type pos struct {
		AccountId     string `json:"account-id"`
		Symbol        string `json:"symbol"`
		Type          string `json:"type"`
		Amount        string `json:"amount"`
		Price         string `json:"price"`
		ClientOrderId string `json:"client-order-id"`
	}
	log.Info("placing an order on huobi")
	ap := "/v1/order/orders/place"
	input, err := json.Marshal(pos{AccountId: accountId, Symbol: symbol, Type: otype, Amount: amount, Price: price, ClientOrderId: clientOrderId})
	if err != nil {
		log.Errorf("failed to prepare order placement input data, %v", err)
		return "", err
	}
	body, err := doReq(apiKey, apiSecret, http.MethodPost, domain, ap, input)
	if err != nil {
		log.Errorf("failed to place order, %v", err)
		return "", err
	}
	return parsePlaceOrder(body)
}

func parsePlaceOrder(body []byte) (string, error) {
	type por struct {
		Status  string `json:"status"`
		OrderId string `json:"data"`
	}
	od := por{}
	err := json.Unmarshal(body, &od)
	if err != nil {
		log.Error("failed to parse: ", string(body))
		return "", err
	}
	if od.Status != "ok" {
		err = fmt.Errorf("placing order on huobi failed, '%s'", string(body))
		log.Error(err)
		return "", err
	}
	log.Infof("order placed, id: %s", od.OrderId)
	return od.OrderId, nil
}
