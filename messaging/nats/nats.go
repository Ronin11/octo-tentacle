package nats

import (
	"fmt"
	"sync"
	
	"github.com/nats-io/go-nats"
)

type natsMessenger struct {
	connection *nats.Conn
}

var messengers = make(map[string]*natsMessenger)
var once sync.Once

func GetNatsMessenger(server string) *natsMessenger{
	if messengers[server] == nil {
		once.Do(func() {
			nc, err := nats.Connect(server)
			if(err != nil){
				fmt.Println(err)
			}
			messengers[server] = &natsMessenger{connection: nc}
		})
	}
	return messengers[server]
}

func (nm natsMessenger) WriteToChannel(channel string, message string){
	nm.connection.Publish(channel, []byte(message))
}

func (nm natsMessenger) SubscribeToChannel(channel string, onEvent func(channel string)){
		// Simple Async Subscriber
		nm.connection.Subscribe(channel, func(m *nats.Msg) {
			onEvent(string(m.Data))
		})
}

func(nm natsMessenger) Close(){
	nm.connection.Close()
}
