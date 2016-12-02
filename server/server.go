package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/suyashkumar/home-automation/server/mqtt"
	"github.com/suyashkumar/home-automation/server/routes"
	"net/http"
	"os"
)

func main() {
	router := httprouter.New()
	router.GET("/api/send/:deviceName/:funcName", routes.AuthMiddlewareGenerator(routes.Send))
	router.GET("/api/streams/:deviceName/:streamName", routes.AuthMiddlewareGenerator(routes.GetStreamedMessages))
	router.GET("/api/users", routes.ListUsers)
	router.POST("/api/auth", routes.Auth)
	router.GET("/api/new", routes.New)
	router.GET("/api/auth/test", routes.AuthMiddlewareGenerator(routes.Test))
	router.GET("/", routes.Hello)
	router.OPTIONS("/api/*sendPath", routes.Headers)

	mqtt.RunServer()
	err := http.ListenAndServe(":"+os.Getenv("PORT"), router)
	panic(err)
}
