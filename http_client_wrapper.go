package shared

import (
	"bytes"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type HttpClientWrapper interface {
	ExecuteRequest(path, method string, body []byte) (*http.Response, error)
}

type httpClientWrapper struct {
	HTTPClient  *http.Client
	baseUrl     string
	headers     map[string]string
	torProxyUrl string
}

func NewHttpClientWrapper(baseUrl, torProxyURL string, timeoutSeconds time.Duration, headers map[string]string, addJsonHeaders bool) (HttpClientWrapper, error) {
	if addJsonHeaders {
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
		HTTPClient:  client,
		baseUrl:     baseUrl,
		headers:     headers,
		torProxyUrl: torProxyURL,
	}, nil
}

// ExecuteRequest calls an endpoint, optional body and error handling. path is appended to the baseUrl, same for json
// headers and authorization. If request results in non 2xx response, will always return error with payload body in err message.
// response should have defer response.Body.Close() after the error check as it could be nil when err is != nil
func (h httpClientWrapper) ExecuteRequest(path, method string, body []byte) (*http.Response, error) {
	req, err := http.NewRequest(method, h.baseUrl+path, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	if len(h.headers) > 0 {
		for hk, hv := range h.headers {
			req.Header.Set(hk, hv)
		}
	}
	// todo retry
	// retry func only supports returning error, so we'd need to have the res http.response be pre-assigned, also for tor
	// return err to func if want it to retry
	res, err := h.HTTPClient.Do(req)
	// handle error status codes
	if err == nil && res != nil && res.StatusCode > 299 {
		defer res.Body.Close() //nolint
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, errors.Wrapf(err, "error reading failed request body")
		}
		return nil, errors.Errorf("received non success status code %d with body: %s", res.StatusCode, string(body))
	}
	return res, err
}
