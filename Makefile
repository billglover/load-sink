DOCKER_IMAGE_NAME = billglover/load-sink

default:
	CGO_ENABLED=0 go build -a -tags netgo -ldflags '-w'

docker:
	CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w'
	docker build -t $(DOCKER_IMAGE_NAME) .

test:
	CGO_ENABLED=0 go test -v -a -tags netgo -ldflags '-w' ./...

bench:
	vegeta attack -targets=load/targets.txt -duration=5s | tee load/results.bin | vegeta report

clean:
	go clean