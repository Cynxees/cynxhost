package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func SendHttpRequest(method string, url string, data interface{}, headers map[string]string) (string, error) {

	fmt.Println("SendPostRequest: ", url)
	// Marshal the data into JSON
	requestData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	// Create a POST request
	req, err := http.NewRequest(method, url, bytes.NewBuffer(requestData))
	if err != nil {
		return "", err
	}

	// Add headers to the request if any
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Make the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Check the response status code for success
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("received non-OK HTTP status: %s", resp.Status)
	}

	// Return the response body as a string
	return string(body), nil
}
