// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"fmt"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/rs/cors"

	"github.com/nicolas2bert/ba-server/apiv1/auth"
	"github.com/nicolas2bert/ba-server/apiv1/restapi/intern/users"
	"github.com/nicolas2bert/ba-server/apiv1/restapi/ui/photos"
	"github.com/nicolas2bert/ba-server/gen/restapi/operations"
	"github.com/nicolas2bert/ba-server/gen/restapi/operations/intern"
	"github.com/nicolas2bert/ba-server/gen/restapi/operations/ui"
)

func jwtDecode(tokenString string, secret string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		signingMethod, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok || signingMethod != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte("secret"), nil
	})

	if err != nil {
		return nil, errors.Unauthenticated("issue: parsing token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.Unauthenticated("issue: claiming token")
	}
	return &claims, nil
}

//go:generate swagger generate server --target ../gen --name ba --spec ../swagger.yaml

func configureFlags(api *operations.BaAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.BaAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.InternGetUsersIDHandler = intern.GetUsersIDHandlerFunc(users.GetUsersIDHandler)
	api.InternSaveUserHandler = intern.SaveUserHandlerFunc(users.SaveUserHandler)
	// api.InternalPostUsersHandler internal.PostUsersHandler

	api.UIGetPhotosHandler = ui.GetPhotosHandlerFunc(photos.GetPhotosHandler)

	api.InternAPIAuth = func(tokenString string) (*auth.PrincipalBA, error) {
		claims, err := jwtDecode(tokenString, "secret")
		if err != nil {
			return nil, errors.Unauthenticated("issue: decoding token")
		}

		role, ok := (*claims)["role"]
		if !ok {
			return nil, errors.Unauthenticated("issue: role not found")
		}

		return &auth.PrincipalBA{
			Role: role.(string),
		}, nil
	}

	api.UIAPIAuth = func(tokenString string) (*auth.PrincipalBA, error) {
		// TODO: USE ENV VAR FOR SECRET
		claims, err := jwtDecode(tokenString, "secret")
		if err != nil {
			return nil, errors.Unauthenticated("issue: decoding token")
		}

		userID, ok := (*claims)["userId"]
		if !ok {
			return nil, errors.Unauthenticated("issue: userId not found")
		}

		role, ok := (*claims)["role"]
		if !ok {
			return nil, errors.Unauthenticated("issue: role not found")
		}

		return &auth.PrincipalBA{
			UserID: userID.(string),
			Role:   role.(string),
		}, nil
	}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

func WithCors(handler http.Handler) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://127.0.0.1:3003",
			"http://localhost:3003",
		},
		AllowedMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			// "OPTIONS",
		},
		AllowedHeaders: []string{
			"x-ui-ba-token",
			// "Content-Type",
		},
		ExposedHeaders: []string{
			"x-ui-ba-token",
		},
	})
	return c.Handler(handler)
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return WithCors(handler)
}
