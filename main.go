package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Pomog/DiscoverNetHTTP/request"
)

func main() {
	// Create an instance of TestHandler
	// Start the HTTP server with your TestHandler
	// Start the HTTP server in a goroutine
	// server.StartLocalServer()

	// Wait for a few seconds to let the server start
	// time.Sleep(2 * time.Second)

	// Create an HTTP client.This concept must be assimilated
	client := http.Client{}

	// URL for the GET request (http://localhost:8080 in this case)
	urlAddress := "http://52.91.213.117/login" // http://52.91.213.117/login test server deployed on AWS

	// Create an instance of HTTPRequestParams for the GET request
	params := request.HTTPRequestParams{
		Client:     &client,
		URL:        urlAddress,
		Method:     "GET",
		TimeoutSec: 10,
		// Add more parameters as needed
	}

	// Make a GET request to server
	response, err := request.MakeRequest(params)
	if err != nil {
		log.Fatal("Error:", err)
	}
	defer response.Body.Close()

	// Process the response
	fmt.Println("response to GET request")
	fmt.Println("Status Code: ", response.Status)
	fmt.Println("Header:", response.Header)
	fmt.Println("Cookies: ", response.Cookies())
	fmt.Println("Request.UR:", response.Request.URL)
	fmt.Println("response.TLS: ", response.TLS)

	// Read and print the response body
	body := make([]byte, 512)
	_, _ = response.Body.Read(body)
	fmt.Println("Response Body:", string(body))

}
