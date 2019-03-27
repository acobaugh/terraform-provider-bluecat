package client

import (
	"net/url"
	"net/http"
	"crypto/tls"
	"fmt"
	"bytes"
	"io/ioutil"
	"encoding/json"
)

type Config struct {
	Username string
	Password string
	BaseURL string
	Insecure bool
}

type Client struct {
	Config *Config
	UserAgent string

	URL *url.URL
	hc *http.Client
	token string
}

func NewClient(config *Config) (*Client, error) {
	tc := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: config.Insecure},
	}

	httpClient := &http.Client{Transport: tc}

	url, err := url.Parse(config.BaseURL)
	if err != nil {
		return nil, err
	}

	c := &Client{Config: config, UserAgent: "", URL: url, hc: httpClient}

	return c, nil
}

// request accepts a HTTP method, resource, payload, and query parameters and returns an *http.Response
func (c *Client) request(method string, resource string, payload interface{}, query map[string]string) (*http.Response, error) {
	url := fmt.Sprintf("%s/Services/REST/v1/%s", c.Config.BaseURL, resource)

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	if len(c.token) > 0 {
		req.Header.Set("Authorization", c.token)
	}

	q := req.URL.Query()
	for argName, argVal := range query {
		q.Add(argName, argVal)
	}
	req.URL.RawQuery = q.Encode()
	resp, err := c.hc.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		errBytes, _ := ioutil.ReadAll(resp.Body)
		errResp := string(errBytes)
		return nil, fmt.Errorf("bluecat: request failed with HTTP status code %d\n Full message: %s",
			resp.StatusCode, errResp)
	}

	return resp, nil
}