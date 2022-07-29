package huobi

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/alphabot-fi/T-801/internal/proto/exa"
)

func doReq(apiKey, apiSecret, method, domain, reqPath string, in []byte) ([]byte, error) {
	u, err := url.Parse(domain)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	u.Path = path.Join(u.Path, reqPath)
	req, err := http.NewRequest(method, u.String(), bytes.NewBuffer(in))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	params, err := signReq(apiKey, apiSecret, req.Method, u.Host, u.Path)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	req.URL.RawQuery = (*params).Encode()
	c := &http.Client{
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).Dial,
			TLSHandshakeTimeout:   10 * time.Second,
			ResponseHeaderTimeout: 10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := c.Do(req)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error(err)
		log.Error(string(body))
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		if err != nil {
			err = fmt.Errorf("request failed with %d, '%s', also: %w", res.StatusCode, res.Status, err)
		} else {
			err = fmt.Errorf("request failed with %d, '%s'", res.StatusCode, res.Status)
		}
	}
	if err != nil {
		log.Error(err)
		log.Error(string(body))
		return nil, err
	}
	return body, nil
}

func sign(secret, params string) (string, error) {
	mac := hmac.New(sha256.New, []byte(secret))
	_, err := mac.Write([]byte(params))
	if err != nil {
		log.Error(err)
		return "", err
	}
	signature := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(signature), nil
}

func signReq(apiKey, secretKey, method, domain, path string) (*url.Values, error) {
	params := url.Values{}
	params.Set("AccessKeyId", apiKey)
	params.Set("SignatureMethod", "HmacSHA256")
	params.Set("SignatureVersion", "2")
	params.Set("Timestamp", time.Now().UTC().Format("2006-01-02T15:04:05"))
	payload := fmt.Sprintf("%s\n%s\n%s\n%s", method, domain, path, params.Encode())
	signature, err := sign(secretKey, payload)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	params.Set("Signature", signature)
	return &params, nil
}

func Pair2string(p *exa.Pair) string {
	if p != nil {
		return strings.ToLower(p.Base.String()) + strings.ToLower(p.Quote.String())
	}
	return ""
}

func String2pair(s string) (*exa.Pair, error) {
	qs := []string{"usdt", "usdc", "usd", "eur", "jpy", "chf", "cad", "krw"}
	for _, q := range qs {
		if !strings.HasSuffix(s, q) {
			continue
		}
		ss := strings.Split(s, q)
		b, q := ss[0], q
		ba, ok := exa.Asset_value[strings.ToUpper(b)]
		if !ok {
			err := fmt.Errorf("unknown base asset: '%s'", b)
			return nil, err
		}
		qa, ok := exa.Asset_value[strings.ToUpper(q)]
		if !ok {
			err := fmt.Errorf("unknown quote asset: '%s'", q)
			return nil, err
		}
		return &exa.Pair{Base: exa.Asset(ba), Quote: exa.Asset(qa)}, nil
	}
	err := fmt.Errorf("unknown pair: '%s'", s)
	log.Error(err)
	return nil, err
}

func String2state(s string) (exa.OrderState, error) {
	switch s {
	case "created":
		return exa.OrderState(exa.OrderState_CREATED), nil
	case "submitted":
		return exa.OrderState(exa.OrderState_SUBMITTED), nil
	case "partial-filled":
		return exa.OrderState(exa.OrderState_PARTIAL_FILLED), nil
	case "filled":
		return exa.OrderState(exa.OrderState_FILLED), nil
	case "partial-canceled":
		return exa.OrderState(exa.OrderState_PARTIAL_CANCELED), nil
	case "canceling":
		return exa.OrderState(exa.OrderState_CANCELING), nil
	case "canceled":
		return exa.OrderState(exa.OrderState_CANCELED), nil
	}
	err := fmt.Errorf("unknown order state: '%s'", s)
	log.Error(err)
	return 0, err
}

func String2type(s string) (exa.OrderType, error) {
	// buy-market, sell-market, buy-limit, sell-limit, buy-ioc, sell-ioc,
	// buy-limit-maker, sell-limit-maker, buy-stop-limit, sell-stop-limit,
	// buy-limit-fok, sell-limit-fok, buy-stop-limit-fok, sell-stop-limit-fok
	if strings.HasSuffix(s, "-market") {
		return exa.OrderType_MARKET, nil
	}
	if strings.HasSuffix(s, "-limit") {
		return exa.OrderType_LIMIT, nil
	}
	if strings.HasSuffix(s, "-limit-fok") {
		return exa.OrderType_LIMIT_FOK, nil
	}
	if strings.HasSuffix(s, "-ioc") {
		return exa.OrderType_IOC, nil
	}
	err := fmt.Errorf("unknown order type: '%s'", s)
	log.Error(err)
	return 0, err
}

func String2side(s string) (exa.Side, error) {
	// buy-market, sell-market, buy-limit, sell-limit, buy-ioc, sell-ioc,
	// buy-limit-maker, sell-limit-maker, buy-stop-limit, sell-stop-limit,
	// buy-limit-fok, sell-limit-fok, buy-stop-limit-fok, sell-stop-limit-fok
	if strings.HasPrefix(s, "buy-") {
		return exa.Side_BUY, nil
	}
	if strings.HasPrefix(s, "sell-") {
		return exa.Side_SELL, nil
	}
	err := fmt.Errorf("unknown order side: '%s'", s)
	log.Error(err)
	return 0, err
}

func TypeAndSide2string(t exa.OrderType, s exa.Side) string {
	switch t {
	case exa.OrderType_MARKET:
		return strings.ToLower(s.String()) + "-market"
	case exa.OrderType_LIMIT:
		return strings.ToLower(s.String()) + "-limit"
	case exa.OrderType_LIMIT_FOK:
		return strings.ToLower(s.String()) + "-limit-fok"
	case exa.OrderType_IOC:
		return strings.ToLower(s.String()) + "-ioc"
	}
	return ""
}
