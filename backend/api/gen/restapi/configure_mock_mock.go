// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/mock-mock/mockmock-meter/backend/api/gen/models"
	"github.com/mock-mock/mockmock-meter/backend/api/gen/restapi/operations"
	"github.com/mock-mock/mockmock-meter/backend/api/gen/restapi/operations/health"
	"github.com/mock-mock/mockmock-meter/backend/api/gen/restapi/operations/mock"
	"github.com/mock-mock/mockmock-meter/backend/api/gen/restapi/operations/web"
)

//go:generate swagger generate server --target ../../gen --name MockMock --spec ../../swagger.yaml

func configureFlags(api *operations.MockMockAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.MockMockAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	/*
		Health
	*/
	api.HealthHelloHandler = health.HelloHandlerFunc(func(params health.HelloParams) middleware.Responder {
		return health.NewHelloOK().WithPayload(&models.Health{Message: "OK"})
	})
	api.HealthHealthcheckHandler = health.HealthcheckHandlerFunc(func(params health.HealthcheckParams) middleware.Responder {
		return health.NewHealthcheckOK().WithPayload(&models.Health{Message: "OK"})
	})

	/*
		Mock
	*/
	api.MockPostMockHandler = mock.PostMockHandlerFunc(func(params mock.PostMockParams) middleware.Responder {
		return middleware.NotImplemented("operation mock.PostMock has not yet been implemented")
	})

	// web
	api.WebWebresourceHandler = web.WebresourceHandlerFunc(func(params web.WebresourceParams) middleware.Responder {
		return web.NewWebresourceOK().WithPayload(&models.Web{Message: "OK"})
	})

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
