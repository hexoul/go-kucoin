# Kucoin API Client
[![GoDoc](https://godoc.org/github.com/hexoul/go-kucoin?status.svg)](https://godoc.org/github.com/hexoul/go-kucoin) [![Go Report Card](https://goreportcard.com/badge/github.com/hexoul/go-kucoin)](https://goreportcard.com/report/github.com/hexoul/go-kucoin)

> Kucoin API v2 Client

## Usage

```golang
package main

import (
	"github.com/hexoul/go-kucoin"
)

func main() {
	// Set your own API key and secret
	k := kucoin.New("API_KEY", "API_SECRET")
	// Call resource
	k.GetCoinBalance("BTC")
}
```

## Features

| API Resource | Type | Done  |
| -------------| ----- | ----- |
| Time | Open | ✔ |
| Ticker | Open | ✔ |
| List accounts | Auth | - |
| Deposit list | Auth | - |

## Reference
- [Kucoin API](https://kucoinapidocs.docs.apiary.io/)
