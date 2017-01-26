package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

var apiMux = http.NewServeMux()
var r = &mux.Router{}

func init() {
	h := &Handler{}
	apiMux.HandleFunc("/", h.SendResponse)

	r = mux.NewRouter()
	r.HandleFunc("/echo", h.Echo).Methods(http.MethodPost)
}

func TestGetAPI(t *testing.T) {
	t.Log("given the need to test the API endpoint")

	req, err := http.NewRequest(http.MethodGet, "/api", nil)
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

	ar := APIResponse{}

	if err := json.NewDecoder(rw.Body).Decode(&ar); err != nil {
		t.Fatal("\tshould decode the response", ballotX, err)
	}
	t.Log("\tshould decode the response", checkMark)

	if ar.Path != "/api" {
		t.Fatal("\tshould contain path of /api in the body", ballotX)
	}
	t.Log("\tshould contain path of /api in the body", checkMark)
}

func TestEcho(t *testing.T) {
	t.Log("given the need to test the /echo endpoint")

	reqBody := "hello there"
	reqContentType := "text/test"
	req, err := http.NewRequest(http.MethodPost, "/echo", bytes.NewBuffer([]byte(reqBody)))
	if err != nil {
		t.Fatal("\tshould be able to create a request", ballotX, err)
	}
	t.Log("\tshould be able to create a request", checkMark)

	req.Header.Add("Content-Type", reqContentType)

	rw := httptest.NewRecorder()
	r.ServeHTTP(rw, req)

	if rw.Code != http.StatusOK {
		t.Fatal("\tshould receive \"200\" response", ballotX, rw.Code)
	}
	t.Log("\tshould receive \"200\" response", checkMark)

	if rw.Body.String() != reqBody {
		t.Fatal("\tshould match the request body", ballotX, rw.Body.String())
	}
	t.Log("\tshould match the request body", checkMark)

	if rw.Header().Get("Content-Type") != reqContentType {
		t.Fatal("\tshould match the request Content-Type", ballotX, rw.Header().Get("Content-Type"))
	}
	t.Log("\tshould match the request Content-Type", checkMark)

}
