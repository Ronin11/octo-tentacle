package octo

import (
	"log"
)


/*~~~~~~ Publics ~~~~~~*/

// ChannelMessenger ...
type ChannelMessenger interface {
	WriteToChannel(channel string, message string)
	SubscribeToChannel(channel string, onEvent func(message string)) func() error
	Close(oof string)
}

// CreateMessenger ...
func CreateMessenger(channel string, network *Network) Messenger {
	switch networkType := network.GetNetworkType(); networkType {
		case NATSNetwork:
			connection, err := getNatsMessenger(network.GetServerAddress())
			if err != nil {
				log.Fatal("CreateMessenger: ", err)
			}
			
			return channelMessenger{
				channel: channel,
				WriteToChannel: connection.WriteToChannel,
				SubscribeToChannel: connection.SubscribeToChannel,
				Close: connection.Close,
			}
		default:
			return nil
	}
}

// CreateListener ...
func CreateListener(network *Network, callback func(message string, subject string)) error {
	switch networkType := network.GetNetworkType(); networkType {
	case NATSNetwork:
		return createNatsListener(network.GetServerAddress(), callback)
	default:
		return nil
	}
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
