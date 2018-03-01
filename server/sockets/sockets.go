package sockets

import (
	gosocketio "github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
	"github.com/sirupsen/logrus"
)

func Build() *gosocketio.Server {
	server := gosocketio.NewServer(transport.GetDefaultWebsocketTransport())

	server.On("hello", func(c *gosocketio.Channel) string {
		logrus.Println("Something successfully handled")

		c.Emit("hello", "Hello emit")

		//you can return result of handler, in caller case
		//handler will be converted from "emit" to "ack"
		return "result"
	})

	server.On(gosocketio.OnConnection, func(c *gosocketio.Channel) {
		logrus.Println("New connection")
	})

	return server
}
