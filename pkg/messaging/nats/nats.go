package nats

import (
	"sync"
	
	"github.com/nats-io/go-nats"
)

type natsMessenger struct {
	connection *nats.Conn
}

var messengers = make(map[string]*natsMessenger)
var once sync.Once

func GetNatsMessenger(server string) (*natsMessenger, error){
	if messengers[server] == nil {
		var nc *nats.Conn
		var err error
		once.Do(func() {
			nc, err = nats.Connect(server)
			messengers[server] = &natsMessenger{connection: nc}
		})
		if(err != nil){
			return nil, err
		}
	}
	return messengers[server], nil
}

func (nm natsMessenger) WriteToChannel(channel string, message string){
	nm.connection.Publish(channel, []byte(message))
}

func (nm natsMessenger) SubscribeToChannel(channel string, onEvent func(channel string)){
		nm.connection.Subscribe(channel, func(m *nats.Msg) {
			onEvent(string(m.Data))
		})
}

func(nm natsMessenger) Close(){
	nm.connection.Close()
}
