// Package kucoin is an implementation of the Kucoin API in Golang.
package kucoin

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
)

const (
	// v1
	// kucoinURL = "https://api.kucoin.com/v1/"
	// v2
	kucoinURL = "https://openapi-v2.kucoin.com"
	// sandbox
	// kucoinURL = "https://openapi-sandbox.kucoin.com"
)

// Client struct
type Client struct {
	accessKey  string
	secretKey  string
	passphrase string
	bSecret    []byte
}

var (
	instance   *Client
	once       sync.Once
	accessKey  string
	secretKey  string
	passphrase string
)

func init() {
	for _, val := range os.Args {
		arg := strings.Split(val, "=")
		if len(arg) < 2 {
			continue
		} else if arg[0] == "-kucoin:accesskey" {
			accessKey = arg[1]
		} else if arg[0] == "-kucoin:secretkey" {
			secretKey = arg[1]
		} else if arg[0] == "-kucoin:passphrase" {
			passphrase = arg[1]
		}
	}
}

// GetInstance returns singleton
func GetInstance() *Client {
	once.Do(func() {
		if accessKey == "" || secretKey == "" || passphrase == "" {
			panic("KEYS FOR BOTH ACCESS AND SECRET AND PASSPHRASE REQUIRED")
		}
		instance = &Client{
			accessKey:  accessKey,
			secretKey:  secretKey,
			passphrase: passphrase,
			bSecret:    []byte(secretKey),
		}
	})
	return instance
}

// GetInstanceWithKey returns singleton
func GetInstanceWithKey(accessKey, secretKey, passphrase string) *Client {
	once.Do(func() {
		if accessKey == "" || secretKey == "" || passphrase == "" {
			panic("KEYS FOR BOTH ACCESS AND SECRET REQUIRED")
		}
		instance = &Client{
			accessKey:  accessKey,
			secretKey:  secretKey,
			passphrase: passphrase,
			bSecret:    []byte(secretKey),
		}
	})
	return instance
}

func doArgs(args ...string) map[string]string {
	m := make(map[string]string)
	var lastK = ""
	for i, v := range args {
		if i&1 == 0 {
			lastK = v
		} else {
			m[lastK] = v
		}
	}
	return m
}

// handleErr gets JSON response from livecoin API en deal with error.
func handleErr(r interface{}) error {
	switch v := r.(type) {
	case map[string]interface{}:
		if code := r.(map[string]interface{})["code"]; code != "200" && code != "200000" {
			errorMessage := r.(map[string]interface{})["msg"]
			return errors.New(errorMessage.(string))
		}
	case []interface{}:
		return nil
	default:
		return fmt.Errorf("don't recognized type %T", v)
	}

	return nil
}

// Time get the API server time.
// doc: https://docs.kucoin.com/#time
func (c *Client) Time() (timestamp int64, err error) {
	r, err := c.do("GET", "/api/v1/timestamp", nil, false, 0)
	if err != nil {
		return
	}
	var response interface{}
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}
	data := response.(map[string]interface{})["data"].(float64)
	timestamp = int64(data)
	return
}
