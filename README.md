# STB Config service example

This is a demonstration of how to create a configuration service for a device, such a a set-top box. It consists of two parts:

* There is a web frontend. This is a single-page application, built using 
[Vue 3](https://vuejs.org/).

* There is a backend API. This is built using Golang. This contains all code that will run on the box itself. Here is where the logic for poking the box's local configuration files.

These components are packaged up as Docker containers, and are ready to be deployed onto the box itself.

# Development & building

* Install Visual Studio Code
* Install Golang 1.18+
* Install [nvm](https://github.com/nvm-sh/nvm)

* Activate the appropriate Node version by doing `nvm use $(cat frontend/.nvmrc)` in each terminal window

## Develop & test locally

* Run backend locally: `make run-local backend` -- can be reached at http://localhost:8084/

* Run frontend locally: `make run-local-frontend` -- http://localhost:5173/

## Build & deploy

* Build backend container: `make build-backend`

* Build frontend site: `make build-frontend-site`

* Start containers: `make run-deployment` -- site is available at http://localhost:80/

## Notes

* The username/password for local use can be found in the [Makefile](Makefile)
* The username/password for deployment can be found in [compose.yaml](deployment/compose.yaml)
* After modifying [the OpenAPI specification](openapi-stb.yaml), regenerate glue code via `make generate-apis`
