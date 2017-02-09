package main

import (
	"github.com/suyashkumar/conduit/server/app"
)

func main() {
	conduitApp := app.NewConduitApp() // Init a new conduit web service
	conduitApp.Run()                  // Run the conduit web service server
}
