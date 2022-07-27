package huobi

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"path"
	"time"
)

func doReq(apiKey, apiSecret, method, domain, reqPath string) ([]byte, error) {
	u, err := url.Parse(domain)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	u.Path = path.Join(u.Path, reqPath)
	req, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	params, err := signReq(apiKey, apiSecret, req.Method, domain, u.Path)
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
	var params url.Values
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
