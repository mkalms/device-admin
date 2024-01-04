# Device admin service example

This is a demonstration of how to create an admin service for a device, such as a set-top box. It consists of two parts:

* There is a web frontend. This is a single-page application, built using 
[Vue 3](https://vuejs.org/).

* There is a backend API. This is built using Golang. This contains all code that will run on the box itself. Here is where the logic for poking the box's local configuration files.

These components are packaged up as Docker containers, and are ready to be deployed onto the box itself.

# Development & building

* Install Visual Studio Code
* Install Docker and the Docker Compose plugin
* Install Golang 1.18+
* Install [nvm](https://github.com/nvm-sh/nvm)
* Install the appropriate Node version by doing `nvm install $(cat web-ui/.nvmrc)`
* Activate the appropriate Node version by doing `nvm use $(cat web-ui/.nvmrc)` in each terminal window

## Develop & test locally

* Run backend locally: `make run-local-backend` -- can be reached at http://localhost:8084/

* Run frontend locally: `make run-local-frontend` -- http://localhost:5173/

## Build & deploy

* Build backend container: `make build-backend`

* Build frontend site: `make build-frontend`

* Start containers: `make run-deployment` -- site is available at http://localhost:80/ \
  (run `ifconfig eth0` in WSL to find the localhost address)

## Notes

* The username/password for local use can be found in the [device-api/Makefile](device-api/Makefile)
* The username/password for deployment can be found in [deployment/compose.yaml](deployment/compose.yaml)
* After modifying [the OpenAPI specification](openapi/openapi.yaml), regenerate glue code via `make generate-apis`
* The current version uses 10MB RAM for the web container, and 5MB RAM for the backend API

## First time installation on WSL2 and Ubuntu 22.04

```
# Install Docker, Golang, Make, and musl
sudo apt-get update && sudo apt install docker.io golang-go make musl-tools

# Install nvm
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.7/install.sh | bash

# Install the Docker Compose plugin
mkdir -p ~/.docker/cli-plugins/
curl -SL https://github.com/docker/compose/releases/download/v2.23.3/docker-compose-linux-x86_64 -o ~/.docker/cli-plugins/docker-compose
chmod +x ~/.docker/cli-plugins/docker-compose

# Add your user to the Docker group (log out and log in to re-evaluate group membership) 
sudo usermod -aG docker ${USER}

# Install the appropritate Node version (after cloning this repo)
nvm install $(cat web-ui/.nvmrc)

# Activate the appropriate Node version (after cloning this repo)
nvm use $(cat web-ui/.nvmrc)
```