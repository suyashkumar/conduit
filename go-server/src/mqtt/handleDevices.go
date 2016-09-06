package mqtt

import (
	"fmt"
	"github.com/surgemq/message"
	"github.com/surgemq/surgemq/service"
	"time"
)

var mClient *service.Client
var handlerMap = make(map[string]func(string, string))

func Register(name string, a func(string, string)) {
	handlerMap[name] = a
}

func onPublish(msg *message.PublishMessage) error {
	/*
		fmt.Println("message arrived")
		fmt.Println(msg.String())
		fmt.Println(string(msg.Payload()))
	*/
	if val, ok := handlerMap[string(msg.Topic())]; ok {
		val(string(msg.Topic()), string(msg.Payload()))
	}
	return nil
}
func createServerClient() *service.Client {
	client := &service.Client{}
	msg := message.NewConnectMessage()
	msg.SetClientId([]byte("surgemq"))
	msg.SetKeepAlive(300)
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

func sendMessage(client *service.Client, device string, payload string) {
	pubMsg := message.NewPublishMessage()
	pubMsg.SetTopic([]byte(device))
	pubMsg.SetPayload([]byte(payload))
	client.Publish(pubMsg, nil)
}

func SendMessage(device string, payload string) {
	sendMessage(mClient, device, payload)
}

func stayAlive() {
	SendMessage("stayinAlive", "")
	time.AfterFunc(50*time.Second, stayAlive)
}

func RunServer() {
	fmt.Println("Starting up MQTT machinery...")
	svr := &service.Server{
		KeepAlive: 300,
	}
	go svr.ListenAndServe("tcp://:1883")
	time.Sleep(200*time.Millisecond)
	mClient = createServerClient()
	go stayAlive()
	fmt.Println("Started and listening")
}
