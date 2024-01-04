package rest_api

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	openapi "github.com/device-admin-org/device-admin/device-api/generated/go"
)

var deviceConfigFilePath string = "/mnt/shared_data/device_config.json"

func GetDeviceConfig(ctx context.Context) (openapi.ImplResponse, error) {
	log.Printf("GetDeviceConfig called\n")
	data, err := ioutil.ReadFile(deviceConfigFilePath)
	if err != nil {
		log.Printf("Could not read json file, using defaults: %s\n", err)
		response := openapi.GetDeviceConfigResponse{
			IpAddress: "1.2.3.4",
			NetMask:   "255.255.255.0",
			BitRate:   115200,
			Duplex:    "full",
		}
		return openapi.Response(http.StatusOK, response), nil
	}
	var response openapi.GetDeviceConfigResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		log.Printf("Could not unmarshal json: %s\n", err)
		return openapi.Response(http.StatusOK, nil), nil
	}
	log.Printf("Unmarshalled json file: %v\n", response)
	return openapi.Response(http.StatusOK, response), nil
}

func SetDeviceConfig(ctx context.Context, request openapi.SetDeviceConfigRequest) (openapi.ImplResponse, error) {
	log.Printf("SetDeviceConfig called with request: %v\n", request)
	jsonData, err := json.Marshal(request)
	if err != nil {
		log.Printf("Could not marshal json: %s\n", err)
		return openapi.Response(http.StatusOK, nil), nil
	}
	log.Printf("Marshalled json: %s\n", &jsonData)
	err = ioutil.WriteFile(deviceConfigFilePath, jsonData, 0)
	if err != nil {
		log.Printf("Could not write marshalled json file: %s\n", err)
		return openapi.Response(http.StatusOK, nil), nil
	}
	return openapi.Response(http.StatusOK, nil), nil
}
