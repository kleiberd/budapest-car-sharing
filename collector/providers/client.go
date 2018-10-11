package providers

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36"
)

type Client struct {
	client *http.Client

	endpointURL *url.URL
	userAgent   string
	referer     string
	header      map[string]string
}

func NewClient(endpointURL string, referer string, header map[string]string) *Client {
	httpClient := http.DefaultClient

	EndpointURL, _ := url.Parse(endpointURL)

	c := &Client{
		client:      httpClient,
		endpointURL: EndpointURL,
		userAgent:   userAgent,
		referer:     referer,
		header:      header,
	}

	return c
}

func (c *Client) SendRequest() ([]byte, error) {
	req, err := c.getRequest()
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func (c *Client) getRequest() (*http.Request, error) {
	req, err := http.NewRequest("GET", c.endpointURL.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "hu,en-US;q=0.9,en;q=0.8")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Referer", c.referer)

	if nil != c.header {
		for key, value := range c.header {
			req.Header.Set(key, value)
		}
	}

	return req, nil
}
