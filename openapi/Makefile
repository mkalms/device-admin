.PHONY: all generate-server generate-client
.PHONY: clean clean-server clean-client

OPENAPI_GENERATOR_VERSION:=v6.2.1
BASE_DIR:=$(abspath ${CURDIR}/..)
SERVER_API_DIR:=device-api/generated
CLIENT_API_DIR:=web-ui/src/generated

all: generate-server generate-client

generate-server: clean-server
	docker run \
		--rm \
		-v "${BASE_DIR}:/local" \
		--user $(shell id -u):$(shell id -g) \
		openapitools/openapi-generator-cli:${OPENAPI_GENERATOR_VERSION} \
		generate \
		--git-user-id=device-admin-org \
		--git-repo-id=device-admin/${SERVER_API_DIR} \
		-i /local/openapi/openapi.yaml \
		-g go-server \
		--additional-properties=enumClassPrefix=true,hideGenerationTimestamp=true,generateAliasAsModel=false \
		-o /local/${SERVER_API_DIR}
# Ensure the Golang server generated code is not in a separate module
	rm ${BASE_DIR}/${SERVER_API_DIR}/go.mod
# The Golang server generator emits an import for "github.com/gorilla/mux" even
# when not necessary. Remove that manually.
	sed -i 's:"github.com/gorilla/mux"::g' ${BASE_DIR}/${SERVER_API_DIR}/go/api_default.go

generate-client: clean-client
	docker run \
		--rm \
		-v "${BASE_DIR}:/local" \
		--user $(shell id -u):$(shell id -g) \
		openapitools/openapi-generator-cli:${OPENAPI_GENERATOR_VERSION} \
		generate \
		-i /local/openapi/openapi.yaml \
		-g typescript-axios \
		--additional-properties=generateAliasAsModel=false \
		-o /local/${CLIENT_API_DIR}
# The TypeScript client generator emits the type 'Bool' for booleans. Change
# that to 'boolean' to make it valid TypeScript code.
	sed -i 's/Bool/boolean/g' ${BASE_DIR}/${CLIENT_API_DIR}/api.ts

clean: clean-server clean-client

clean-server:
	rm -rf ${BASE_DIR}/${SERVER_API_DIR}

clean-client:
	rm -rf ${BASE_DIR}/${CLIENT_API_DIR}
