package shared

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/time/rate"

	"github.com/avast/retry-go/v4"
	"github.com/pkg/errors"
)

type HTTPClientWrapper interface {
	ExecuteRequest(path, method string, body []byte) (*http.Response, error)
}

type httpClientWrapper struct {
	httpClient  *http.Client
	baseURL     string
	headers     map[string]string
	torProxyURL string
	retry       uint
	limiter     *rate.Limiter
}

type HTTPResponseError struct {
	error
	StatusCode int
}

func BuildResponseError(statusCode int, err error) error {
	return HTTPResponseError{StatusCode: statusCode, error: err}
}

func NewHTTPClientWrapper(baseURL, torProxyURL string, timeout time.Duration, headers map[string]string, addJSONHeaders bool, opts ...HTTPClientWrapperOption) (HTTPClientWrapper, error) {
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

	return &httpClientWrapper{
		httpClient:  client,
		baseURL:     baseURL,
		headers:     headers,
		torProxyURL: torProxyURL,
		retry:       params.Retry,
		limiter:     params.Limiter,
	}, nil
}

type HTTPClientWrapperOptions struct {
	Retry     uint
	Transport *http.Transport
	Limiter   *rate.Limiter
}

var defaultHTTPClientWrapperOptions = HTTPClientWrapperOptions{
	Retry: 5,
}

type HTTPClientWrapperOption func(*HTTPClientWrapperOptions)

func WithRetry(retry uint) HTTPClientWrapperOption {
	return func(opts *HTTPClientWrapperOptions) {
		opts.Retry = retry
	}
}

func WithLimiter(limiter *rate.Limiter) HTTPClientWrapperOption {
	return func(opts *HTTPClientWrapperOptions) {
		opts.Limiter = limiter
	}
}

func WithTransport(transport *http.Transport) HTTPClientWrapperOption {
	return func(opts *HTTPClientWrapperOptions) {
		opts.Transport = transport
	}
}

// ExecuteRequest calls an endpoint, optional body and error handling. path is appended to the baseURL, same for json
// headers and authorization. If request results in non 2xx response, will always return error with payload body in err message.
// response should have defer response.Body.Close() after the error check as it could be nil when err is != nil
func (h httpClientWrapper) ExecuteRequest(path, method string, body []byte) (*http.Response, error) {
	req, err := http.NewRequest(method, h.baseURL+path, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	for hk, hv := range h.headers {
		req.Header.Set(hk, hv)
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
			if httpError, ok := e.(HTTPResponseError); ok {
				return res, httpError
			}
		}
	}

	return res, err
}
