// Package kucoin is an implementation of the Kucoin API in Golang.
package kucoin

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

const (
	// v1
	// kucoinURL = "https://api.kucoin.com/v1/"
	// v2
	// kucoinURL = "https://openapi-v2.kucoin.com"
	// sandbox
	kucoinURL = "https://openapi-sandbox.kucoin.com"
)

// New returns an instantiated Kucoin struct.
func New(apiKey, apiSecret, passphrase string) *Kucoin {
	client := newClient(apiKey, apiSecret, passphrase)
	return &Kucoin{client}
}

// NewCustomClient returns an instantiated Kucoin struct with custom http client.
func NewCustomClient(apiKey, apiSecret, passphrase string, httpClient http.Client) *Kucoin {
	client := newClient(apiKey, apiSecret, passphrase)
	client.httpClient = httpClient
	return &Kucoin{client}
}

// NewCustomTimeout returns an instantiated Kucoin struct with custom timeout.
func NewCustomTimeout(apiKey, apiSecret, passphrase string, timeout time.Duration) *Kucoin {
	client := newClient(apiKey, apiSecret, passphrase)
	client.httpClient.Timeout = timeout
	return &Kucoin{client}
}

func doArgs(args ...string) map[string]string {
	m := make(map[string]string)
	var lastK = ""
	for i, v := range args {
		if i&1 == 0 {
			lastK = v
		} else {
			m[lastK] = v
		}
	}
	return m
}

// handleErr gets JSON response from livecoin API en deal with error.
func handleErr(r interface{}) error {
	switch v := r.(type) {
	case map[string]interface{}:
		if code := r.(map[string]interface{})["code"]; code != "200" && code != "200000" {
			errorMessage := r.(map[string]interface{})["msg"]
			return errors.New(errorMessage.(string))
		}
	case []interface{}:
		return nil
	default:
		return fmt.Errorf("don't recognized type %T", v)
	}

	return nil
}

// Kucoin represent a Kucoin client.
type Kucoin struct {
	client *client
}

// SetDebug enables/disables http request/response dump.
func (b *Kucoin) SetDebug(enable bool) {
	b.client.debug = enable
}

// Time get the API server time.
// doc: https://docs.kucoin.com/#time
func (b *Kucoin) Time() (timestamp int64, err error) {
	r, err := b.client.do("GET", "/api/v1/timestamp", nil, false, 0)
	if err != nil {
		return
	}
	var response interface{}
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}
	data := response.(map[string]interface{})["data"].(float64)
	timestamp = int64(data)
	return
}

// GetTicker include only the inside (i.e. best) bid and ask data , last price and last trade size.
// docs: https://docs.kucoin.com/#get-ticker
func (b *Kucoin) GetTicker(symbol string) (ticker Ticker, err error) {
	r, err := b.client.do("GET", "/api/v1/market/orderbook/level1", doArgs("symbol", strings.ToUpper(symbol)), false, 0)
	if err != nil {
		return
	}
	var response interface{}
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}
	var rawRes rawTicker
	err = json.Unmarshal(r, &rawRes)
	ticker = rawRes.Data
	return
}
