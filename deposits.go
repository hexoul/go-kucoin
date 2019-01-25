package kucoin

import (
	"encoding/json"
	"strings"
)

// GetDepositList get deposit page list.
// docs: https://docs.kucoin.com/#get-deposit-address
func (b *Kucoin) GetDepositList(currency, status string) (depositList DepositList, err error) {
	timestamp, err := b.Time()
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
	r, err := b.client.do("GET", "/api/v1/deposits", payload, true, timestamp)
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
	depositList = rawRes.Data
	return
}

// DepositList struct
type DepositList struct {
	CurrentPage int                      `json:"currentPage"`
	PageSize    int                      `json:"pageSize"`
	TotalNum    int                      `json:"totalNum"`
	TotalPage   int                      `json:"totalPage"`
	Items       []map[string]interface{} `json:"items"`
}

type rawDepositList struct {
	Data DepositList `json:"data"`
}
