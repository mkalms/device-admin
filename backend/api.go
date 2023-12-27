package backend

import (
	"context"

	openapi "github.com/stb-org/stb/backend/generated/go"
	rest_api "github.com/stb-org/stb/backend/rest-api"
)

// store API

func (s *ApiService) GetDeviceConfig(ctx context.Context) (openapi.ImplResponse, error) {
	response, err := rest_api.GetDeviceConfig(ctx)
	return response, err
}

func (s *ApiService) SetDeviceConfig(ctx context.Context, request openapi.SetDeviceConfigRequest) (openapi.ImplResponse, error) {

	response, err := rest_api.SetDeviceConfig(ctx, request)
	return response, err
}

func (s *ApiService) Reset(ctx context.Context) (openapi.ImplResponse, error) {

	response, err := rest_api.Reset(ctx)
	return response, err
}
