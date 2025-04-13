package main

import (
    "github.com/andelf/go-curl"
    "strings"
    "testing"
)

func TestPutRequest(t *testing.T) {
    // Capture response
    var response strings.Builder
    writeCallback := func(data []byte, userdata interface{}) bool {
        response.Write(data)
        return true
    }

    // Initialize curl
    easy := curl.EasyInit()
    if easy == nil {
        t.Fatal("Failed to initialize curl")
    }
    defer easy.Cleanup()

    // JSON payload
    payload := `{"id": 1, "title": "Updated Post", "body": "Updated content", "userId": 1}`

    // Configure PUT request
    easy.Setopt(curl.OPT_URL, "https://jsonplaceholder.typicode.com/posts/1")
    easy.Setopt(curl.OPT_CUSTOMREQUEST, "PUT")
    easy.Setopt(curl.OPT_POSTFIELDS, payload)
    easy.Setopt(curl.OPT_HTTPHEADER, []string{"Content-Type: application/json"})
    easy.Setopt(curl.OPT_WRITEFUNCTION, writeCallback)

    // Perform the request
    if err := easy.Perform(); err != nil {
        t.Fatalf("PUT request failed: %v", err)
    }

    // Verify response contains updated data
    resp := response.String()
    if !strings.Contains(resp, `"title": "Updated Post"`) {
        t.Errorf("Expected response to contain updated title, got: %s", resp)
    }
}