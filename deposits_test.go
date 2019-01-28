package kucoin

import (
	"testing"
)

func TestGetDepositList(t *testing.T) {
	k := GetInstanceWithKey(APIKey, SecretKey, Passphrase)
	if depositList, err := k.GetDepositList(nil); err != nil {
		t.Fatal(err)
	} else {
		t.Logf("%v\n", depositList)
	}
}
