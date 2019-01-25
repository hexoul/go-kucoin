package kucoin

import (
	"testing"
)

func TestListAccounts(t *testing.T) {
	k := New(APIKey, SecretKey, Passphrase)
	if accounts, err := k.ListAccounts("", ""); err != nil {
		t.Fatal(err)
	} else {
		for _, v := range accounts {
			t.Log(v)
		}
	}
}
