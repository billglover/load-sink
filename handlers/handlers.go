package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/billglover/load-sink/har"
)

// APIResponse defines the structure of the response to a GET request on
// the main API endpoint
type APIResponse struct {
	Status      int    `json:"status"`
	Delay       int    `json:"delay"`
	Payload     string `json:"payload"`
	PayloadSize int    `json:"payload_size"`
	Path        string `json:"path"`
}

// Handler holds the configuration for the API handler
type Handler struct {
	PayloadSize    int
	PayloadVar     int
	ResponseDelay  int
	ResponseJitter int
}

// SendResponse handles a request to the main API endpoint. It uses configuration
// stored in the Handler struct to determine how to respond to the request
func (h *Handler) SendResponse(w http.ResponseWriter, r *http.Request) {
	d := int(randomOffset(int32(h.ResponseDelay), int32(h.ResponseJitter)))

	// calculate the desired payload size
	s := int(randomOffset(int32(h.PayloadSize), int32(h.PayloadVar)))

	res := &APIResponse{
		Status:      http.StatusOK,
		Delay:       d,
		Payload:     string(randStringBytesMaskImprSrc(s)),
		PayloadSize: s,
		Path:        r.URL.String(),
	}
	log.Printf("request: %s", r.RequestURI)

	resBytes, _ := json.Marshal(res)

	// hold the response based on the specified delay
	time.Sleep(time.Millisecond * time.Duration(d))

	w.Header().Set("Content-Type", "application/json")
	w.Write(resBytes)
}

func randomOffset(m int32, r int32) (t int32) {

	// handle the case where our range is zero
	if r == 0 {
		t = m
		return
	}

	// otherwise apply random offset to the mid-point
	// based on the range provided
	t = rand.Int31n((2*r)+1) - r + m
	return
}

// Echo handles requests to `POST /echo` and returns a response with identical
// body and Content-Type as the request.
func (h *Handler) Echo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))

	if r.Body != nil {
		bodyBytes, _ := ioutil.ReadAll(r.Body)
		w.Write(bodyBytes)
	}
}

// Request handles requests to `ALL /request/{path}` and returns a response
// encoded as an HAR format object.
func (h *Handler) Request(w http.ResponseWriter, r *http.Request) {
	l, err := har.FromHTTPRequest(r)

	// unable to parse request
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(l)

	// unable to send response
	if err != nil {
		panic(err)
	}
}
