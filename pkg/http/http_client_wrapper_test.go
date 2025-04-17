package http

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

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

func Test_clientWrapper_GraphQLQuery(t *testing.T) {
	const baseURL = "http://graphql.test.com/query"
	retryCount := uint(5)

	hcw, _ := NewClientWrapper(baseURL, "", 1, nil, true, WithRetry(retryCount))
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	query := `{
  vehicle(tokenId: 172541) {
    id
    owner
    definition {
      id
    }
    name
    privileges {
      nodes {
        user
      }
    }
  }
}`
	identityResponse := `{
  "data": {
    "vehicle": {
      "id": "V_kc4AAqH9",
      "owner": "0x593CAD66A44B63EEe8B54Ff4Be9d1E2bA7C1dA40",
      "definition": {
        "id": "dodge_1500_2025"
      },
      "name": "add manual liar",
      "privileges": {
        "nodes": [],
        "edges": []
      }
    }
  }
}`
	var wrapper struct {
		Data struct {
			Vehicle Vehicle `json:"vehicle"`
		} `json:"data"`
	}

	httpmock.RegisterResponder(http.MethodPost, baseURL, httpmock.NewStringResponder(200, identityResponse))
	err := hcw.GraphQLQuery("", query, &wrapper)
	require.NoError(t, err)

	countInfo := httpmock.GetCallCountInfo()
	c := countInfo["POST "+baseURL]
	assert.Equal(t, 1, c, "expected 1 request count")
	assert.Equal(t, "V_kc4AAqH9", wrapper.Data.Vehicle.ID)
	assert.Equal(t, "0x593CAD66A44B63EEe8B54Ff4Be9d1E2bA7C1dA40", wrapper.Data.Vehicle.Owner)
	assert.Equal(t, "add manual liar", wrapper.Data.Vehicle.Name)
	assert.Equal(t, "dodge_1500_2025", wrapper.Data.Vehicle.Definition.ID)
	assert.Equal(t, 0, len(wrapper.Data.Vehicle.Privileges.Nodes))
	assert.Equal(t, 0, len(wrapper.Data.Vehicle.Privileges.Edges))
}

type Vehicle struct {
	ID         string `json:"id"`
	Owner      string `json:"owner"`
	Definition struct {
		ID string `json:"id"`
	} `json:"definition"`
	Name       string `json:"name"`
	Privileges struct {
		Nodes []interface{} `json:"nodes"`
		Edges []interface{} `json:"edges"`
	} `json:"privileges"`
}
