package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"mqtt"
	"net/http"
	"time"
)

type RpcResponse struct {
	Success bool   `json:"success"`
	Data    string `json:"data"`
}

func Send(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	mqtt.SendMessage(ps.ByName("deviceName"), ps.ByName("funcName"))
	c := make(chan string)
	end := make(chan string)

	mqtt.Register(ps.ByName("deviceName")+"/device", func(topic string, payload string) {
		res := &RpcResponse{
			Success: true,
			Data:    payload,
		}
		resBytes, _ := json.Marshal(res)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, string(resBytes))
		c <- "done"
	})
	time.AfterFunc(2*time.Second, func() {
		fmt.Fprintf(w, "ERROR")
		end <- "done"
	})
	select {
	case <-c:
		return
	case <-end:
		return
	}

}

func main() {
	router := httprouter.New()
	router.GET("/send/:deviceName/:funcName", Send)
	mqtt.RunServer()
	http.ListenAndServe(":8080", router)
}
