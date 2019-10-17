// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"PlanetLabs/AuthService/app"
	"PlanetLabs/AuthService/restapi/operations"
)

//go:generate swagger generate server --target ..\..\AuthService --name PlanetAuth --spec ..\swagger.yml

func configureFlags(api *operations.PlanetAuthAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.PlanetAuthAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.CreateGroupHandler = operations.CreateGroupHandlerFunc(func(params operations.CreateGroupParams) middleware.Responder {
		return app.CreateGroup(params)
	})
	api.CreateUserHandler = operations.CreateUserHandlerFunc(func(params operations.CreateUserParams) middleware.Responder {
		return app.CreateUser(params)
	})
	api.DeleteGroupHandler = operations.DeleteGroupHandlerFunc(func(params operations.DeleteGroupParams) middleware.Responder {
		return app.DeleteGroup(params)
	})
	api.DeleteUserHandler = operations.DeleteUserHandlerFunc(func(params operations.DeleteUserParams) middleware.Responder {
		return app.DeleteUser(params)
	})
	api.GetGroupHandler = operations.GetGroupHandlerFunc(func(params operations.GetGroupParams) middleware.Responder {
		return app.GetGroup(params)
	})
	api.GetUserHandler = operations.GetUserHandlerFunc(func(params operations.GetUserParams) middleware.Responder {
		return app.GetUser(params)
	})
	api.UpdateGroupHandler = operations.UpdateGroupHandlerFunc(func(params operations.UpdateGroupParams) middleware.Responder {
		return app.UpdateGroup(params)
	})
	api.UpdateUserHandler = operations.UpdateUserHandlerFunc(func(params operations.UpdateUserParams) middleware.Responder {
		return app.UpdateUser(params)
	})

	err := app.InitializeDataStore()
	if err != nil {
		panic("Could not initialize server. " + err.Error())
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

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
