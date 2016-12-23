# load-sink
 This is a simple HTTP(S) endpoint that can be used to act as a sink for load testing.

## Building the Container

[![Build Status](https://travis-ci.org/billglover/load-sink.svg?branch=master)](https://travis-ci.org/billglover/load-sink)

```
CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w'
docker build -t load-sink .
```

The latest container is available on the Docker Hub registry: [billglover/load-sink/](https://hub.docker.com/r/billglover/load-sink/)

## Feature Wishlist

 - HTTP responses (200)
 - HTTPS responses (200)
 - Configurable response time
 - Health endpoint
