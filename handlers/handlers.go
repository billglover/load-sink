package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
)

// HealthResponse defines the structure of the response to a GET request on
// the health API endpoint
type HealthResponse struct {
	Status          string `json:"status"`
	Delay           int    `json:"delay"`
	Jitter          int    `json:"jitter"`
	PayloadSize     int    `json:"size"`
	PayloadVariance int    `json:"variance"`
}

// APIResponse defines the structure of the response to a GET request on
// the main API endpoint
type APIResponse struct {
	Status      int    `json:"status"`
	Delay       int    `json:"delay"`
	Payload     string `json:"payload"`
	PayloadSize int    `json:"payload_size"`
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
	}

	resBytes, _ := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")
	w.Write(resBytes)
}

// SendHealthResponse handles a request to the health API endpoint. It returns
// details of the configuration currently stored in the Handler struct
func (h *Handler) SendHealthResponse(w http.ResponseWriter, r *http.Request) {

	res := &HealthResponse{
		Status:          "ok",
		Delay:           h.ResponseDelay,
		Jitter:          h.ResponseJitter,
		PayloadSize:     h.PayloadSize,
		PayloadVariance: h.PayloadVar,
	}

	resBytes, _ := json.Marshal(res)

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
	t = rand.Int31n(2*r) - r + m
	return
}
