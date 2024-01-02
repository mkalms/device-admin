.PHONY: all run-local build-executable build-image

LOCAL_API_USER:="test"
LOCAL_API_TOKEN:="1234"

all: build-executable build-image

run-local:
	cd cmd \
	&&	PORT=8084 \
		API_USER=${LOCAL_API_USER} \
		API_TOKEN=${LOCAL_API_TOKEN} \
		go run main.go

build-executable:
	GOOS=linux \
	GOARCH=amd64 \
	CGO_ENABLED=1 \
	CC=musl-gcc \
	go build --ldflags '-linkmode=external -extldflags=-static' -o build/app cmd/main.go

build-image:
	docker build . -t device-api:latest
