.PHONY: all run-local build clean

all: build

DEPLOYMENT_DIR:=../deployment/web-ui/static

run-local:
	npm install \
	&&	VITE_BACKEND_API_ENDPOINT="http://localhost:8084" \
		npm run dev

build: clean
	VITE_BACKEND_API_ENDPOINT="/api" \
	npm run build
	cp -r dist ${DEPLOYMENT_DIR}

clean:
	rm -rf ${DEPLOYMENT_DIR}
