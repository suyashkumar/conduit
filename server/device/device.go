package device

import (
	"net/http"

	"fmt"

	"github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
)

type Handler interface {
	Call(deviceName, deviceID, functionName string)
	On(deviceName, deviceID, eventName string, callback func(deviceName, eventName, body string))
	GetHTTPHandler() http.Handler
}

type handler struct {
	server *gosocketio.Server
}

var globalDeviceHandler *handler

func NewHandler() Handler {
	s := gosocketio.NewServer(transport.GetDefaultWebsocketTransport())

	// Attach socket event handlers
	s.On("hello", onHello)
	s.On(gosocketio.OnConnection, onConnection)

	globalDeviceHandler = &handler{
		server: s,
	}
	return globalDeviceHandler
}

func getRoomName(deviceName, deviceID string) string {
	return fmt.Sprintf("%s_%s", deviceID, deviceName)
}

func (h *handler) Call(deviceName, deviceID, functionName string) {
	h.server.BroadcastTo(getRoomName(deviceName, deviceID), "server_directives", functionName)
}

func (h *handler) On(deviceName, deviceID, eventName string, callback func(deviceName, eventName, body string)) {

}

func (h *handler) GetHTTPHandler() http.Handler {
	return h.server
}
