.PHONY: default
.PHONY: run-local-backend
.PHONY: run-local-frontend

.PHONY: build
.PHONY: build-backend build-backend-executable build-backend-image
.PHONY: build-frontend-site

.PHONY: run-containers

.PHONY: generate-apis generate-go-server-api generate-typescript-client-api

OPENAPI_GENERATOR_VERSION:=v6.2.1

LOCAL_API_USER:="test"
LOCAL_API_TOKEN:="1234"

########################################################
# Default command, in case someone does juat 'make' w/o target
########################################################

default:
	@echo "See README.md for instructions on the common Makefile targets"
	@echo "Available targets:"
	@make -qp | awk -F':' '/^[a-zA-Z0-9][^$$#\/\t=]*:([^=]|$$)/ {split($$1,A,/ /);for(i in A)print A[i]}' | sort -u

#########################################################
# Local (emulator) commands
#########################################################

run-local-backend:
	cd device-api/cmd \
	&&	PORT=8084 \
		API_USER=${LOCAL_API_USER} \
		API_TOKEN=${LOCAL_API_TOKEN} \
		go run main.go

run-local-frontend:
	cd web-ui \
	&&  npm install \
	&&	VITE_BACKEND_API_ENDPOINT="http://localhost:8084" \
		npm run dev

#########################################################
# Image/site build commands
#########################################################

build: build-backend build-frontend-site

build-backend: build-backend-executable build-backend-image

build-backend-executable:
	cd device-api \
	&&	GOOS=linux \
		GOARCH=amd64 \
		CGO_ENABLED=1 \
		CC=musl-gcc \
		go build --ldflags '-linkmode=external -extldflags=-static' -o build/app cmd/main.go

build-backend-image:
	docker build device-api -t device-api:latest

build-frontend-site:
	cd web-ui \
	&&	VITE_BACKEND_API_ENDPOINT="/api" \
		npm run build

	rm -r deployment/web/static
	cp -r web-ui/dist deployment/web/static

#########################################################
# Deployment commands
#########################################################

run-deployment:
	cd deployment \
	&&	docker compose up

#########################################################
# API regeneration commands
#########################################################

generate-apis: generate-go-server-api generate-typescript-client-api

generate-go-server-api:

	rm -rf device-api/generated
	docker run \
		--rm \
		-v "${PWD}:/local" \
		--user $(shell id -u):$(shell id -g) \
		openapitools/openapi-generator-cli:${OPENAPI_GENERATOR_VERSION} \
		generate \
		--git-user-id=stb-org \
		--git-repo-id=stb/device-api/generated \
		-i /local/openapi/openapi.yaml \
		-g go-server \
		--additional-properties=enumClassPrefix=true,hideGenerationTimestamp=true,generateAliasAsModel=false \
		-o /local/device-api/generated

# Ensure the Golang server generated code is not in a separate module
	rm device-api/generated/go.mod

#	The Golang server generator emits an import for "github.com/gorilla/mux" even when not necessary. Remove that manually.
	sed -i 's:"github.com/gorilla/mux"::g' device-api/generated/go/api_default.go


generate-typescript-client-api:

#	rm -rf web-ui/src/generated/
	docker run \
		--rm \
		-v "${PWD}:/local" \
		--user $(shell id -u):$(shell id -g) \
		openapitools/openapi-generator-cli:${OPENAPI_GENERATOR_VERSION} \
		generate \
		-i /local/openapi/openapi.yaml \
		-g typescript-axios \
		--additional-properties=generateAliasAsModel=false \
		-o /local/web-ui/src/generated
	
#	The TypeScript client generator emits the type 'Bool' for booleans. Change that to 'boolean' to make it valid TS code
	sed -i 's/Bool/boolean/g' web-ui/src/generated/api.ts
