package kucoin

import (
	"testing"
)

const (
	APIKey     = "YOUR_API_KEY"
	SecretKey  = "YOUR_SECRET_KEY"
	Passphrase = "YOUR_PASSPHRASE"
)

func TestTime(t *testing.T) {
	k := New(APIKey, SecretKey, Passphrase)
	if timestamp, err := k.Time(); err != nil {
		t.Fatal(err)
	} else {
		t.Logf("%d", timestamp)
	}
}

func TestGetTicker(t *testing.T) {
	k := New(APIKey, SecretKey, Passphrase)
	if ticker, err := k.GetTicker("ETH-BTC"); err != nil {
		t.Fatal(err)
	} else {
		t.Logf("%v", ticker)
	}
}

func TestListMergedDealtOrders(t *testing.T) {
	k := New(APIKey, SecretKey, Passphrase)
	if ret, err := k.ListMergedDealtOrders("ETH-USDT", "BUY", 20, 1, 0, 0); err != nil {
		t.Fatal(err)
	} else {
		t.Logf("%d", len(ret.Datas))
	}
}
