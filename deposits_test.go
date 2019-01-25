package kucoin

import (
	"testing"
)

func TestGetDepositList(t *testing.T) {
	k := New(APIKey, SecretKey, Passphrase)
	if depositList, err := k.GetDepositList("", ""); err != nil {
		t.Fatal(err)
	} else {
		t.Logf("%v\n", depositList)
	}
}
