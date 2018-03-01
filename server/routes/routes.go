package routes

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/suyashkumar/conduit/server/handlers"
)

// Build returns a new Router configured to serve all application routes
func Build() *httprouter.Router {
	r := httprouter.New()

	// RESTful API Routes
	r.GET("/", handlers.Index)
	r.GET("/hello", handlers.Hello)

	// Configure static file serving from /static
	r.ServeFiles("/static/*filepath", http.Dir("public/static"))

	return r
}
