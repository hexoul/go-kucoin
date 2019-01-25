package kucoin

import (
	"encoding/json"
	"fmt"
	"strings"
)

// ListAccounts get a list of accounts.
// doc: https://docs.kucoin.com/#list-accounts
func (c *Client) ListAccounts(currency, accountType string) (ret []Account, err error) {
	timestamp, err := c.Time()
	if err != nil {
		return
	}
	payload := make(map[string]string)
	if currency != "" {
		payload["currency"] = strings.ToUpper(currency)
	}
	if accountType == "main" || accountType == "trade" {
		payload["type"] = accountType
	}
	r, err := c.do("GET", "/api/v1/accounts", payload, true, timestamp)
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
	var accounts Accounts
	err = json.Unmarshal(r, &accounts)
	ret = accounts.Data
	return
}

// GetHolds that Holds are placed on an account for any active orders or pending withdraw requests.
// As an order is filled, the hold amount is updated.
// If an order is canceled, any remaining hold is removed.
// For a withdraw, once it is completed, the hold is removed.
// doc: https://docs.kucoin.com/#get-holds
func (c *Client) GetHolds(accountID string) (ret []Hold, err error) {
	if accountID == "" {
		return
	}
	timestamp, err := c.Time()
	if err != nil {
		return
	}
	r, err := c.do("GET", fmt.Sprintf("/api/v1/accounts/%s/holds", accountID), nil, true, timestamp)
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
	var rawRes rawHoldList
	err = json.Unmarshal(r, &rawRes)
	ret = rawRes.Data.Items
	return
}
