package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloYou(t *testing.T) {
	// setup
	person := Person{
		Name: "Augusto",
	}

	var person_bytes bytes.Buffer
	err := json.NewEncoder(&person_bytes).Encode(person)
	if err != nil {
		t.Error(err)
	}

	request_recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/hello", &person_bytes)

	// check status code
	HelloYou(request_recorder, request)
	if http.StatusOK != request_recorder.Code {
		t.Errorf("status code expected: %d, got: %d", http.StatusOK, request_recorder.Code)
	}

	// check response
	response, err := io.ReadAll(request_recorder.Body)
	if err != nil {
		t.Error(err)
	}
	expected_body := fmt.Sprintf("Hello %s\n", person.Name)
	if expected_body != string(response) {
		t.Errorf("response expected: %s, got: %s", expected_body, string(response))
	}
}
