package rest_api

import (
	"context"
	"log"
	"net/http"

	openapi "github.com/device-admin-org/device-admin/device-api/generated/go"
)

func GetDeviceConfig(ctx context.Context) (openapi.ImplResponse, error) {
	log.Printf("GetDeviceConfig called")
	getDeviceConfigResponse := openapi.GetDeviceConfigResponse{
		IpAddress: "1.2.3.4",
		NetMask:   "255.255.255.0",
		BitRate:   115200,
		Duplex:    "full",
	}
	return openapi.Response(http.StatusOK, getDeviceConfigResponse), nil
}

func SetDeviceConfig(ctx context.Context, request openapi.SetDeviceConfigRequest) (openapi.ImplResponse, error) {
	log.Printf("SetDeviceConfig called with request: %v", request)
	return openapi.Response(http.StatusOK, nil), nil
}
