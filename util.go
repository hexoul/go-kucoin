package kucoin

import (
	"crypto/hmac"
	"crypto/sha256"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func parseOptions(options *Options) map[string]string {
	payload := make(map[string]string)
	if options == nil {
		return payload
	}

	if options.Currency != "" {
		payload["currency"] = strings.ToUpper(options.Currency)
	}
	if options.AccountType == "main" || options.AccountType == "trade" {
		payload["type"] = options.AccountType
	}
	if options.Status == "PROCESSING" || options.Status == "SUCCESS" || options.Status == "FAILURE" {
		payload["status"] = options.Status
	}
	return payload
}

func doReq(req *http.Request) (body []byte, err error) {
	requestTimeout := time.Duration(5 * time.Second)
	client := &http.Client{
		Timeout: requestTimeout,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		return nil, err
	}
	return
}

func computeHmac256(message string, secret []byte) []byte {
	h := hmac.New(sha256.New, secret)
	h.Write([]byte(message))
	return h.Sum(nil)
}
