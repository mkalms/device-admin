.PHONY: all
.PHONY: generate-apis
.PHONY: run-local-backend run-local-frontend
.PHONY: build-backend build-frontend
.PHONY: run-deployment
.PHONY: clean

all: clean generate-apis build-backend build-frontend run-deployment

generate-apis:
	$(MAKE) -C openapi

run-local-backend:
	$(MAKE) -C device-api run-local

run-local-frontend:
	$(MAKE) -C web-ui run-local

build-backend:
	$(MAKE) -C device-api

build-frontend:
	$(MAKE) -C web-ui

run-deployment:
	$(MAKE) -C deployment

clean:
	$(MAKE) -C openapi clean
	$(MAKE) -C web-ui clean
