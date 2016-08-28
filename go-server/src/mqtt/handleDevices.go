package mqtt

import (
	"fmt"
	"github.com/surgemq/message"
	"github.com/surgemq/surgemq/service"
	"time"
)

var mClient *service.Client

func onPublish(msg *message.PublishMessage) error {
	fmt.Println("message arrived")
	fmt.Println(msg.String())
	fmt.Println(string(msg.Payload()))
	return nil
}
func createServerClient() *service.Client {
	client := &service.Client{}
	msg := message.NewConnectMessage()
	msg.SetClientId([]byte("surgemq"))
	msg.SetKeepAlive(30)
	msg.SetCleanSession(true)
	msg.SetWillTopic([]byte("will"))
	msg.SetWillMessage([]byte("send me home"))
	msg.SetVersion(3)
	if err := client.Connect("tcp://:1883", msg); err != nil {
		fmt.Println("problem")
		fmt.Println(err)
	}

	submsg := message.NewSubscribeMessage()
	submsg.AddTopic([]byte("#"), 0)
	client.Subscribe(submsg, nil, onPublish)
	//client.Subscribe(msg2, nil, nil)

	pubMsg := message.NewPublishMessage()
	pubMsg.SetTopic([]byte("suyash1"))
	pubMsg.SetPayload(make([]byte, 10))
	client.Publish(pubMsg, nil)
	return client
}

func sendMessage(client *service.Client) {
	pubMsg := message.NewPublishMessage()
	pubMsg.SetTopic([]byte("suyash1"))
	pubMsg.SetPayload([]byte("hello"))
	client.Publish(pubMsg, nil)
}

func SendMessage() {
	sendMessage(mClient)
}

func RunServer() {
	fmt.Println("Starting up...")
	svr := &service.Server{
		KeepAlive: 300,
	}
	go svr.ListenAndServe("tcp://:1883")
	mClient = createServerClient()
	time.AfterFunc(100*time.Second, SendMessage)

}
