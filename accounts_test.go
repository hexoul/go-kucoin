package kucoin

import (
	"strconv"
	"testing"
)

func TestListAccounts(t *testing.T) {
	k := New(APIKey, SecretKey, Passphrase)
	if accounts, err := k.ListAccounts("", ""); err != nil {
		t.Fatal(err)
	} else {
		for _, v := range accounts {
			bal, _ := strconv.ParseFloat(v.Balance, 32)
			avail, _ := strconv.ParseFloat(v.Available, 32)
			t.Logf("%s: %s %s %f %f\n", v.Type, v.ID, v.Currency, bal, avail)
		}
	}
}

func TestGetHolds(t *testing.T) {
	k := New(APIKey, SecretKey, Passphrase)
	if holds, err := k.GetHolds("5c45748cef83c7101c066906"); err != nil {
		t.Fatal(err)
	} else {
		for _, v := range holds {
			t.Logf("%v\n", v)
		}
	}
}
