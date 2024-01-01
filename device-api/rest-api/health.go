package rest_api

import (
	"context"
	"log"
	"net/http"

	openapi "github.com/stb-org/stb/device-api/generated/go"
)

func Health(ctx context.Context) (openapi.ImplResponse, error) {
	log.Printf("Health check called")
	return openapi.Response(http.StatusOK, nil), nil
}
