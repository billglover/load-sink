package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"sync"
)

var apiAddress string
var healthAddress string

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
		io.WriteString(w, "hello world")
	})

	healthMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		io.WriteString(w, "healthy world")
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
