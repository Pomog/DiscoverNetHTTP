package server

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func StartLocalServer() {
	testHandler := &TestHandler{}

	s := &http.Server{
		Addr:           ":8080",
		Handler:        testHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {

		log.Fatal(s.ListenAndServe())
	}()
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

	// Set a custom status code (e.g., 202 Accepted)
	w.WriteHeader(http.StatusAccepted)

	// Send a response to the client
	fmt.Fprintln(w, "Test Handler - GET")
}

func handlePOST(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Test Handler - POST")
}
