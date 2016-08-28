// Our first program will print the classic "hello world"
// message. Here's the full source code.
package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"mqtt"
	"net/http"
	"time"
)

func Send(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	mqtt.SendMessage()
	time.AfterFunc(2*time.Second, func() {
		fmt.Println("hi")
		fmt.Fprintf(w, "Sent")
	})

}

func main() {
	router := httprouter.New()
	router.GET("/send", Send)
	mqtt.RunServer()
	http.ListenAndServe(":8080", router)
}
