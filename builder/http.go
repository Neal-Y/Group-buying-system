package builder

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type HTTPClient struct {
	method   string
	url      string
	formData url.Values
}

func NewHttpClient() *HTTPClient {
	return &HTTPClient{
		formData: make(url.Values),
	}
}

func (b *HTTPClient) WithMethod(method string) *HTTPClient {
	b.method = method
	return b
}

func (b *HTTPClient) WithURL(url string) *HTTPClient {
	b.url = url
	return b
}

func (b *HTTPClient) WithFormData(key, value string) *HTTPClient {
	b.formData.Set(key, value)
	return b
}

func (b *HTTPClient) Build() ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest(b.method, b.url, strings.NewReader(b.formData.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
