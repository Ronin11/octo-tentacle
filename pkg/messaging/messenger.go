package messaging

import (
	"github.com/octo-tentacle/pkg/messaging/nats"
)

type Messenger interface {
	WriteToChannel(channel string, message string)
	SubscribeToChannel(message string, onEvent func(channel string))
	Close()
}

// var messengers = make(map[string]*natsMessenger)
// var once sync.Once

// func CreateNatsMessanger(server string) (Messenger, error){

// }

// func createMessenger(server string)

func GetNatsMessenger(server string) (Messenger, error){
	return nats.GetNatsMessenger(server)
}

