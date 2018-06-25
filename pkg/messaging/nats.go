package messaging

import (
	"sync"
	
	"github.com/nats-io/go-nats"
)

type natsMessenger struct {
	connection *nats.Conn
}

var messengers = make(map[string]*natsMessenger)
var once sync.Once

func getNatsMessenger(server string) (*natsMessenger, error){
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

func createNatsListener(server string, callback func(message string, subject string)) error{
	messenger, err := getNatsMessenger(server)
	messenger.connection.Subscribe("*", func(m *nats.Msg) {
		callback(string(m.Data), m.Subject)
	})
	return err
}

func (nm natsMessenger) WriteToChannel(channel string, message string){
	nm.connection.Publish(channel, []byte(message))
}

func (nm natsMessenger) SubscribeToChannel(channel string, onEvent func(message string)){
		nm.connection.Subscribe(channel, func(m *nats.Msg) {
			onEvent(string(m.Data))
		})
}

func(nm natsMessenger) Close(oof string){
	nm.connection.Close()
}
