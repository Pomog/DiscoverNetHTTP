package request

import (
	"context"
	"net/http"
	"net/url"
	"time"
)

// generateGETrequest generates a sample GET request with cookies and values.
func generateGETrequest(params HTTPRequestParams) (*http.Request, error) {
	// Parse the request URL
	reqURL, err := url.Parse(params.URL)
	if err != nil {
		return nil, err
	}

	// Create a new http.Request with method, URL, and an empty header
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

func MakeGetRequest(params HTTPRequestParams) (*http.Response, error) {
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

	// Do sends an HTTP request and returns an HTTP response
	return params.Client.Do(req)
}
