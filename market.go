package kucoin

import (
	"encoding/json"
	"strings"
)

// GetTicker include only the inside (i.e. best) bid and ask data , last price and last trade size.
// docs: https://docs.kucoin.com/#get-ticker
func (c *Client) GetTicker(symbol string) (ticker Ticker, err error) {
	r, err := c.do("GET", "/api/v1/market/orderbook/level1", doArgs("symbol", strings.ToUpper(symbol)), false, 0)
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
