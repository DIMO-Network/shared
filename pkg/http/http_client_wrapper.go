package http

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/time/rate"

	"github.com/avast/retry-go/v4"
	"github.com/pkg/errors"
)

var ErrNotFound = errors.New("not found")
var ErrBadRequest = errors.New("bad request")

type ClientWrapper interface {
	ExecuteRequest(path, method string, body []byte) (*http.Response, error)
	ExecuteRequestWithAuth(path, method string, body []byte, authHeader string) (*http.Response, error)
}

type clientWrapper struct {
	httpClient  *http.Client
	baseURL     string
	headers     map[string]string
	torProxyURL string
	retry       uint
	limiter     *rate.Limiter
}

type ResponseError struct {
	error
	StatusCode int
}

func BuildResponseError(statusCode int, err error) error {
	return ResponseError{StatusCode: statusCode, error: err}
}

func NewClientWrapper(baseURL, torProxyURL string, timeout time.Duration, headers map[string]string, addJSONHeaders bool, opts ...ClientWrapperOption) (ClientWrapper, error) {
	if headers == nil {
		headers = map[string]string{}
	}
	if addJSONHeaders {
		headers["Accept"] = "application/json"
		headers["Content-Type"] = "application/json"
	}
	client := &http.Client{
		Timeout: timeout,
	}
	if torProxyURL != "" {
		proxyURL, err := url.Parse(torProxyURL)
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse Tor proxy URL")
		}
		client.Transport = &http.Transport{Proxy: http.ProxyURL(proxyURL)}
	}

	params := defaultHTTPClientWrapperOptions
	for _, opt := range opts {
		opt(&params)
	}

	if params.Transport != nil {
		client.Transport = params.Transport
	}

	return &clientWrapper{
		httpClient:  client,
		baseURL:     baseURL,
		headers:     headers,
		torProxyURL: torProxyURL,
		retry:       params.Retry,
		limiter:     params.Limiter,
	}, nil
}

type ClientWrapperOptions struct {
	Retry     uint
	Transport *http.Transport
	Limiter   *rate.Limiter
}

var defaultHTTPClientWrapperOptions = ClientWrapperOptions{
	Retry: 5,
}

type ClientWrapperOption func(*ClientWrapperOptions)

func WithRetry(retry uint) ClientWrapperOption {
	return func(opts *ClientWrapperOptions) {
		opts.Retry = retry
	}
}

func WithLimiter(limiter *rate.Limiter) ClientWrapperOption {
	return func(opts *ClientWrapperOptions) {
		opts.Limiter = limiter
	}
}

func WithTransport(transport *http.Transport) ClientWrapperOption {
	return func(opts *ClientWrapperOptions) {
		opts.Transport = transport
	}
}

// GraphQLQuery calls ExecuteRequestWithAuth with the prepared gql query. authHeader is optional. The baseURL must be
// set in the clientWrapper with the full path eg. including /query path.
//
// Example:
// client := NewClientWrapper("https://identity-api.dimo.zone/query",...)
// query := `{vehicle(tokenId:123){id,owner}}`
//
//	var wrapper struct {
//			Data struct {
//				Vehicle coremodels.Vehicle `json:"vehicle"`
//			} `json:"data"`
//		}
//
// err := client.GraphQLQuery(query, &wrapper)
func (h clientWrapper) GraphQLQuery(authHeader, query string, result interface{}) error {
	// GraphQL request
	requestPayload := GraphQLRequest{Query: query}
	payloadBytes, err := json.Marshal(requestPayload)
	if err != nil {
		return err
	}

	// POST request
	res, err := h.ExecuteRequestWithAuth("", "POST", payloadBytes, authHeader)
	if err != nil {
		// not sure why we do this
		if _, ok := err.(ResponseError); !ok {
			return errors.Wrapf(err, "error calling identity api from url %s. request: %s", h.baseURL, string(payloadBytes))
		}
	}
	defer res.Body.Close() // nolint

	if res.StatusCode == 404 {
		return ErrNotFound
	}
	if res.StatusCode == 400 {
		return ErrBadRequest
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return errors.Wrapf(err, "error reading response body from url %s", h.baseURL)
	}

	if err := json.Unmarshal(bodyBytes, result); err != nil {
		return err
	}

	return nil
}

func (h clientWrapper) ExecuteRequest(path, method string, body []byte) (*http.Response, error) {
	return h.ExecuteRequestWithAuth(path, method, body, "")
}

// ExecuteRequestWithAuth calls an endpoint, optional body and error handling. path is appended to the baseURL, same for json
// headers and authorization. If request results in non 2xx response, will always return error with payload body in err message.
// response should have defer response.Body.Close() after the error check as it could be nil when err is != nil. Adds auth if passed in
func (h clientWrapper) ExecuteRequestWithAuth(path, method string, body []byte, authHeader string) (*http.Response, error) {
	req, err := http.NewRequest(method, h.baseURL+path, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	for hk, hv := range h.headers {
		req.Header.Set(hk, hv)
	}
	if authHeader != "" {
		req.Header.Set("Authorization", authHeader)
	}

	var res *http.Response

	if h.limiter != nil {
		// Wait for permission from limiter to proceed
		err = h.limiter.Wait(context.Background())
		if err != nil {
			return nil, err
		}
	}

	err = retry.Do(func() error {
		res, err = h.httpClient.Do(req)
		// handle error status codes
		if err == nil && res != nil && res.StatusCode > 299 {
			defer res.Body.Close() //nolint
			body, err := io.ReadAll(res.Body)
			if err != nil {
				return errors.Wrapf(err, "error reading failed request body")
			}
			errResp := BuildResponseError(res.StatusCode, errors.Errorf("received non success status code %d for url %s with body: %s", res.StatusCode, req.URL.String(), string(body)))
			if code := res.StatusCode; 400 <= code && code < 500 {
				// unrecoverable since probably bad payload or n
				return retry.Unrecoverable(BuildResponseError(code, errResp))
			}
			return errResp
		}
		return err
	}, retry.Attempts(h.retry), retry.Delay(500*time.Millisecond), retry.MaxDelay(9*time.Second))

	if res == nil {
		return nil, err
	}

	if retryErrors, ok := err.(retry.Error); ok {
		for _, e := range retryErrors {
			if httpError, ok := e.(ResponseError); ok {
				return res, httpError
			}
		}
	}

	return res, err
}

type GraphQLRequest struct {
	Query string `json:"query"`
}
