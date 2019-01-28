package kucoin

import (
	"encoding/json"
)

// ListFills get a list of recent fills.
// doc: https://docs.kucoin.com/#fills
func (c *Client) ListFills(options *Options) (ret []Fill, err error) {
	timestamp, err := c.Time()
	if err != nil {
		return
	}
	r, err := c.do("GET", "/api/v1/fills", parseOptions(options), true, timestamp)
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
	var res rawFillList
	err = json.Unmarshal(r, &res)
	ret = res.Data.Items
	return
}
