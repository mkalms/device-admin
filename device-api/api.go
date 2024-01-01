package deviceapi

import (
	"context"

	openapi "github.com/device-admin-org/device-admin/device-api/generated/go"
	rest_api "github.com/device-admin-org/device-admin/device-api/rest-api"
)

func (s *ApiService) Login(ctx context.Context, loginRequest openapi.LoginRequest) (openapi.ImplResponse, error) {
	response, err := rest_api.Login(ctx, loginRequest)
	return response, err
}

func (s *ApiService) Health(ctx context.Context) (openapi.ImplResponse, error) {
	response, err := rest_api.Health(ctx)
	return response, err
}

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
