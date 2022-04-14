package shared

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

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
}

func NewHTTPClientWrapper(baseURL, torProxyURL string, timeoutSeconds time.Duration, headers map[string]string, addJSONHeaders bool) (HTTPClientWrapper, error) {
	if headers == nil {
		headers = map[string]string{}
	}
	if addJSONHeaders {
		headers["Accept"] = "application/json"
		headers["Content-Type"] = "application/json"
	}
	client := &http.Client{
		Timeout: timeoutSeconds * time.Second,
	}
	if torProxyURL != "" {
		proxyURL, err := url.Parse(torProxyURL)
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse Tor proxy URL")
		}
		client.Transport = &http.Transport{Proxy: http.ProxyURL(proxyURL)}
	}
	return &httpClientWrapper{
		httpClient:  client,
		baseURL:     baseURL,
		headers:     headers,
		torProxyURL: torProxyURL,
	}, nil
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
	res := new(http.Response)

	err = retry.Do(func() error {
		res, err = h.httpClient.Do(req)
		// handle error status codes
		if err == nil && res != nil && res.StatusCode > 299 {
			defer res.Body.Close() //nolint
			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				return errors.Wrapf(err, "error reading failed request body")
			}
			errResp := errors.Errorf("received non success status code %d with body: %s", res.StatusCode, string(body))
			if res.StatusCode == 400 || res.StatusCode == 401 || res.StatusCode == 404 {
				// unrecoverable since probably bad payload or n
				return retry.Unrecoverable(errResp)
			}
			return errResp
		}
		return nil
	}, retry.Attempts(5), retry.Delay(500*time.Millisecond), retry.MaxDelay(9*time.Second))

	return res, err
}
