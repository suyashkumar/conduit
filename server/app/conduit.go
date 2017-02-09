package app

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/suyashkumar/conduit/server/middleware"
	"github.com/suyashkumar/conduit/server/mqtt"
	"github.com/suyashkumar/conduit/server/routes"
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
	context *routes.HandlerContext
}

func (c *ConduitAppImpl) Run() {
	c.attachRoutes()
	mqtt.RunServer()
	c.startWebServer()
}

func (c *ConduitAppImpl) attachRoutes() {
	c.Router.GET("/api/send/:deviceName/:funcName", c.WrapAuthHandler(routes.Send))
	c.Router.GET("/api/streams/:deviceName/:streamName", c.WrapAuthHandler(routes.GetStreamedMessages))
	c.Router.POST("/api/auth", c.WrapHandler(routes.Auth))
	c.Router.POST("/api/register", c.WrapHandler(routes.New))
	c.Router.GET("/api/me", c.WrapAuthHandler(routes.GetUser))
	c.Router.GET("/", routes.Hello)
	c.Router.OPTIONS("/api/*sendPath", routes.Headers)
	c.Router.ServeFiles("/static/*filepath", http.Dir("public/static"))
}

func (c *ConduitAppImpl) WrapAuthHandler(next routes.AuthHandler) httprouter.Handle {
	return middleware.ConduitAuthMiddleware(next, c.context)
}

func (c *ConduitAppImpl) WrapHandler(next routes.ConduitHandler) httprouter.Handle {
	return middleware.ConduitMiddleware(next, c.context)
}

func (c *ConduitAppImpl) startWebServer() {
	fmt.Printf("Web server to listen on port :%s", os.Getenv("PORT"))
	if c.AppConfig.IsDev {
		err := http.ListenAndServe(":"+os.Getenv("PORT"), c.Router)
		panic(err)
	} else {
		go http.ListenAndServeTLS(":443", os.Getenv("CERT"), os.Getenv("PRIV_KEY"), c.Router)
		err := http.ListenAndServe(":"+os.Getenv("PORT"), http.HandlerFunc(routes.RedirectToHttps))
		panic(err)
	}

}

func NewConduitApp() *ConduitAppImpl {
	session, err := mgo.Dial(secrets.DB_DIAL_URL)
	if err != nil {
		panic(err)
	}

	return &ConduitAppImpl{
		Router: httprouter.New(),
		AppConfig: AppConfig{
			IsDev: os.Getenv("DEV") == "TRUE",
		},
		context: &routes.HandlerContext{
			DbSession: session,
		},
	}
}
