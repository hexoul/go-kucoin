# Kucoin API Client
[![GoDoc](https://godoc.org/github.com/hexoul/go-kucoin?status.svg)](https://godoc.org/github.com/hexoul/go-kucoin) [![Go Report Card](https://goreportcard.com/badge/github.com/hexoul/go-kucoin)](https://goreportcard.com/report/github.com/hexoul/go-kucoin)

> Kucoin API v2 Client

## Usage

```golang
import (
	"github.com/hexoul/go-kucoin"
)

func main() {
	// Set your own API key, secret and passphrase
	k := kucoin.New("API_KEY", "API_SECRET", "API_PASSPHRASE")
	// Call resource
	k.GetCoinBalance("BTC")
}
```

## Test

1. Fill your information in `kucoin_test.go`

  ```golang
  const (
    APIKey     = "YOUR_API_KEY"
    SecretKey  = "YOUR_SECRET_KEY"
	Passphrase = "YOUR_PASSPHRASE"
  )
  ```

2. ```go test -v```

## Features

| API Resource | Type | Done  |
| -------------| ----- | ----- |
| Time | Open | ✔ |
| Ticker | Open | ✔ |
| List accounts | Auth | ✔ |
| Deposit list | Auth | ✔ |

## Reference
- [Kucoin API](https://docs.kucoin.com/)
