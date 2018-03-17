package handlers

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Headers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	SetCorsHeaders(w)
	if r.Method == "OPTIONS" {
		return
	}
}

func SetCorsHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "x-access-token, Content-Type")
}
