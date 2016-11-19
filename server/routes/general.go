package routes

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Welcome to Suyash's Home automation server. If you know what you're doing, you can use the routes on this server to control your devices all over the world. A proper font-end for this service is coming soon.")
}
