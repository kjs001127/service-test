package e2e

import (
	"testing"
)

type request struct {
	method string
	path   string
	header map[string]string
	body   map[string]any
}
type assertionFunc func(t *testing.T, resMap map[string]any)
type expectedResponse struct {
	statusCode    int
	body          map[string]any
	assertionFunc assertionFunc
}
type mockServer struct {
	url              string
	expectedRequests []struct {
		req              request
		expectedResponse expectedResponse
	}
}

type testInfo struct {
	name             string
	req              request
	mockServers      []mockServer
	expectedResponse expectedResponse
	beforeTest       func() map[string]string
}
