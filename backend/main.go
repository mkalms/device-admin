package backend

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	authentication "github.com/stb-org/stb/backend/authentication"
	openapi "github.com/stb-org/stb/backend/generated/go"
)

type ApiService struct {
}

var corsHandler *cors.Cors
var router *mux.Router

func init() {

	apiService := &ApiService{}
	DefaultApiController := openapi.NewDefaultApiController(apiService)

	router = openapi.NewRouter(DefaultApiController)

	// Set up CORS handler

	corsHandler = cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowedHeaders:     []string{"*"},
		MaxAge:             3600,
		OptionsPassthrough: false,
		Debug:              false,
	})

	// Set up authentication

	usernamePasswordValidator := authentication.CreateUsernamePasswordValidator()
	authenticationMiddleware := authentication.CreateAuthenticationMiddleware([]authentication.Validator{
		usernamePasswordValidator,
	})

	router.Use(authenticationMiddleware.Handler)
}

func BackendAPI(w http.ResponseWriter, r *http.Request) {
	log.Printf("Backend API request URL path: %v", r.URL.Path)

	// Invoke CORS handler, and forward request to router in case it is a non-preflight request
	//
	// While the router has a mechanism for chaining to middlewares, the router will only do so for routes
	//  that it knows about. We don't list any of the OPTIONS-method routes in the OpenAPI spec,
	//  thus the router doesn't know about them, and thus the router will automatically return HTTP 405 Method Not Allowed
	//  for those routes, without giving the CORS middleware any chance to handle them.
	// We work around this by invoking the CORS middleware explicitly, before the router.
	corsHandler.ServeHTTP(w, r, router.ServeHTTP)
}
