package main

import (
    "fmt"
    "github.com/andelf/go-curl"
)

// writeCallback is used to capture response data
func writeCallback(data []byte, userdata interface{}) bool {
    fmt.Print(string(data))
    return true
}

func main() {
    // Initialize curl
    easy := curl.EasyInit()
    if easy == nil {
        fmt.Println("Failed to initialize curl")
        return
    }
    defer easy.Cleanup()

    // Configure GET request
    easy.Setopt(curl.OPT_URL, "https://jsonplaceholder.typicode.com/posts/1")
    easy.Setopt(curl.OPT_HTTPGET, true) // Explicitly set GET method
    easy.Setopt(curl.OPT_WRITEFUNCTION, writeCallback)

    // Perform the request
    if err := easy.Perform(); err != nil {
        fmt.Printf("GET request failed: %v\n", err)
        return
    }
    fmt.Println("\nGET request completed")
}