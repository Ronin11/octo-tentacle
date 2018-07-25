package messaging


//Publics
type ChannelMessenger interface {
	WriteToChannel(channel string, message string)
	SubscribeToChannel(channel string, onEvent func(message string)) func() error
	Close(oof string)
}

func CreateNatsMessenger(channel string, server string) Messenger {
	return createChannelMessenger(channel, server, createNatsMessenger)
}

func CreateNatsListener(server string, callback func(message string, subject string)) error {
	return createNatsListener(server, callback)
}

//Privates
type newChannelMessenger func(server string) (ChannelMessenger, error)

type channelMessenger struct {
	channel string
	unsub func() error
	WriteToChannel func(channel string, message string)
	SubscribeToChannel func(channel string, onEvent func(message string)) func() error
	Close func(oof string)
}

func (c channelMessenger) Write(message string){
	c.WriteToChannel(c.channel, message)
}

func (c channelMessenger) WriteAndListen(message string, onEvent func(message string)){
	c.WriteToChannel(c.channel, message)
	c.unsub = c.SubscribeToChannel(c.channel, func(msg string){
		onEvent(msg)
		c.unsub()
	})
}

func (c channelMessenger) Subscribe(onEvent func(message string)){
	c.unsub = c.SubscribeToChannel(c.channel, onEvent)
}

func (c channelMessenger) Unsubscribe(){
	if(c.unsub != nil){
		c.unsub()
	}
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
