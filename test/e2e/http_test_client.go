package e2e

import (
	"context"
	"encoding/json"
	"github.com/go-resty/resty/v2"
)

type HttpTestClient struct {
	ctx     context.Context
	client  *resty.Client
	method  string
	body    json.RawMessage
	baseURL string
}

func NewHttpTestClient(baseURL string) *HttpTestClient {
	return &HttpTestClient{client: resty.New(), baseURL: baseURL}
}

func (m *HttpTestClient) Get(path string) *HttpTestClient {
	m.method = resty.MethodGet
	m.client.SetBaseURL(m.baseURL + path)
	return m
}

func (m *HttpTestClient) Post(path string) *HttpTestClient {
	m.method = resty.MethodPost
	m.client.SetBaseURL(m.baseURL + path)
	return m
}

func (m *HttpTestClient) Put(path string) *HttpTestClient {
	m.method = resty.MethodPut
	m.client.SetBaseURL(m.baseURL + path)
	return m
}

func (m *HttpTestClient) Delete(path string) *HttpTestClient {
	m.method = resty.MethodDelete
	m.client.SetBaseURL(m.baseURL + path)
	return m
}

func (m *HttpTestClient) SetBody(body json.RawMessage) *HttpTestClient {
	m.body = body
	return m
}

func (m *HttpTestClient) SetHeader(key, value string) *HttpTestClient {
	m.client.SetHeader(key, value)
	return m
}

func (m *HttpTestClient) Do() (*resty.Response, error) {
	r := m.client.R()
	r.SetContext(m.ctx)
	r.SetBody(m.body)
	resp, err := r.Execute(m.method, m.client.BaseURL)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
