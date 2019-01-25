package kucoin

import (
	"crypto/hmac"
	"crypto/sha256"
	b64 "encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"
)

type client struct {
	apiKey     string
	apiSecret  string
	passPhrase string
	bSecret    []byte
	httpClient http.Client
	debug      bool
}

func newClient(apiKey, apiSecret, passPhrase string) (c *client) {
	c = &client{
		apiKey:     apiKey,
		apiSecret:  apiSecret,
		passPhrase: passPhrase,
		bSecret:    []byte(apiSecret),
	}
	c.httpClient.Timeout = time.Second * 30
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
func (c *client) do(method, resource string, payload map[string]string, authNeeded bool, nonce int64) ([]byte, error) {
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
		if len(c.apiKey) == 0 || len(c.apiSecret) == 0 || len(c.passPhrase) == 0 {
			return nil, errors.New("API Key, API Secret and Passphrase must be set")
		}
		if nonce == 0 {
			nonce = time.Now().UnixNano() / int64(time.Millisecond)
		}
		req.Header.Add("KC-API-KEY", c.apiKey)
		req.Header.Add("KC-API-PASSPHRASE", c.passPhrase)
		req.Header.Add("KC-API-TIMESTAMP", fmt.Sprintf("%v", nonce))
		req.Header.Add("KC-API-SIGN", c.sign(method, resource, queryString, nonce))
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
	}
	return data, err
}

func computeHmac256(message string, secret []byte) []byte {
	h := hmac.New(sha256.New, secret)
	h.Write([]byte(message))
	return h.Sum(nil)
}

func (c *client) sign(method, resource, queryString string, nonce int64) (signature string) {
	if queryString != "" {
		queryString = "?" + queryString
	}
	strForSign := fmt.Sprintf("%v%s%s%s", nonce, method, resource, queryString)
	fmt.Println(strForSign)
	return b64.StdEncoding.EncodeToString(computeHmac256(strForSign, c.bSecret))
}
