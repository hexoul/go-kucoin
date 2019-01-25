package kucoin

import (
	"encoding/json"
	"strings"
)

// ListAccounts get a list of accounts.
// doc: https://docs.kucoin.com/#list-accounts
func (b *Kucoin) ListAccounts(currency, accountType string) (ret []Account, err error) {
	timestamp, err := b.Time()
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
	r, err := b.client.do("GET", "/api/v1/accounts", payload, true, timestamp)
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

//----------------------- WILL BE DEPRECATED BELOW ---------------------------//

// UserInfo struct represents kucoin data model.
type UserInfo struct {
	ReferrerCode             string      `json:"referrer_code"`
	PhotoCredentialValidated bool        `json:"photoCredentialValidated"`
	VideoValidated           bool        `json:"videoValidated"`
	Language                 string      `json:"language"`
	Currency                 string      `json:"currency"`
	Oid                      string      `json:"oid"`
	BaseFeeRate              float64     `json:"baseFeeRate"`
	HasCredential            bool        `json:"hasCredential"`
	CredentialNumber         string      `json:"credentialNumber"`
	PhoneValidated           bool        `json:"phoneValidated"`
	Phone                    string      `json:"phone"`
	CredentialValidated      bool        `json:"credentialValidated"`
	GoogleTwoFaBinding       bool        `json:"googleTwoFaBinding"`
	Nickname                 interface{} `json:"nickname"`
	Name                     string      `json:"name"`
	HasTradePassword         bool        `json:"hasTradePassword"`
	EmailValidated           bool        `json:"emailValidated"`
	Email                    string      `json:"email"`
	LoginRecord              struct {
		Last struct {
			IP      string      `json:"ip"`
			Context interface{} `json:"context"`
			Time    int64       `json:"time"`
		} `json:"last"`
		Current struct {
			IP      string      `json:"ip"`
			Context interface{} `json:"context"`
			Time    int64       `json:"time"`
		} `json:"current"`
	} `json:"loginRecord"`
}

type rawUserInfo struct {
	Data UserInfo `json:"data"`
}

// AccountHistory struct represents kucoin data model.
type AccountHistory struct {
	Datas []struct {
		Fee             float64     `json:"fee"`
		Oid             string      `json:"oid"`
		Type            string      `json:"type"`
		Amount          float64     `json:"amount"`
		Remark          string      `json:"remark"`
		Status          string      `json:"status"`
		Address         string      `json:"address"`
		Context         string      `json:"context"`
		UserOid         string      `json:"userOid"`
		CoinType        string      `json:"coinType"`
		CreatedAt       int64       `json:"createdAt"`
		DeletedAt       interface{} `json:"deletedAt"`
		UpdatedAt       int64       `json:"updatedAt"`
		OuterWalletTxid interface{} `json:"outerWalletTxid"`
	} `json:"datas"`
	Total           int         `json:"total"`
	Limit           int         `json:"limit"`
	PageNos         int         `json:"pageNos"`
	CurrPageNo      int         `json:"currPageNo"`
	NavigatePageNos []int       `json:"navigatePageNos"`
	CoinType        string      `json:"coinType"`
	Type            interface{} `json:"type"`
	UserOid         string      `json:"userOid"`
	Status          interface{} `json:"status"`
	FirstPage       bool        `json:"firstPage"`
	LastPage        bool        `json:"lastPage"`
	StartRow        int         `json:"startRow"`
}

type rawAccountHistory struct {
	Success bool           `json:"success"`
	Code    string         `json:"code"`
	Data    AccountHistory `json:"data"`
}
