package builder

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"shopping-cart/util"
	"strings"
)

type HTTPClient[T any] struct {
	method   string
	url      string
	headers  map[string]string
	formData url.Values
}

func NewHttpClient[T any]() *HTTPClient[T] {
	return &HTTPClient[T]{
		formData: make(url.Values),
	}
}

func (b *HTTPClient[T]) WithMethodPost() *HTTPClient[T] {
	b.method = "POST"
	return b
}

func (b *HTTPClient[T]) WithMethodGet() *HTTPClient[T] {
	b.method = "GET"
	return b
}

func (b *HTTPClient[T]) WithURL(url string) *HTTPClient[T] {
	b.url = url
	return b
}

func (b *HTTPClient[T]) WithFormData(key, value string) *HTTPClient[T] {
	b.formData.Set(key, value)
	return b
}

func (b *HTTPClient[T]) SetHeader(key, value string) *HTTPClient[T] {
	b.headers[key] = value
	return b
}

func (b *HTTPClient[T]) UserHeaderFormUrlencoded() *HTTPClient[T] {
	b.headers["Content-Type"] = "application/x-www-form-urlencoded"
	return b
}

func (b *HTTPClient[T]) Build(dto *T) error {
	client := &http.Client{}

	req, err := http.NewRequest(b.method, b.url, strings.NewReader(b.formData.Encode()))
	if err != nil {
		return err
	}

	for key, value := range b.headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return util.ParseJSONResponse(body, dto)
}
