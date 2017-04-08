package app

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/suyashkumar/conduit/server/handlers"
	"github.com/suyashkumar/conduit/server/middleware"
	"github.com/suyashkumar/conduit/server/mqtt"
	"github.com/suyashkumar/conduit/server/secrets"
	"gopkg.in/mgo.v2"
	"net/http"
	"os"
)

type AppConfig struct {
	IsDev bool
}

type ConduitApp interface {
	Run()
}

type ConduitAppImpl struct {
	AppConfig
	Router  *httprouter.Router
	context *handlers.HandlerContext
}

func (c *ConduitAppImpl) Run() {
	c.attachRoutes()
	mqtt.RunServer()
	c.startWebServer()
}

func (c *ConduitAppImpl) attachRoutes() {
	c.Router.GET("/api/send/:deviceName/:funcName", c.WrapAuthHandler(handlers.Send))
	c.Router.GET("/api/streams/:deviceName/:streamName", c.WrapAuthHandler(handlers.GetStreamedMessages))
	c.Router.POST("/api/auth", c.WrapHandler(handlers.Auth))
	c.Router.POST("/api/register", c.WrapHandler(handlers.New))
	c.Router.GET("/api/me", c.WrapAuthHandler(handlers.GetUser))
	c.Router.GET("/", handlers.Hello)
	c.Router.OPTIONS("/api/*sendPath", handlers.Headers)
	c.Router.ServeFiles("/static/*filepath", http.Dir("public/static"))
}

func (c *ConduitAppImpl) WrapAuthHandler(next handlers.AuthHandler) httprouter.Handle {
	return middleware.ConduitAuthMiddleware(next, c.context)
}

func (c *ConduitAppImpl) WrapHandler(next handlers.ConduitHandler) httprouter.Handle {
	return middleware.ConduitMiddleware(next, c.context)
}

func (c *ConduitAppImpl) startWebServer() {
	fmt.Printf("Web server to listen on port :%s", os.Getenv("PORT"))
	if c.AppConfig.IsDev {
		err := http.ListenAndServe(":"+os.Getenv("PORT"), c.Router)
		panic(err)
	} else {
		go http.ListenAndServeTLS(":443", os.Getenv("CERT"), os.Getenv("PRIV_KEY"), c.Router)
		err := http.ListenAndServe(":"+os.Getenv("PORT"), http.HandlerFunc(handlers.RedirectToHttps))
		panic(err)
	}

}

func NewConduitApp() *ConduitAppImpl {
	// Init DB Session
	session, err := mgo.Dial(secrets.DB_DIAL_URL)
	if err != nil {
		panic(err)
	}

	return &ConduitAppImpl{
		Router: httprouter.New(),
		AppConfig: AppConfig{
			IsDev: os.Getenv("DEV") == "TRUE",
		},
		context: &handlers.HandlerContext{
			DbSession: session,
		},
	}
}
