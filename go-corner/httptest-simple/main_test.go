package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// $ go run httptest-simple.go
// then run:
func TestHelloHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello", nil)
	w := httptest.NewRecorder()
	hello(w, req)
	resp := w.Result()
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("read response body failed, err:%v", err)
	}

	if string(data) != "Hello World" {
		t.Errorf("expect response data:Hello World, but got:%s", string(data))
	}
	t.Logf("response data:%s", string(data))
}
