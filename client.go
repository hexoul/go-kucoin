package kucoin

import (
	"crypto/hmac"
	"crypto/sha256"
	b64 "encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"
)

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

// do prepare and process HTTP request to Kucoin API.
/*
	 *  Example
	 *  POST parameters：
	 *    type: BUY
	 *    amount: 10
	 *    price: 1.1
	 *    Arrange the parameters in ascending alphabetical order (lower cases first),
		  then combine them with & (don't urlencode them, don't add ?, don't add extra &),
		  e.g. amount=10&price=1.1&type=BUY
*/
func (c *Client) do(method, resource string, payload map[string]string, authNeeded bool, nonce int64) ([]byte, error) {
	var req *http.Request

	URL, err := url.Parse(kucoinURL)
	if err != nil {
		return nil, err
	}
	URL.Path = path.Join(URL.Path, resource)
	queryString := ""
	if method == "GET" {
		q := URL.Query()
		for key, value := range payload {
			q.Set(key, value)
		}
		URL.RawQuery = q.Encode()
		req, err = http.NewRequest("GET", URL.String(), nil)
		queryString = URL.Query().Encode()
	} else {
		postValues := url.Values{}
		for key, value := range payload {
			postValues.Set(key, value)
		}
		queryString = postValues.Encode()
		req, err = http.NewRequest(
			method, URL.String(), strings.NewReader(
				queryString,
			),
		)
	}
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	// Auth
	if authNeeded {
		if nonce == 0 {
			nonce = time.Now().UnixNano() / int64(time.Millisecond)
		}
		req.Header.Add("KC-API-KEY", c.accessKey)
		req.Header.Add("KC-API-PASSPHRASE", c.passphrase)
		req.Header.Add("KC-API-TIMESTAMP", fmt.Sprintf("%v", nonce))
		req.Header.Add("KC-API-SIGN", c.sign(method, resource, queryString, nonce))
	}

	return doReq(req)
}

func computeHmac256(message string, secret []byte) []byte {
	h := hmac.New(sha256.New, secret)
	h.Write([]byte(message))
	return h.Sum(nil)
}

func (c *Client) sign(method, resource, queryString string, nonce int64) (signature string) {
	if queryString != "" {
		queryString = "?" + queryString
	}
	strForSign := fmt.Sprintf("%v%s%s%s", nonce, method, resource, queryString)
	return b64.StdEncoding.EncodeToString(computeHmac256(strForSign, c.bSecret))
}
