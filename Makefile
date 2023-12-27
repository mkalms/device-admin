.PHONY: default
.PHONY: run-local-backend

.PHONY: generate-apis generate-go-server-api generate-go-client-api generate-typescript-client-api

OPENAPI_GENERATOR_VERSION:=v6.2.1

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

run-local-sql-auth-proxy:
	./binaries/cloud_sql_proxy -instances "$(shell jq -r ".cloudSQLInstance" < $(ENV)/config.json)=tcp:5432" -fd_rlimit 1024 -enable_iam_login -credential_file=$(ENV)/database/google_application_credentials.json

run-local-psql:
	psql "host=127.0.0.1 sslmode=disable dbname=performance_backend user=$(shell jq -r ".psqlUser" < $(ENV)/config.json)"

run-local-backend:
	cd backend/cmd \
	&&	PORT=8084 \
		go run main.go

#########################################################
# API regeneration commands
#########################################################

generate-apis: generate-go-server-api generate-go-client-api generate-typescript-client-api

generate-go-server-api:

	rm -rf backend/generated
	docker run \
		--rm \
		-v "${PWD}:/local" \
		--user $(shell id -u):$(shell id -g) \
		openapitools/openapi-generator-cli:${OPENAPI_GENERATOR_VERSION} \
		generate \
		--git-user-id=stb-org \
		--git-repo-id=stb/backend/generated \
		-i /local/openapi-stb.yaml \
		-g go-server \
		--additional-properties=enumClassPrefix=true,hideGenerationTimestamp=true,generateAliasAsModel=false \
		-o /local/backend/generated

	rm backend/generated/go.mod

generate-typescript-client-api:

#	rm -rf frontend/src/generated/
	docker run \
		--rm \
		-v "${PWD}:/local" \
		--user $(shell id -u):$(shell id -g) \
		openapitools/openapi-generator-cli:${OPENAPI_GENERATOR_VERSION} \
		generate \
		-i /local/openapi-stb.yaml \
		-g typescript-axios \
		--additional-properties=generateAliasAsModel=false \
		-o /local/frontend/src/generated
	
#	The TypeScript client generator emits the type 'Bool' for booleans. Change that to 'boolean' to make it valid TS code
	sed -i 's/Bool/boolean/g' frontend/src/generated/api.ts
