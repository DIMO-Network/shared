package http

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func Test_httpClientWrapper_ExecuteRequest_failsTooManyRetries(t *testing.T) {
	const baseURL = "http://test.com"
	retryCount := uint(5)

	hcw, _ := NewClientWrapper(baseURL, "", 1, nil, true, WithRetry(retryCount))
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	path := "/api/vehicle/v2/makes"
	httpmock.RegisterResponder(http.MethodGet, baseURL+path, httpmock.NewStringResponder(503, "error: too many requests"))
	response, err := hcw.ExecuteRequest(path, "GET", nil)
	countInfo := httpmock.GetCallCountInfo()
	c := countInfo["GET "+baseURL+path]

	assert.Equal(t, int(retryCount), c, "expected five retries")
	assert.Error(t, err, "expected error")
	assert.ErrorIs(t, err, err.(ResponseError), "expected ResponseError")
	assert.Equal(t, 503, err.(ResponseError).StatusCode, "expected 409")
	assert.Equal(t, 503, response.StatusCode)
}

func Test_httpClientWrapper_ExecuteRequest_doesNotRetryCertainStatusCodes(t *testing.T) {
	const baseURL = "http://test.com"
	hcw, _ := NewClientWrapper(baseURL, "", 1, nil, true)

	tests := []struct {
		name       string
		statusCode int
	}{
		{
			name:       "don't retry 400",
			statusCode: 400,
		},
		{
			name:       "don't retry 401",
			statusCode: 401,
		},
		{
			name:       "don't retry 404",
			statusCode: 404,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()
			path := "/api/vehicle/v2/makes"
			httpmock.RegisterResponder(http.MethodGet, baseURL+path, httpmock.NewStringResponder(test.statusCode, "error: whatever"))
			response, err := hcw.ExecuteRequest(path, "GET", nil)
			countInfo := httpmock.GetCallCountInfo()
			c := countInfo["GET "+baseURL+path]

			assert.Equal(t, 1, c, "expected single call")
			assert.Equal(t, test.statusCode, response.StatusCode)
			assert.Error(t, err, "expected error")
		})
	}
}
