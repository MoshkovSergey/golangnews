package main

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestUndefinedRouteReturns404(t *testing.T) {
	go main()

	time.Sleep(1 * time.Second)

	resp, err := http.Get("http://localhost:8080/undefined")
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNotFound {
		t.Fatalf("Expected status code 404, got %d", resp.StatusCode)
	} else {
		fmt.Println("page 404 is ok!")
	}
}
