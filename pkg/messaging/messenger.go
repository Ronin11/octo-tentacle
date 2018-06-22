package messaging

import (
	"github.com/octo-tentacle/pkg/messaging/nats"
)

// type 

type Messenger interface {
	Write(message string)
	Subscribe(onEvent func(message string))
}

type ChannelMessenger interface {
	WriteToChannel(channel string, message string)
	SubscribeToChannel(channel string, onEvent func(message string))
	Close(oof string)
}

type channelMessenger struct {
	channel string
	WriteToChannel func(channel string, message string)
	SubscribeToChannel func(channel string, onEvent func(message string))
	Close func(oof string)
}

func (c channelMessenger) Write(message string){
	c.WriteToChannel(c.channel, message)
}

func (c channelMessenger) Subscribe(onEvent func(message string)){
	c.SubscribeToChannel(c.channel, onEvent)
}

type newChannelMessenger func(server string) (ChannelMessenger, error)

func createNatsMessenger(server string) (ChannelMessenger, error){
	return nats.GetNatsMessenger(server)
}

func createChannelMessenger(channel string, server string, construction newChannelMessenger) Messenger{
	messenger, err := construction(server)
	if(err != nil){
		panic(err)
	}

	return channelMessenger{
		channel: channel,
		WriteToChannel: messenger.WriteToChannel,
		SubscribeToChannel: messenger.SubscribeToChannel,
		Close: messenger.Close,
	}
}

func CreateNatsMessenger(channel string, server string) Messenger {
	return createChannelMessenger(channel, server, createNatsMessenger)
}





