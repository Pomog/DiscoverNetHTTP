package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

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
	urlAddress := "http://localhost:8080/login" // http://52.91.213.117/login test server deployed on AWS

	// Create an instance of HTTPRequestParams for the GET request
	params := request.HTTPRequestParams{
		Client:     &client,
		URL:        urlAddress,
		Method:     "GET",
		TimeoutSec: 10,
		// Add more parameters as needed
	}

	// Make a GET request to server
	response, err := request.MakeGetRequest(params)
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

	// Read and print the response body
	body := make([]byte, 512)
	_, _ = response.Body.Read(body)
	fmt.Println("Response Body:", string(body))

	// Create an instance of HTTPRequestParams for the GET request
	paramsPOST := request.HTTPRequestParams{
		Client:     &client,
		URL:        urlAddress,
		Method:     "POST",
		TimeoutSec: 10,
	}
	// Add form data to the request
	paramsPOST.Values = url.Values{}
	paramsPOST.Values.Set("emailLogIn", "beavis@mtv.com")
	paramsPOST.Values.Set("passwordLogIn", "123456")

	// Make a Post request to server
	response, err = request.MakePostRequest(paramsPOST)
	if err != nil {
		log.Fatal("Error:", err)
	}
	defer response.Body.Close()

	// Process the response
	fmt.Println("response to POST request")
	fmt.Println("Status Code: ", response.Status)
	fmt.Println("Header:", response.Header)
	fmt.Println("Cookies: ", response.Cookies())
	fmt.Println("Request.UR: ", response.Request.URL)

	// Read and print the response body
	POSTbody := make([]byte, 5000)
	_, _ = response.Body.Read(POSTbody)
	fmt.Println("Response Body:\n", string(POSTbody))
}

// // Create an instance of HTTPRequestParams for the GET request
// paramsPOST := request.HTTPRequestParams{
// 	Client:     &client,
// 	URL:        urlAddress,
// 	Method:     "POST",
// 	TimeoutSec: 10,
// }
// // Add form data to the request
// paramsPOST.Values = url.Values{}
// paramsPOST.Values.Set("firstName", "John")
// paramsPOST.Values.Set("lastName", "Doe")
// paramsPOST.Values.Set("nickName", "JohnD")
// paramsPOST.Values.Set("emailRegistr", "john.doe@example.com")
// paramsPOST.Values.Set("passwordReg", "securePassword123")

// // Open the file that you want to include in the request
// file, err := os.Open("johvi.png")
// if err != nil {
// 	fmt.Println("Error opening file:", err)
// 	return
// }
// defer file.Close()

// // Create a new buffer to store the file contents
// fileBuffer := &bytes.Buffer{}
// if _, err := io.Copy(fileBuffer, file); err != nil {
// 	fmt.Println("Error reading file:", err)
// 	return
// }

// // Add the file to the request
// paramsPOST.Files = map[string]io.Reader{
// 	"avatar": fileBuffer,
// }

// // Make a Post request to server
// response, err = request.MakePostRequest(paramsPOST)
// if err != nil {
// 	log.Fatal("Error:", err)
// }
// defer response.Body.Close()

// // Process the response
// fmt.Println("response to POST request")
// fmt.Println("Status Code: ", response.Status)
// fmt.Println("Header:", response.Header)
// fmt.Println("Cookies: ", response.Cookies())
// fmt.Println("Request.UR: ", response.Request.URL)
// location, _ := response.Location()
// fmt.Println("location: ", location)

// // Read and print the response body
// POSTbody := make([]byte, 5000)
// _, _ = response.Body.Read(POSTbody)
// fmt.Println("Response Body:\n", string(POSTbody))
