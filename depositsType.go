package kucoin

// Deposit struct
type Deposit struct {
	Address    string `json:"address"`
	Memo       string `json:"memo"`
	Amount     string `json:"amount"`
	Fee        string `json:"fee"`
	Currency   string `json:"currency"`
	IsInner    string `json:"isInner"`
	WalletTxID string `json:"walletTxId"`
	Status     string `json:"status"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
}

// DepositList struct
type DepositList struct {
	CurrentPage int       `json:"currentPage"`
	PageSize    int       `json:"pageSize"`
	TotalNum    int       `json:"totalNum"`
	TotalPage   int       `json:"totalPage"`
	Items       []Deposit `json:"items"`
}

type rawDepositList struct {
	Data DepositList `json:"data"`
}
