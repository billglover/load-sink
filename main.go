package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/billglover/load-sink/handlers"
)

var apiMux = http.NewServeMux()

func main() {
	log.Println("starting")

	// parse the configuration
	h := &handlers.Handler{}

	payloadSize := flag.Int("size", 1000, "the size of the payload in bytes")
	payloadVar := flag.Int("variance", 0, "requested variance in payload size in bytes")
	responseDelay := flag.Int("delay", 0, "the delay before returning API response")
	responseJitter := flag.Int("jitter", 0, "requested variance in the delay before returning an API response")
	apiAddress := flag.String("endpoint", ":8080", "requested variance in the delay before returning an API response")

	flag.Parse()

	h.PayloadSize = *payloadSize
	h.PayloadVar = *payloadVar
	h.ResponseDelay = *responseDelay
	h.ResponseJitter = *responseJitter
	APIAddress := *apiAddress

	log.Printf("configuration parameter: size=%d\n", h.PayloadSize)
	log.Printf("configuration parameter: variance=%d\n", h.PayloadVar)
	log.Printf("configuration parameter: delay=%d\n", h.ResponseDelay)
	log.Printf("configuration parameter: jitter=%d\n", h.ResponseJitter)
	log.Printf("configuration parameter: endpoint=%s\n", APIAddress)

	apiMux.HandleFunc("/", h.SendResponse)

	log.Fatal(http.ListenAndServe(APIAddress, apiMux))
}
