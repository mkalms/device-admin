.PHONY: run-local-frontend build-frontend-site

.DEFAULT: build-frontend-site

DEPLOYMENT_DIR:=../deployment/web/static

run-local-frontend:
	npm install \
	&&	VITE_BACKEND_API_ENDPOINT="http://localhost:8084" \
		npm run dev

build-frontend-site:
	VITE_BACKEND_API_ENDPOINT="/api" \
	npm run build
	rm -r ${DEPLOYMENT_DIR}
	cp -r dist ${DEPLOYMENT_DIR}