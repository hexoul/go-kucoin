package kucoin

import (
	"testing"
)

const (
	APIKey    = "YOUR_API_KEY"
	SecretKey = "YOUR_SECRET_KEY"
)

func TestGetSymbol(t *testing.T) {
	k := New(APIKey, SecretKey)
	if symbol, err := k.GetSymbol("ETH-USDT"); err != nil {
		t.Fatal(err)
	} else {
		t.Logf("%v", symbol.Datetime)
	}
}

func TestGetBalance(t *testing.T) {
	k := New(APIKey, SecretKey)
	if bal, err := k.GetCoinBalance("BTC"); err != nil {
		t.Fatal(err)
	} else {
		t.Logf("%f %f", bal.FreezeBalance, bal.Balance)
	}
}

func TestListMergedDealtOrders(t *testing.T) {
	k := New(APIKey, SecretKey)
	if ret, err := k.ListMergedDealtOrders("ETH-USDT", "BUY", 20, 1, 0, 0); err != nil {
		t.Fatal(err)
	} else {
		t.Logf("%d", len(ret.Datas))
	}
}
