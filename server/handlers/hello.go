package handlers

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/suyashkumar/conduit/server/device"
)

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params, d device.Handler) {
	d.Call("Suyash", "ID", "myFunc")
	fmt.Fprintf(w, "Hello, world")
}
