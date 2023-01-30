package webapi

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

type (
	headerOptController func(http.Header)
	queryOptController  func(url.Values)
)

func makeGetRequest(url string, headers ...headerOptController) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	for _, c := range headers {
		c(req.Header)
	}
	if err != nil {
		return nil, err
	}
	client := http.Client{
		Timeout: 30 * time.Second,
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, err
}

func makePostRequest(url string, body map[string]string, headers ...headerOptController) (*http.Response, error) {
	postBody, _ := json.Marshal(body)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(postBody))
	if err != nil {
		return nil, err
	}
	for _, c := range headers {
		c(req.Header)
	}
	client := http.Client{
		Timeout: 30 * time.Second,
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, err
}

func addQueryParams(base string, opts ...queryOptController) (string, error) {
	u, err := url.Parse(base)
	if err != nil {
		return "", err
	}
	q := u.Query()
	for _, c := range opts {
		c(q)
	}
	u.RawQuery = q.Encode()
	return u.String(), nil
}
