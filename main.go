package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"sync"
)

// HealthResponse defines the structure of the response to a GET request on
// the health API
type HealthResponse struct {
	Status      string  `json:"status"`
	Delay       float32 `json:"delay"`
	Jitter      float32 `json:"jitter"`
	PayloadSize int32   `json:"payloadsize"`
}

var apiAddress string
var healthAddress string
var responseDelay float32
var responseJitter float32
var payloadSize int32

func init() {
	if apiAddress = os.Getenv("API_ADDRESS"); apiAddress == "" {
		apiAddress = ":8080"
	}
	log.Println("configuration parameter: apiAddress", apiAddress)

	if healthAddress = os.Getenv("HEALTH_ADDRESS"); healthAddress == "" {
		healthAddress = ":8081"
	}
	log.Println("configuration parameter: healthAddress", healthAddress)

	log.Println("API initialised")
}

func main() {
	log.Println("starting")

	var apiMux = http.NewServeMux()
	var healthMux = http.NewServeMux()

	apiMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write(RandStringBytesMaskImprSrc(1000))
	})

	healthMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		res := &HealthResponse{
			Status:      "ok",
			Delay:       responseDelay,
			Jitter:      responseJitter,
			PayloadSize: payloadSize,
		}

		resBytes, _ := json.Marshal(res)

		w.Header().Set("Content-Type", "text/plain")
		w.Write(resBytes)
	})

	var servers sync.WaitGroup
	servers.Add(1)
	go func() {
		defer servers.Done()
		log.Println("api starting on", apiAddress)
		err := http.ListenAndServe(apiAddress, apiMux)
		if err != nil {
			log.Fatal(err)
		}
	}()

	servers.Add(1)
	go func() {
		defer servers.Done()
		log.Println("health monitor starting on", healthAddress)
		err := http.ListenAndServe(healthAddress, healthMux)
		if err != nil {
			log.Fatal(err)
		}
	}()

	servers.Wait()
}
