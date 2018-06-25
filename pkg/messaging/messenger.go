package messaging


type Messenger interface {
	Write(message string)
	Subscribe(onEvent func(message string))
}

type ChannelMessenger interface {
	WriteToChannel(channel string, message string)
	SubscribeToChannel(channel string, onEvent func(message string))
	Close(oof string)
}

func CreateNatsMessenger(channel string, server string) Messenger {
	return createChannelMessenger(channel, server, createNatsMessenger)
}

func CreateNatsListener(server string, callback func(message string, subject string)) error {
	return createNatsListener(server, callback)
}


type newChannelMessenger func(server string) (ChannelMessenger, error)

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

func createNatsMessenger(server string) (ChannelMessenger, error){
	return getNatsMessenger(server)
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





