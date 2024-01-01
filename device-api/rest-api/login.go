package rest_api

import (
	"context"
	"log"
	"net/http"

	authentication "github.com/stb-org/stb/device-api/authentication"
	openapi "github.com/stb-org/stb/device-api/generated/go"
)

func Login(ctx context.Context, loginRequest openapi.LoginRequest) (openapi.ImplResponse, error) {
	log.Printf("Login called")
	if authentication.IsUserAndAPITokenValid(loginRequest.Username, loginRequest.Password) {
		log.Printf("Username/password valid")
		return openapi.Response(http.StatusOK, nil), nil
	} else {
		log.Printf("Username/password invalid")
		return openapi.Response(http.StatusUnauthorized, &openapi.MessageResponse{Message: "Invalid username/password combination"}), nil
	}
}
