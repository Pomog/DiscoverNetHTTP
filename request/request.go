package request

import (
	"context"
	"crypto/tls"
	"net/http"
	"net/url"
	"time"
)

// HTTPRequestParams represents the parameters for an HTTP request.
type HTTPRequestParams struct {
	Client         *http.Client
	URL            string            // URL of the request
	Method         string            // HTTP method (GET, POST, etc.)
	Headers        map[string]string // HTTP headers
	Cookies        []*http.Cookie    // Cookies
	Values         url.Values        // Form values (for POST requests, etc.)
	Body           string            // Request body for non-form data
	TimeoutSec     int               // Timeout in seconds
	UserAgent      string            // User-Agent header for the request
	QueryParams    url.Values        // Additional query parameters
	Username       string            // Username for authentication
	Password       string            // Password for authentication
	RequestContext context.Context   // Context for the request
	ProxyURL       *url.URL          // URL of the proxy to be used for the request
	TLSConfig      *tls.Config       // TLS configuration for the request
	// Add more parameters as needed
}

// generateGETrequest generates a sample GET request with cookies and values.
func generateGETrequest(params HTTPRequestParams) (*http.Request, error) {
	reqURL, err := url.Parse(params.URL)
	if err != nil {
		return nil, err
	}

	req := &http.Request{
		Method: params.Method,
		URL:    reqURL,
		Header: make(http.Header),
	}

	// Set headers
	for key, value := range params.Headers {
		req.Header.Set(key, value)
	}

	// Add cookies to the request
	for _, cookie := range params.Cookies {
		req.AddCookie(cookie)
	}

	// Set values for form data (for GET requests)
	req.Form = params.Values

	// Set request context
	if params.RequestContext != nil {
		req = req.WithContext(params.RequestContext)
	}

	return req, nil
}

func MakeRequest(params HTTPRequestParams) (*http.Response, error) {
	req, err := generateGETrequest(params)
	if err != nil {
		return nil, err
	}

	// Set a timeout if specified
	if params.TimeoutSec > 0 {
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(params.TimeoutSec)*time.Second)
		defer cancel()
		req = req.WithContext(ctx)
	}

	return params.Client.Do(req)
}
