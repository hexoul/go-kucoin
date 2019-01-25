package kucoin

// Hold struct
type Hold struct {
	Currency   string `json:"currency"`
	HoldAmount string `json:"holdAmount"`
	BizType    string `json:"bizType"`
	OrderID    string `json:"orderId"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
}

// HoldList struct
type HoldList struct {
	CurrentPage int    `json:"currentPage"`
	PageSize    int    `json:"pageSize"`
	TotalNum    int    `json:"totalNum"`
	TotalPage   int    `json:"totalPage"`
	Items       []Hold `json:"items"`
}

type rawHoldList struct {
	Data HoldList `json:"data"`
}

// Account struct
type Account struct {
	ID        string `json:"id"`
	Currency  string `json:"currency"`
	Type      string `json:"type"`
	Balance   string `json:"balance"`
	Available string `json:"available"`
	Holds     string `json:"holds"`
}

// Accounts struct
type Accounts struct {
	Data []Account `json:"data"`
}
