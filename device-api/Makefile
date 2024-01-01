.PHONY: run-local-backend
.PHONY: build build-backend-executable build-backend-image

LOCAL_API_USER:="test"
LOCAL_API_TOKEN:="1234"

build: build-backend-executable build-backend-image

run-local-backend:
	cd cmd \
	&&	PORT=8084 \
		API_USER=${LOCAL_API_USER} \
		API_TOKEN=${LOCAL_API_TOKEN} \
		go run main.go

build-backend-executable:
	GOOS=linux \
	GOARCH=amd64 \
	CGO_ENABLED=1 \
	CC=musl-gcc \
	go build --ldflags '-linkmode=external -extldflags=-static' -o build/app cmd/main.go

build-backend-image:
	docker build . -t device-api:latest
