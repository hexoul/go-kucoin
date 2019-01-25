package kucoin

// Coin struct represents kucoin data model.
type Coin struct {
	WithdrawMinFee    float64     `json:"withdrawMinFee"`
	WithdrawMinAmount float64     `json:"withdrawMinAmount"`
	WithdrawFeeRate   float64     `json:"withdrawFeeRate"`
	ConfirmationCount int         `json:"confirmationCount"`
	WithdrawRemark    string      `json:"withdrawRemark"`
	InfoURL           interface{} `json:"infoUrl"`
	Name              string      `json:"name"`
	TradePrecision    int         `json:"tradePrecision"`
	DepositRemark     interface{} `json:"depositRemark"`
	EnableWithdraw    bool        `json:"enableWithdraw"`
	EnableDeposit     bool        `json:"enableDeposit"`
	Coin              string      `json:"coin"`
}

type rawCoins struct {
	Success   bool   `json:"success"`
	Code      string `json:"code"`
	Msg       string `json:"msg"`
	Timestamp int64  `json:"timestamp"`
	Data      []Coin `json:"data"`
}

type rawCoin struct {
	Success   bool   `json:"success"`
	Code      string `json:"code"`
	Msg       string `json:"msg"`
	Timestamp int64  `json:"timestamp"`
	Data      Coin   `json:"data"`
}

// CoinBalance struct represents kucoin data model.
type CoinBalance struct {
	CoinType      string  `json:"coinType"`
	Balance       float64 `json:"balance"`
	FreezeBalance float64 `json:"freezeBalance"`
}

type rawCoinBalances struct {
	Success bool          `json:"success"`
	Code    string        `json:"code"`
	Data    []CoinBalance `json:"data"`
}

type rawCoinBalance struct {
	Success bool        `json:"success"`
	Code    string      `json:"code"`
	Data    CoinBalance `json:"data"`
}

// SpecificDealtOrder struct represents kucoin data model.
type SpecificDealtOrder struct {
	Datas []struct {
		Oid       string  `json:"oid"`
		DealPrice float64 `json:"dealPrice"`
		OrderOid  string  `json:"orderOid"`
		Direction string  `json:"direction"`
		Amount    float64 `json:"amount"`
		DealValue float64 `json:"dealValue"`
		CreatedAt int64   `json:"createdAt"`
	} `json:"datas"`
	Total           int         `json:"total"`
	Limit           int         `json:"limit"`
	PageNos         int         `json:"pageNos"`
	CurrPageNo      int         `json:"currPageNo"`
	NavigatePageNos []int       `json:"navigatePageNos"`
	UserOid         string      `json:"userOid"`
	Direction       interface{} `json:"direction"`
	StartRow        int         `json:"startRow"`
	FirstPage       bool        `json:"firstPage"`
	LastPage        bool        `json:"lastPage"`
}

type rawSpecificDealtOrder struct {
	Success bool               `json:"success"`
	Code    string             `json:"code"`
	Msg     string             `json:"msg"`
	Data    SpecificDealtOrder `json:"data"`
}

