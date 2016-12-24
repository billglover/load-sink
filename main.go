package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"sync"
)

// HealthResponse defines the structure of the response to a GET request on
// the health API
type HealthResponse struct {
	Status          string `json:"status"`
	Delay           int    `json:"delay"`
	Jitter          int    `json:"jitter"`
	PayloadSize     int    `json:"size"`
	PayloadVariance int    `json:"variance"`
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
		w.Header().Set("Content-Type", "text/plain")
		w.Write(RandStringBytesMaskImprSrc(*payloadSize + *payloadVar))
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

		w.Header().Set("Content-Type", "text/plain")
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
