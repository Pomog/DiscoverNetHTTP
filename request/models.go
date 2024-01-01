package request

import (
	"context"
	"crypto/tls"
	"io"
	"net/http"
	"net/url"
)

// HTTPRequestParams represents the parameters for an HTTP request.
type HTTPRequestParams struct {
	Client         *http.Client         // Client is the HTTP client used to make the request.
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