// MergedDealtOrder struct represents kucoin data model.
type MergedDealtOrder struct {
	Total int `json:"total"`
	Datas []struct {
		CreatedAt     int64   `json:"createdAt"`
		Amount        float64 `json:"amount"`
		DealValue     float64 `json:"dealValue"`
		DealPrice     float64 `json:"dealPrice"`
		Fee           float64 `json:"fee"`
		FeeRate       float64 `json:"feeRate"`
		Oid           string  `json:"oid"`
		OrderOid      string  `json:"orderOid"`
		CoinType      string  `json:"coinType"`
		CoinTypePair  string  `json:"coinTypePair"`
		Direction     string  `json:"direction"`
		DealDirection string  `json:"dealDirection"`
	} `json:"datas"`
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

type rawMergedDealtOrder struct {
	Success   bool             `json:"success"`
	Code      string           `json:"code"`
	Msg       string           `json:"msg"`
	Timestamp int64            `json:"timestamp"`
	Data      MergedDealtOrder `json:"data"`
}

// ActiveMapOrder struct represents kucoin data model.
type ActiveMapOrder struct {
	SELL []struct {
		Oid           string      `json:"oid"`
		Type          string      `json:"type"`
		UserOid       interface{} `json:"userOid"`
		CoinType      string      `json:"coinType"`
		CoinTypePair  string      `json:"coinTypePair"`
		Direction     string      `json:"direction"`
		Price         float64     `json:"price"`
		DealAmount    float64     `json:"dealAmount"`
		PendingAmount float64     `json:"pendingAmount"`
		CreatedAt     int64       `json:"createdAt"`
		UpdatedAt     int64       `json:"updatedAt"`
	} `json:"SELL"`
	BUY []struct {
		Oid           string      `json:"oid"`
		Type          string      `json:"type"`
		UserOid       interface{} `json:"userOid"`
		CoinType      string      `json:"coinType"`
		CoinTypePair  string      `json:"coinTypePair"`
		Direction     string      `json:"direction"`
		Price         float64     `json:"price"`
		DealAmount    float64     `json:"dealAmount"`
		PendingAmount float64     `json:"pendingAmount"`
		CreatedAt     int64       `json:"createdAt"`
		UpdatedAt     int64       `json:"updatedAt"`
	} `json:"BUY"`
}

type rawActiveMapOrder struct {
	Success   bool           `json:"success"`
	Code      string         `json:"code"`
	Msg       string         `json:"msg,omitempty"`
	Timestamp int64          `json:"timestamp,omitempty"`
	Data      ActiveMapOrder `json:"data"`
}

// ActiveOrder struct represents kucoin data model.
type ActiveOrder struct {
	SELL [][]interface{} `json:"SELL"`
	BUY  [][]interface{} `json:"BUY"`
}

type rawActiveOrder struct {
	Comment   string      `json:"_comment"`
	Success   bool        `json:"success"`
	Code      string      `json:"code"`
	Msg       string      `json:"msg,omitempty"`
	Timestamp int64       `json:"timestamp,omitempty"`
	Data      ActiveOrder `json:"data"`
}

// Order structs represents kucoin data model.
type Order struct {
	OrderOid string `json:"orderOid"`
}

type rawOrder struct {
	Success bool   `json:"success"`
	Code    string `json:"code"`
	Msg     string `json:"msg"`
	Data    Order  `json:"data"`
}

// OrderDetails structs represents kucoin data model.
type OrderDetails struct {
	CoinType         string  `json:"coinType"`
	DealValueTotal   float64 `json:"dealValueTotal"`
	DealPriceAverage float64 `json:"dealPriceAverage"`
	FeeTotal         float64 `json:"feeTotal"`
	UserOid          string  `json:"userOid"`
	DealAmount       float64 `json:"dealAmount"`
	DealOrders       struct {
		Total     int  `json:"total"`
		FirstPage bool `json:"firstPage"`
		LastPage  bool `json:"lastPage"`
		Datas     []struct {
			Amount    float64 `json:"amount"`
			DealValue float64 `json:"dealValue"`
			Fee       float64 `json:"fee"`
			DealPrice float64 `json:"dealPrice"`
			FeeRate   float64 `json:"feeRate"`
		} `json:"datas"`
		CurrPageNo int `json:"currPageNo"`
		Limit      int `json:"limit"`
		PageNos    int `json:"pageNos"`
	} `json:"dealOrders"`
	CoinTypePair  string  `json:"coinTypePair"`
	OrderPrice    float64 `json:"orderPrice"`
	Type          string  `json:"type"`
	OrderOid      string  `json:"orderOid"`
	PendingAmount float64 `json:"pendingAmount"`
}

type rawOrderDetails struct {
	Success   bool         `json:"success"`
	Code      string       `json:"code"`
	Msg       string       `json:"msg"`
	Timestamp int64        `json:"timestamp"`
	Data      OrderDetails `json:"data"`
}

// OrdersBook struct represents kucoin data model.
type OrdersBook struct {
	Comment string      `json:"_comment"`
	SELL    [][]float64 `json:"SELL"`
	BUY     [][]float64 `json:"BUY"`
}

type rawOrdersBook struct {
	Success bool       `json:"success"`
	Code    string     `json:"code"`
	Msg     string     `json:"msg"`
	Data    OrdersBook `json:"data"`
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

// Withdrawal struct represents kucoin data model.
type Withdrawal struct {
}

type rawWithdrawal struct {
	Success   bool       `json:"success,omitempty"`
	Code      string     `json:"code,omitempty"`
	Msg       string     `json:"msg,omitempty"`
	Timestamp int64      `json:"timestamp,omitempty"`
	Data      Withdrawal `json:"data,omitempty"`
}
