# load-sink
 This is a simple HTTP(S) endpoint that can be used to act as a sink for load testing.

## Building the Container

[![Build Status](https://travis-ci.org/billglover/load-sink.svg?branch=master)](https://travis-ci.org/billglover/load-sink)

```
CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w'
docker build -t load-sink .
```

The latest container is available on the Docker Hub registry: [billglover/load-sink/](https://hub.docker.com/r/billglover/load-sink/)

## Running the Container

```
docker run -p 8080:8080 -p 8081:8081 -d load-sink
```

Query the main API: 

```
curl -i -XGET http://localhost:8080
```
```
HTTP/1.1 200 OK
Content-Type: text/plain
Date: Fri, 23 Dec 2016 21:48:04 GMT
Content-Length: 11

hello world
```

Query the health API: 

```
curl -i -XGET http://localhost:8081
```
```
HTTP/1.1 200 OK
Content-Type: text/plain
Date: Fri, 23 Dec 2016 21:48:30 GMT
Content-Length: 13

healthy world
```


## Feature Wishlist

 - HTTP responses (200)
 - HTTPS responses (200)
 - Configurable response time
 - Health endpoint
