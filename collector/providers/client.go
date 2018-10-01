package providers

import (
	"net/http"
	"net/url"
)

const (
	userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36"
)

type Client struct {
	client      *http.Client
	EndpointURL *url.URL
	UserAgent   string
}

func NewClient(endpointURL string) *Client {
	httpClient := http.DefaultClient

	EndpointURL, _ := url.Parse(endpointURL)

	c := &Client{client: httpClient, EndpointURL: EndpointURL, UserAgent: userAgent}

	return c
}

func (c *Client) GetRequest(header *http.Header) (*http.Request, error) {
	req, err := http.NewRequest("GET", c.EndpointURL.String(), nil)
	if err != nil {
		return nil, err
	}

	if nil != header {
		req.Header = *header
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "hu,en-US;q=0.9,en;q=0.8")
	req.Header.Set("Connection", "keep-alive")

	return req, nil
}

func (c *Client) SendRequest(req *http.Request) (*http.Response, error) {
	return c.client.Do(req)
}
