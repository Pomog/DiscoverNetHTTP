package request

import (
	"bytes"
	"context"
	"mime/multipart"
	"net/http"
	"time"
)

/*
GeneratePOSTRequestWithFile generates an HTTP POST request with files and form values.
Parameters:
- params: An HTTPRequestParams struct containing various parameters for the request.
Returns:
- *http.Request: The generated HTTP request.
- error: An error if any occurred during the request generation.
This function is designed to create an HTTP POST request with file attachments and other form values.
It uses the multipart/form-data encoding for sending files. The function takes a set of parameters
such as files, form values, request method, and others to build the request.
*/
func GeneratePOSTRequestForAudit(params HTTPRequestParams) (*http.Request, error) {
	var req *http.Request
	var err error

	// Check if there are files to include in the request
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

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

	// Set request headers
	for key, value := range params.Headers {
		req.Header.Set(key, value)
	}

	// Set the context for the request
	// The context in Go is a way to carry deadlines, cancellations, and other request-scoped values across API boundaries and between processes.
	// params.RequestContext is an optional field in the HTTPRequestParams
	if params.RequestContext != nil {
		// If a specific request context is provided in the parameters, use it.
		req = req.WithContext(params.RequestContext)
	} else {
		// If no specific context is provided, create a new empty context using context.TODO().
		req = req.WithContext(context.TODO())
	}

	return req, nil
}

// MakePostRequest makes a POST request using the provided HTTPRequestParams and returns the response.
func MakePostRequestForAudit(params HTTPRequestParams) (*http.Response, error) {
	// Generate the POST request
	req, err := GeneratePOSTRequestForAudit(params)
	if err != nil {
		return nil, err
	}

	// Print the details of the generated request
	PrintRequestParams(req, params)

	// Set a timeout if specified
	if params.TimeoutSec > 0 {
		// Create a context with timeout and update the request context
		ctx, cancel := context.WithTimeout(req.Context(), time.Duration(params.TimeoutSec)*time.Second)
		defer cancel()

		req = req.WithContext(ctx)
	}

	// Do sends an HTTP request and returns an HTTP response
	return params.Client.Do(req)
}
