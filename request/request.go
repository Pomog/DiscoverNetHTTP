package request

import (
	"bytes"
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"net/url"
	"path/filepath"
	"time"
)

// HTTPRequestParams represents the parameters for an HTTP request.
type HTTPRequestParams struct {
	Client         *http.Client
	URL            string               // URL of the request
	Method         string               // HTTP method (GET, POST, etc.)
	Headers        map[string]string    // HTTP headers
	Cookies        []*http.Cookie       // Cookies
	Values         url.Values           // Form values (for POST requests, etc.)
	Body           string               // Request body for non-form data
	TimeoutSec     int                  // Timeout in seconds
	UserAgent      string               // User-Agent header for the request
	QueryParams    url.Values           // Additional query parameters
	Username       string               // Username for authentication
	Password       string               // Password for authentication
	RequestContext context.Context      // Context for the request
	ProxyURL       *url.URL             // URL of the proxy to be used for the request
	TLSConfig      *tls.Config          // TLS configuration for the request
	Files          map[string]io.Reader // Files to be included in the request
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

	return params.Client.Do(req)
}

func GeneratePOSTRequestWithFile(params HTTPRequestParams) (*http.Request, error) {
	var req *http.Request
	var err error

	// Check if there are files to include in the request
	if len(params.Files) > 0 {
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		// Add files to the request
		for filename, content := range params.Files {
			part, err := writer.CreatePart(textproto.MIMEHeader{
				"Content-Disposition": []string{fmt.Sprintf(`form-data; name="avatar"; filename="%s"`, filepath.Base(filename))},
				"Content-Type":        []string{"image/png"},
			})
			if err != nil {
				return nil, err
			}

			if _, err := io.Copy(part, content); err != nil {
				return nil, err
			}
		}

		// Add other form values from params.Values
		for key, values := range params.Values {
			for _, value := range values {
				_ = writer.WriteField(key, value)
			}
		}

		err = writer.Close()
		if err != nil {
			return nil, err
		}

		// Create a new request with a reader for the multipart body
		req, err = http.NewRequest(params.Method, params.URL, body)
		if err != nil {
			return nil, err
		}

		// Set the Content-Type header for the multipart form data
		req.Header.Set("Content-Type", writer.FormDataContentType())
	} else {
		return nil, errors.New("no files provided")
	}

	// Set request headers
	for key, value := range params.Headers {
		req.Header.Set(key, value)
	}

	// Set the context for the request
	if params.RequestContext != nil {
		req = req.WithContext(params.RequestContext)
	} else {
		req = req.WithContext(context.TODO())
	}

	return req, nil
}

func PrintRequestParams(req *http.Request, params HTTPRequestParams) {
	fmt.Printf("****************** Print the request details ******************\n")
	fmt.Printf("Request Method: %s\n", req.Method)
	fmt.Printf("Request URL: %s\n", req.URL)
	fmt.Printf("Request Headers:\n")
	for key, values := range req.Header {
		for _, value := range values {
			fmt.Printf("  %s: %s\n", key, value)
		}
	}
	fmt.Printf("Request Cookies:\n")
	for _, cookie := range req.Cookies() {
		fmt.Printf("  %s: %s\n", cookie.Name, cookie.Value)
	}
	fmt.Printf("Request Values From HTTPRequestParams:\n")
	for key, values := range params.Values {
		for _, value := range values {
			fmt.Printf("  %s: %s\n", key, value)
		}
	}
	fmt.Printf("Request Timeout (seconds): %d\n", params.TimeoutSec)
	fmt.Printf("Request UserAgent: %s\n", params.UserAgent)
	fmt.Printf("Request Query Params:\n")
	for key, values := range params.QueryParams {
		for _, value := range values {
			fmt.Printf("  %s: %s\n", key, value)
		}
	}
	fmt.Printf("Request Username: %s\n", params.Username)
	fmt.Printf("Request Password: %s\n", params.Password)
	fmt.Printf("Request TLS Config: %v\n", params.TLSConfig)
	fmt.Printf("Request Files:\n")
	for key, _ := range params.Files {
		fmt.Printf("	%s: \n", key)

	}
	fmt.Printf("***************************************************************\n")
}

func MakePostRequest(params HTTPRequestParams) (*http.Response, error) {
	req, err := GeneratePOSTRequestWithFile(params)
	if err != nil {
		return nil, err
	}

	PrintRequestParams(req, params)

	// Set a timeout if specified
	if params.TimeoutSec > 0 {
		ctx, cancel := context.WithTimeout(req.Context(), time.Duration(params.TimeoutSec)*time.Second)
		defer cancel()

		req = req.WithContext(ctx)
	}

	return params.Client.Do(req)
}
