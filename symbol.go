package kucoin

// Ticker struct
type Ticker struct {
	Sequence    string `json:"sequence"`
	Price       string `json:"price"`
	Size        string `json:"size"`
	BestBid     string `json:"bestBid"`
	BestBidSize string `json:"bestBidSize"`
	BestAsk     string `json:"bestAsk"`
	BestAskSize string `json:"bestAskSize"`
}

type rawTicker struct {
	Code string `json:"code"`
	Data Ticker `json:"data"`
}

// Symbol struct represents kucoin data model.
type Symbol struct {
	CoinType      string  `json:"coinType"`
	Trading       bool    `json:"trading"`
	Symbol        string  `json:"symbol"`
	LastDealPrice float64 `json:"lastDealPrice,omitempty"`
	Buy           float64 `json:"buy,omitempty"`
	Sell          float64 `json:"sell,omitempty"`
	Change        float64 `json:"change,omitempty"`
	CoinTypePair  string  `json:"coinTypePair"`
	Sort          int     `json:"sort"`
	FeeRate       float64 `json:"feeRate"`
	VolValue      float64 `json:"volValue"`
	High          float64 `json:"high,omitempty"`
	Datetime      int64   `json:"datetime"`
	Vol           float64 `json:"vol"`
	Low           float64 `json:"low,omitempty"`
	ChangeRate    float64 `json:"changeRate,omitempty"`
	Stick         bool    `json:"stick,omitempty"`
	Fav           bool    `json:"fav,omitempty"`
}

type rawSymbols struct {
	Success   bool     `json:"success"`
	Code      string   `json:"code"`
	Msg       string   `json:"msg"`
	Timestamp int64    `json:"timestamp"`
	Data      []Symbol `json:"data"`
}

type rawSymbol struct {
	Success   bool   `json:"success"`
	Code      string `json:"code"`
	Msg       string `json:"msg"`
	Timestamp int64  `json:"timestamp"`
	Data      Symbol `json:"data"`
}
