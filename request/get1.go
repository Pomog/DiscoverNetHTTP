package request

import (
	"fmt"
	"log"
	"net/http"
)

func Get1() {
	// Create an HTTP client.This concept must be assimilated
	client := http.Client{}

	// URL for the GET request (http://localhost:8080 in this case)
	urlAddress := "http://localhost:8080/registration" // http://52.91.213.117/login test server deployed on AWS

	// Create an instance of HTTPRequestParams for the GET request
	params := HTTPRequestParams{
		Client:     &client,
		URL:        urlAddress,
		Method:     "GET",
		TimeoutSec: 10,
		// Add more parameters as needed
	}

	// Make a GET request to server
	response, err := MakeGetRequest(params)
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
}
