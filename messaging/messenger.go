package messaging

import (
	"github.com/octo-tentacle/messaging/nats"
)

type Messenger interface {
	WriteToChannel(channel string, message string)
	SubscribeToChannel(message string, onEvent func(channel string))
	Close()
}

func GetNatsMessenger(server string) Messenger{
	return nats.GetNatsMessenger(server)
}

