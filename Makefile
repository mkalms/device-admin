.PHONY: default
.PHONY: run-local-backend
.PHONY: run-local-frontend

.PHONY: build
.PHONY: build-backend
.PHONY: build-frontend-site

.PHONY: run-containers

.PHONY: generate-apis

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
	$(MAKE) -C device-api run-local-backend

run-local-frontend:
	cd web-ui \
	&&  npm install \
	&&	VITE_BACKEND_API_ENDPOINT="http://localhost:8084" \
		npm run dev

#########################################################
# Image/site build commands
#########################################################

build: build-backend build-frontend-site

build-backend:
	$(MAKE) -C device-api

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

generate-apis:
	$(MAKE) -C openapi