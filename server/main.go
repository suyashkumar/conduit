package main

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/suyashkumar/conduit/server/config"
	"github.com/suyashkumar/conduit/server/log"
	"github.com/suyashkumar/conduit/server/routes"
	"github.com/suyashkumar/conduit/server/sockets"
)

func main() {
	log.Configure()

	r := routes.Build()
	s := sockets.Build()

	r.Handler("GET", "/socket.io/", s)

	p := fmt.Sprintf(":%s", config.Get(config.Port))

	if config.Get(config.UseSSL) == "false" {
		logrus.WithField("port", p).Info("Serving without SSL")
		err := http.ListenAndServe(p, r)
		logrus.Fatal(err)
	} else {
		logrus.Info("Serving with SSL")
		err := http.ListenAndServeTLS(
			p,
			config.Get(config.CertKey),
			config.Get(config.PrivKey),
			r,
		)
		// TODO: reroute http requests to https
		logrus.Fatal(err)
	}

}
