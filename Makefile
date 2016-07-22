
build: build-go build-docker

build-go:
	docker run --rm \
		-e GOBIN=/go/bin/ -e CGO_ENABLED=0 -e GOPATH=/go \
		-v $$(pwd):/go/src/ovhapi -w /go/src/ovhapi \
			golang:1.6.2-alpine go build

build-docker:
	docker build -t krkr/ovhapi .