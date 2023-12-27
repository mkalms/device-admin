/*
 * STB config API
 *
 * This is the API that is used to configure an STB
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"net/http"
	"errors"
)

// DefaultApiService is a service that implements the logic for the DefaultApiServicer
// This service should implement the business logic for every endpoint for the DefaultApi API.
// Include any external packages or services that will be required by this service.
type DefaultApiService struct {
}

// NewDefaultApiService creates a default api service
func NewDefaultApiService() DefaultApiServicer {
	return &DefaultApiService{}
}

// GetDeviceConfig - Fetch device config
func (s *DefaultApiService) GetDeviceConfig(ctx context.Context) (ImplResponse, error) {
	// TODO - update GetDeviceConfig with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, GetDeviceConfigResponse{}) or use other options such as http.Ok ...
	//return Response(200, GetDeviceConfigResponse{}), nil

	//TODO: Uncomment the next line to return response Response(401, MessageResponse{}) or use other options such as http.Ok ...
	//return Response(401, MessageResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetDeviceConfig method not implemented")
}

// SetDeviceConfig - Set device config
func (s *DefaultApiService) SetDeviceConfig(ctx context.Context, setDeviceConfigRequest SetDeviceConfigRequest) (ImplResponse, error) {
	// TODO - update SetDeviceConfig with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	//return Response(200, nil),nil

	//TODO: Uncomment the next line to return response Response(400, MessageResponse{}) or use other options such as http.Ok ...
	//return Response(400, MessageResponse{}), nil

	//TODO: Uncomment the next line to return response Response(401, MessageResponse{}) or use other options such as http.Ok ...
	//return Response(401, MessageResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("SetDeviceConfig method not implemented")
}
