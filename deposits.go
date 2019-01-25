package kucoin

// DepositList struct
type DepositList struct {
	CurrentPage string                   `json:"currentPage"`
	PageSize    string                   `json:"pageSize"`
	TotalNum    interface{}              `json:"totalNum"`
	TotalPage   string                   `json:"totalPage"`
	Items       []map[string]interface{} `json:"items"`
}

type rawDepositList struct {
	Data DepositList `json:"data"`
}
