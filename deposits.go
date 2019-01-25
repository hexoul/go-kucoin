package kucoin

import (
	"encoding/json"
	"strings"
)

// GetDepositList get deposit page list.
// docs: https://docs.kucoin.com/#get-deposit-address
func (c *Client) GetDepositList(currency, status string) (deposits []Deposit, err error) {
	timestamp, err := c.Time()
	if err != nil {
		return
	}
	payload := make(map[string]string)
	if currency != "" {
		payload["currency"] = strings.ToUpper(currency)
	}
	if status == "PROCESSING" || status == "SUCCESS" || status == "FAILURE" {
		payload["status"] = status
	}
	r, err := c.do("GET", "/api/v1/deposits", payload, true, timestamp)
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
