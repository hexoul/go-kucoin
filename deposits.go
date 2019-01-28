package kucoin

import (
	"encoding/json"
)

// GetDepositList get deposit page list.
// docs: https://docs.kucoin.com/#get-deposit-address
func (c *Client) GetDepositList(options *Options) (deposits []Deposit, err error) {
	timestamp, err := c.Time()
	if err != nil {
		return
	}
	r, err := c.do("GET", "/api/v1/deposits", parseOptions(options), true, timestamp)
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
	var rawRes rawDepositList
	err = json.Unmarshal(r, &rawRes)
	deposits = rawRes.Data.Items
	return
}
