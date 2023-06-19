package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGreet(t *testing.T) {
	greetServer := httptest.NewServer(http.HandlerFunc(greet))
	defer greetServer.Close()

	greetClient := greetServer.Client()

	resp, err := greetClient.Get(greetServer.URL)
	if err != nil {
		t.Fatalf("Error making GET request to greet server: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code %v, got %v", http.StatusOK, resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Error reading response body: %v", err)
	}
	defer resp.Body.Close()

	if string(data) != "Hello World!" {
		t.Fatalf("Expected response body to be %v, got %v", "Hello, World!", string(data))
	}

	t.Log("TestGreet passed!")
}
