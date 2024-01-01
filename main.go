package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/Pomog/DiscoverNetHTTP/request"
)

func main() {

	// URL for the request (http://localhost:8080 in this case)
	urlAddress := "http://localhost:8080/registration" // http://52.91.213.117/login test server deployed on AWS

	// Create an HTTP client.This concept must be assimilated
	client := http.Client{}
	
	// Create a POST request for user registration.
	// This request registers a new user in the FFForum web application (https://github.com/Pomog/ForumFFF),
	// which can be verified
	// by checking the creation of a new row in the SQL Database, specifically in the "users" table.

	// Include a file
	// The type of file can be defined in the request by setting
	// "Content-Type" of the part, which is a section of the request body
	filePath := "static/audit.rtf"
	fileContents, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	formValues := url.Values{}
	// Add other form values if needed.
	// These values are part of the registration form for the FFForum web application.
	formValues.Add("firstName", "John")
	formValues.Add("lastName", "Doe")
	formValues.Add("nickName", "JohnD")
	formValues.Add("emailRegistr", "john.doe@example.com")
	formValues.Add("passwordReg", "securePassword123")

	// Create an instance of HTTPRequestParams for the POST request
	paramsPOST := request.HTTPRequestParams{
		Client:     &client,
		URL:        urlAddress,
		Values:     formValues,
		Method:     "POST",
		TimeoutSec: 10,
		Files: map[string]io.Reader{
			filePath: bytes.NewReader(fileContents),
		},
	}

	// Make the POST request
	responsePost, err := request.MakePostRequest(paramsPOST)
	if err != nil {
		panic(err)
	}
	defer responsePost.Body.Close()

	// Process the response
	fmt.Println("response to POST request")
	fmt.Println("Status Code: ", responsePost.Status)
	fmt.Println("Header:", responsePost.Header)
	fmt.Println("Cookies: ", responsePost.Cookies())
	fmt.Println("Request.URL:", responsePost.Request.URL)

	// Read and print the response body
	bodyResp := make([]byte, 15000)
	_, _ = responsePost.Body.Read(bodyResp)
	fmt.Println("Response Body:", string(bodyResp))
}
