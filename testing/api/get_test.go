package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/hello-world", nil)

	HelloWorld(rr, req)

	if http.StatusOK != rr.Code {
		t.Errorf("status code expected: %d, got: %d", http.StatusOK, rr.Code)
	}
}
