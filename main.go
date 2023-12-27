package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

func main() {
	// Create an instance of TestHandler
	testHandler := &TestHandler{}

	// Start the HTTP server with your TestHandler
	s := &http.Server{
		Addr:           ":8080",
		Handler:        testHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		// Start the HTTP server in a goroutine
		log.Fatal(s.ListenAndServe())
	}()

	// Wait for a few seconds to let the server start
	time.Sleep(2 * time.Second)

	// Create an HTTP client.This concept must be assimilated
	client := http.Client{}

	// Make a GET request to server
	response, err := client.Get("http://localhost:8080")
	if err != nil {
		log.Fatal("Error:", err)
	}
	defer response.Body.Close()

	// Process the response
	fmt.Println("response to GET request")
	fmt.Println("Status Code:", response.Status)
	fmt.Println("Header:", response.Header)
	fmt.Println("Cookies:", response.Cookies())

	// Read and print the response body
	body := make([]byte, 512)
	_, _ = response.Body.Read(body)
	fmt.Println("Response Body:", string(body))

	// Make a POST request to your server
	postResponse, err := client.PostForm("http://localhost:8080", url.Values{"key": {"value"}})
	if err != nil {
		log.Fatal("POST Error:", err)
	}
	defer postResponse.Body.Close()

	// Process the response
	fmt.Println("response to POST request")
	fmt.Println("POST Status Code:", postResponse.Status)
	fmt.Println("POST Header:", postResponse.Header)
	fmt.Println("Cookies:", postResponse.Cookies())

	// Read and print the POST response body
	postBody := make([]byte, 512)
	_, _ = postResponse.Body.Read(postBody)
	fmt.Println("POST Response Body:", string(postBody))
}

/*
define a function that satisfies the http.Handler interface
the http.Handler interface requires implementing a method called ServeHTTP,
which takes an http.ResponseWriter and an http.Request as parameters
any type implementing the Handler interface must have a ServeHTTP method with the specified signature
ServeHTTP(ResponseWriter, *Request)
*/
type TestHandler struct{} // TestHandler is a custom type that implements the http.Handler interface.

// TestHandler implementation of a single method: ServeHTTP.
func (h *TestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGET(w, r)
	case http.MethodPost:
		handlePOST(w, r)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func handleGET(w http.ResponseWriter, r *http.Request) {
	// Create a new cookie
	cookie := &http.Cookie{
		Name:  "TestCookie",
		Value: "Result",
	}

	// Add the cookie to the response
	http.SetCookie(w, cookie)

	// Send a response to the client
	fmt.Fprintln(w, "Test Handler - GET")
}


func handlePOST(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Test Handler - POST")
}

// // generateGETrequest generates a sample GET request.
// func generateGETrequest() *http.Request {
// 	reqURL := &url.URL{
// 		Scheme: "http",
// 		Host:   "example.com",
// 		Path:   "/",
// 	}

// 	getRequest := &http.Request{
// 		Method: "GET",
// 		URL:    reqURL,
// 	}

// 	return getRequest
// }
