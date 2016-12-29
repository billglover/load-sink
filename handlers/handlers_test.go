package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

var apiMux = http.NewServeMux()
var healthMux = http.NewServeMux()

func init() {
	h := &Handler{}

	healthMux.HandleFunc("/", h.SendHealthResponse)
	apiMux.HandleFunc("/", h.SendResponse)
}

func TestGetHealth(t *testing.T) {
	t.Log("given the need to test the health endpoint")

	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal("\tshould be able to create a request", ballotX, err)
	}
	t.Log("\tshould be able to create a request", checkMark)

	rw := httptest.NewRecorder()
	healthMux.ServeHTTP(rw, req)

	if rw.Code != http.StatusOK {
		t.Fatal("\tshould receive \"200\" response", ballotX, rw.Code)
	}
	t.Log("\tshould receive \"200\" response", checkMark)

	hr := HealthResponse{}

	if err := json.NewDecoder(rw.Body).Decode(&hr); err != nil {
		t.Fatal("\tshould decode the response", ballotX, err)
	}
}

func TestGetAPI(t *testing.T) {
	t.Log("given the need to test the API endpoint")

	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal("\tshould be able to create a request", ballotX, err)
	}
	t.Log("\tshould be able to create a request", checkMark)

	rw := httptest.NewRecorder()
	apiMux.ServeHTTP(rw, req)

	if rw.Code != http.StatusOK {
		t.Fatal("\tshould receive \"200\" response", ballotX, rw.Code)
	}
	t.Log("\tshould receive \"200\" response", checkMark)

	hr := APIResponse{}

	if err := json.NewDecoder(rw.Body).Decode(&hr); err != nil {
		t.Fatal("\tshould decode the response", ballotX, err)
	}
}
