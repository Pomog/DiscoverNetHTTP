package request

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"path/filepath"
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
func GeneratePOSTRequestWithFile(params HTTPRequestParams) (*http.Request, error) {
	var req *http.Request
	var err error

	// Check if there are files to include in the request
	if len(params.Files) > 0 {
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		// Add files to the request
		for filename, content := range params.Files {
			// writer.CreatePart:
			// This function is a method of the multipart.Writer type.
			// It creates a new part in the multipart message.
			// A part is a section of the request body that contains a file, form data, or other data.
			// textproto.MIMEHeader:
			// It's a type representing MIME headers.
			// MIME (Multipurpose Internet Mail Extensions) headers are key-value pairs associated with parts of a MIME message.
			// In this case, it is used to define the headers of the multipart part
			part, err := writer.CreatePart(textproto.MIMEHeader{
				// This header field indicates the part's role.
				"Content-Disposition": []string{fmt.Sprintf(`form-data; name="avatar"; filename="%s"`, filepath.Base(filename))},
				// This header field specifies the media type of the part's data.
				"Content-Type": []string{"image/png"},
			})
			if err != nil {
				return nil, err
			}

			// Copy the content of the file (content) to the part of the multipart message (part).
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
func MakePostRequest(params HTTPRequestParams) (*http.Response, error) {
	// Generate the POST request
	req, err := GeneratePOSTRequestWithFile(params)
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
