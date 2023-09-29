package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type HttpClient interface {
	Execute(method string, url *url.URL, body interface{}) (
		*http.Response,
		error,
	)
}

type HttpClientImpl struct {
	*http.Client
	host  string
	token string
}

func NewHttpClient(basePath string) HttpClient {
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.MaxIdleConns = 100
	transport.MaxIdleConnsPerHost = 100
	transport.MaxConnsPerHost = 100

	return &HttpClientImpl{
		Client: &http.Client{
			Transport: transport,
			Timeout:   time.Second * 10,
		},
		host: basePath,
	}
}

func (c *HttpClientImpl) SetToken(token string) {
	c.token = token
}

func (c *HttpClientImpl) Execute(
	method string, url *url.URL,
	body interface{},
) (*http.Response, error) {

	reqBody, err := marshalBody(body)
	if err != nil {
		return nil, err
	}

	req, err := c.createRequest(method, url, reqBody)
	if err != nil {
		return nil, err
	}

	return c.Do(req)
}

func marshalBody(body interface{}) (io.Reader, error) {
	if body == nil {
		return nil, nil
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(jsonBody), nil
}

func (c *HttpClientImpl) createUrl(path string) string {
	return fmt.Sprintf("%s%s", c.host, path)
}

func (c *HttpClientImpl) createRequest(
	method string,
	path *url.URL,
	body io.Reader,
) (*http.Request, error) {

	path.Host = c.host
	path.Scheme = "https"

	req, err := http.NewRequest(method, path.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", c.token)
	return nil, nil
}
