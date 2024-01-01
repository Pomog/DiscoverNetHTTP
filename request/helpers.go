package request

import (
	"fmt"
	"net/http"
)

// PrintRequestParams prints the details of an HTTP request and its associated parameters.
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
