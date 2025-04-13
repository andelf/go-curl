package main

import (
    "fmt"
    "github.com/andelf/go-curl"
    "strings"
)

// writeCallback captures response data
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

    // JSON payload for the PUT request
    payload := `{"id": 1, "title": "Updated Post", "body": "This is an updated post", "userId": 1}`
    reader := strings.NewReader(payload)

    // Configure PUT request
    easy.Setopt(curl.OPT_URL, "https://jsonplaceholder.typicode.com/posts/1")
    easy.Setopt(curl.OPT_CUSTOMREQUEST, "PUT") // Set method to PUT
    easy.Setopt(curl.OPT_POSTFIELDS, payload) // Set request body
    easy.Setopt(curl.OPT_HTTPHEADER, []string{"Content-Type: application/json"})
    easy.Setopt(curl.OPT_WRITEFUNCTION, writeCallback)

    // Perform the request
    if err := easy.Perform(); err != nil {
        fmt.Printf("PUT request failed: %v\n", err)
        return
    }
    fmt.Println("\nPUT request completed")
}