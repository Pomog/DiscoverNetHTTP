package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	testHandler := &TestHandler{}

	s := &http.Server{
		Addr:           ":8080",
		Handler:        testHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // bitwise left shift operation that creates an integer value 1048576 bytes or 1 megabyte
	}
	log.Fatal(s.ListenAndServe())

}


/*
	define a function that satisfies the http.Handler interface
	the http.Handler interface requires implementing a method called ServeHTTP,
	which takes an http.ResponseWriter and an http.Request as parameters
	any type implementing the Handler interface must have a ServeHTTP method with the specified signature
	ServeHTTP(ResponseWriter, *Request)
*/
type TestHandler struct{}

func (h *TestHandler) ServeHTTP (w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Test Handler")
}