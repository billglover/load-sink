package main

import (
	"flag"
	"log"
	"net/http"
	"sync"

	"github.com/billglover/load-sink/handlers"
)

var apiMux = http.NewServeMux()
var healthMux = http.NewServeMux()

func main() {
	log.Println("starting")

	// parse the configuration
	h := &handlers.Handler{}

	payloadSize := flag.Int("size", 1000, "the size of the payload in bytes")
	payloadVar := flag.Int("variance", 0, "requested variance in payload size in bytes")
	responseDelay := flag.Int("delay", 0, "the delay before returning API response")
	responseJitter := flag.Int("jitter", 0, "requested variance in the delay before returning an API response")
	apiAddress := flag.String("endpoint", ":8080", "requested variance in the delay before returning an API response")
	healthAddress := flag.String("health", ":8081", "requested variance in the delay before returning an API response")

	flag.Parse()

	h.PayloadSize = *payloadSize
	h.PayloadVar = *payloadVar
	h.ResponseDelay = *responseDelay
	h.ResponseJitter = *responseJitter
	APIAddress := *apiAddress
	HealthAddress := *healthAddress

	log.Printf("configuration parameter: size=%d\n", h.PayloadSize)
	log.Printf("configuration parameter: variance=%d\n", h.PayloadVar)
	log.Printf("configuration parameter: delay=%d\n", h.ResponseDelay)
	log.Printf("configuration parameter: jitter=%d\n", h.ResponseJitter)
	log.Printf("configuration parameter: endpoint=%s\n", APIAddress)
	log.Printf("configuration parameter: health=%s\n", HealthAddress)

	apiMux.HandleFunc("/", h.SendResponse)
	healthMux.HandleFunc("/", h.SendHealthResponse)

	var servers sync.WaitGroup
	servers.Add(1)
	go func() {
		defer servers.Done()
		log.Println("api starting on", APIAddress)
		err := http.ListenAndServe(APIAddress, apiMux)
		if err != nil {
			log.Fatal(err)
		}
	}()

	servers.Add(1)
	go func() {
		defer servers.Done()
		log.Println("health monitor starting on", HealthAddress)
		err := http.ListenAndServe(HealthAddress, healthMux)
		if err != nil {
			log.Fatal(err)
		}
	}()

	servers.Wait()
}
