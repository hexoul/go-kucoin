package kucoin

import (
	"testing"
)

func TestListFills(t *testing.T) {
	k := GetInstanceWithKey(APIKey, SecretKey, Passphrase)
	if ret, err := k.ListFills(&Options{}); err != nil {
		t.Fatal(err)
	} else {
		for _, v := range ret {
			t.Log(v.Symbol, v.OrderID, v.TradeID, v.Funds)
		}
	}
}
