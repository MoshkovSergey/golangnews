// Server starts successfully on port 8080
package main

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

// Server starts successfully on port 8080
func TestServerStartsSuccessfully(t *testing.T) {
	go main()

	time.Sleep(1 * time.Second)

	resp, err := http.Get("http://localhost:8080/")
	if err != nil {
		t.Fatalf("Failed to start server: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	} else {
		fmt.Println("server running, ok!")
	}
}
