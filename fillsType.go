package kucoin

// Fill struct
type Fill struct {
	Symbol         string `json:"symbol"`
	TradeID        string `json:"tradeId"`
	OrderID        string `json:"orderId"`
	CounterOrderID string `json:"counterOrderId"`
	Side           string `json:"side"`
	Liquidity      string `json:"liquidity"`
	ForceTaker     bool   `json:"forceTaker"`
	Price          string `json:"price"`
	Size           string `json:"size"`
	Funds          string `json:"funds"`
	Fee            string `json:"fee"`
	FeeRate        string `json:"feeRate"`
	FeeCurrency    string `json:"feeCurrency"`
	Stop           string `json:"stop"`
	Type           string `json:"type"`
	CreatedAt      int64  `json:"createdAt"`
}

// FillList struct
type FillList struct {
	CurrentPage int    `json:"currentPage"`
	PageSize    int    `json:"pageSize"`
	TotalNum    int    `json:"totalNum"`
	TotalPage   int    `json:"totalPage"`
	Items       []Fill `json:"items"`
}

type rawFillList struct {
	Data FillList `json:"data"`
}
