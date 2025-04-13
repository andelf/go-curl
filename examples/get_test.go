package main

import (
    "github.com/andelf/go-curl"
    "strings"
    "testing"
)

func TestGetRequest(t *testing.T) {
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

    // Configure GET request
    easy.Setopt(curl.OPT_URL, "https://jsonplaceholder.typicode.com/posts/1")
    easy.Setopt(curl.OPT_HTTPGET, true)
    easy.Setopt(curl.OPT_WRITEFUNCTION, writeCallback)

    // Perform the request
    if err := easy.Perform(); err != nil {
        t.Fatalf("GET request failed: %v", err)
    }

    // Verify response contains expected JSON fields
    resp := response.String()
    if !strings.Contains(resp, `"id": 1`) {
        t.Errorf("Expected response to contain id: 1, got: %s", resp)
    }
}