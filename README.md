# Kucoin API Client
[![GoDoc](https://godoc.org/github.com/hexoul/go-kucoin?status.svg)](https://godoc.org/github.com/hexoul/go-kucoin) [![Go Report Card](https://goreportcard.com/badge/github.com/hexoul/go-kucoin)](https://goreportcard.com/report/github.com/hexoul/go-kucoin)


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
| Tick (symbols) | Open | ✔ |
| Get coin info | Open | ✔ |
| List coins | Open | ✔ |
| Tick (symbols) for logged user | Auth | - |
| Get coin deposit address | Auth | - |
| Get balance of coin | Auth | ✔ |
| Create an order | Auth | - |
| Get user info | Auth | - |
| List active orders (Both map and array) | Auth | - |
| List deposit & withdrawal records | Auth | - |
| List dealt orders (Both Specific and Merged) | Auth | ✔ |
| Order details | Auth | - |
| Create withdrawal apply | Auth | - |
| Cancel withdrawal | Auth | - |
| Cancel orders | Auth | - |
| Cancel all orders | Auth | - |
| Order books | Auth | - |

## Reference
- [Kucoin API](https://kucoinapidocs.docs.apiary.io/)
