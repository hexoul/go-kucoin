package kucoin

import (
	"testing"
)

const (
	APIKey    = "YOUR_API_KEY"
	SecretKey = "YOUR_SECRET_KEY"
)

func TestGetBalance(t *testing.T) {
	k := New(APIKey, SecretKey)
	if bal, err := k.GetCoinBalance("BTC"); err != nil {
		t.Fatal(err)
	} else {
		t.Logf("%f %f", bal.FreezeBalance, bal.Balance)
	}
}
