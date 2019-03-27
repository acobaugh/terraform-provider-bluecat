package client

import (
	"net/url"
	"net/http"
	"crypto/tls"
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
	client *http.Client
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

	c := &Client{Config: config, UserAgent: "", URL: url, client: httpClient}

	return c, nil
}