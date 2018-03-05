package routes

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/suyashkumar/conduit/server/device"
	"github.com/suyashkumar/conduit/server/handlers"
)

type deviceAPIHandler func(w http.ResponseWriter, r *http.Request, p httprouter.Params, d device.Handler)

// injectDeviceHandler is middleware that injects the device.Handler into the RESTful API route handler functions
func injectDeviceHandler(h deviceAPIHandler, d device.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		h(w, r, p, d)
	}
}

// Build returns a new Router configured to serve all application routes
func Build(d device.Handler) *httprouter.Router {
	r := httprouter.New()

	// RESTful API Routes
	r.GET("/", handlers.Index)
	r.GET("/hello", injectDeviceHandler(handlers.Hello, d))

	// Configure static file serving from /static
	r.ServeFiles("/static/*filepath", http.Dir("public/static"))

	// Configure device handler socket routing
	r.Handler("GET", "/socket.io/", d.GetHTTPHandler())

	return r
}
