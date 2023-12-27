package authentication

import (
	"log"
	"net/http"
)

type UsernamePasswordValidator struct {
}

func CreateUsernamePasswordValidator() *UsernamePasswordValidator {
	return &UsernamePasswordValidator{}
}

func (*UsernamePasswordValidator) Validate(r *http.Request) (string, int, string) {

	// Fetch username + password from Basic Authentication header of WWW request

	username, password, basicAuthPresent := r.BasicAuth()

	if !basicAuthPresent {

		log.Print("Basic auth header (with email/token) not provided")
		return "", 0, ""
	}

	log.Printf("Basic Auth header present, username: %v, password: %v", username, password)

	// Validate that the API user + API token matches what is configured for this backend

	success := isUserAndAPITokenValid(username, password)

	if !success {
		log.Printf("API user/API token pair %v, %v is not valid", username, password)
		return "", http.StatusUnauthorized, "Unauthorized; please provide valid API user / API token pair"
	} else {

		// API user + API token are valid

		log.Printf("API user/API Token pair %v, %v is valid; authentication successful", username, password)
		return username, 0, ""
	}
}

func isUserAndAPITokenValid(username string, password string) bool {

	return username == getApiUser() && password == getApiToken()
}
