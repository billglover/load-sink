package main

import (
	"encoding/json"
	"flag"
	"log"
	"math/rand"
	"net/http"
	"sync"
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
	Payload     []byte `json:"string"`
	PayloadSize int    `json:"payload_size"`
}

func main() {
	log.Println("starting")

	payloadSize := flag.Int("size", 1000, "the size of the payload in bytes")
	payloadVar := flag.Int("variance", 0, "requested variance in payload size in bytes")
	responseDelay := flag.Int("delay", 0, "the delay before returning API response")
	responseJitter := flag.Int("jitter", 0, "requested variance in the delay before returning an API response")
	apiAddress := flag.String("endpoint", ":8080", "requested variance in the delay before returning an API response")
	healthAddress := flag.String("health", ":8081", "requested variance in the delay before returning an API response")

	flag.Parse()

	log.Printf("configuration parameter: size=%d\n", *payloadSize)
	log.Printf("configuration parameter: variance=%d\n", *payloadVar)
	log.Printf("configuration parameter: delay=%d\n", *responseDelay)
	log.Printf("configuration parameter: jitter=%d\n", *responseJitter)
	log.Printf("configuration parameter: endpoint=%s\n", *apiAddress)
	log.Printf("configuration parameter: health=%s\n", *apiAddress)

	var apiMux = http.NewServeMux()
	var healthMux = http.NewServeMux()

	apiMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// calculate the desired response time
		d := int(randomOffset(int32(*responseDelay), int32(*responseJitter)))

		// calculate the desired payload size
		s := int(randomOffset(int32(*payloadSize), int32(*payloadVar)))

		res := &APIResponse{
			Status:      http.StatusOK,
			Delay:       d,
			Payload:     RandStringBytesMaskImprSrc(s),
			PayloadSize: s,
		}

		resBytes, _ := json.Marshal(res)

		w.Header().Set("Content-Type", "application/json")
		w.Write(resBytes)
	})

	healthMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		res := &HealthResponse{
			Status:          "ok",
			Delay:           *responseDelay,
			Jitter:          *responseJitter,
			PayloadSize:     *payloadSize,
			PayloadVariance: *payloadVar,
		}

		resBytes, _ := json.Marshal(res)

		w.Header().Set("Content-Type", "application/json")
		w.Write(resBytes)
	})

	var servers sync.WaitGroup
	servers.Add(1)
	go func() {
		defer servers.Done()
		log.Println("api starting on", *apiAddress)
		err := http.ListenAndServe(*apiAddress, apiMux)
		if err != nil {
			log.Fatal(err)
		}
	}()

	servers.Add(1)
	go func() {
		defer servers.Done()
		log.Println("health monitor starting on", *healthAddress)
		err := http.ListenAndServe(*healthAddress, healthMux)
		if err != nil {
			log.Fatal(err)
		}
	}()

	servers.Wait()
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
