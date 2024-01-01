package request

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func PostByteSlice() {
	// URL for the GET request (http://localhost:8080 in this case)
	urlAddress := "http://localhost:8080/registration" // http://52.91.213.117/login test server deployed on AWS

	// Create an instance of HTTPRequestParams for the POST request
	// Prepare form data
	bodyPost := &bytes.Buffer{}
	writer := multipart.NewWriter(bodyPost)

	_ = writer.WriteField("firstName", "John")
	_ = writer.WriteField("lastName", "Doe")
	_ = writer.WriteField("nickName", "JohnD")
	_ = writer.WriteField("emailRegistr", "john.doe@example.com")
	_ = writer.WriteField("passwordReg", "securePassword123")

	// Add file to the request
	// Open the file to be included in the request
	file, err := os.Open("johvi.png")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create the form field for the file
	part, err := writer.CreateFormFile("avatar", filepath.Base(file.Name()))
	if err != nil {
		fmt.Println("Error creating form file:", err)
		return
	}

	// Copy the file content to the form field
	_, err = io.Copy(part, file)
	if err != nil {
		fmt.Println("Error copying file to form file:", err)
		return
	}

	// Write a custom header for the file part
	part.Write([]byte("Content-Type: image/png\r\n\r\n"))

	// Close the writer to finalize the form data
	_ = writer.Close()

	// Create the request
	request, err := http.NewRequest("POST", urlAddress, bodyPost)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set the Content-Type header to the writer's boundary
	request.Header.Set("Content-Type", writer.FormDataContentType())

	// Send the request
	clientPost := &http.Client{}
	responsePost, err := clientPost.Do(request)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer responsePost.Body.Close()

	// Process the response
	fmt.Println("response to POST request")
	fmt.Println("Status Code: ", responsePost.Status)
	fmt.Println("Header:", responsePost.Header)
	fmt.Println("Cookies: ", responsePost.Cookies())
	fmt.Println("Request.URL:", responsePost.Request.URL)

	// Read and print the response body
	bodyResp := make([]byte, 5000)
	_, _ = responsePost.Body.Read(bodyResp)
	fmt.Println("Response Body:", string(bodyResp))
}
