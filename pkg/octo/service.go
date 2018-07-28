package octo


import (
	"fmt"
	"time"

	"github.com/octo-tentacle/pkg/messaging"
)

// Service ...
type Service interface{
	onMessage(message string)
}

// Characteristics ...
type Characteristics struct {
	Read bool `json:"read"`
	Write bool `json:"write"`
}

// CreateService ...
func CreateService(network *Network, config *Config, data interface{}){
	server := network.GetServerAddress()

	for _, channel := range config.OutputChannels {
		messenger := messaging.CreateNatsMessenger(channel.Name, server)
		startServiceWriter(messenger, data)
	}

	discoveryMessenger := messaging.CreateNatsMessenger("discovery", server)
	startServiceDiscoveryListener(discoveryMessenger, data)
}

func startServiceDiscoveryListener(messenger messaging.Messenger, data interface{}){
	messenger.Subscribe(func(message string){
		if(message == "?"){
			messenger.Write(fmt.Sprintf("%+v", data))
		}
	})
}

func startServiceListener(messenger messaging.Messenger){
	messenger.Subscribe(func(message string){
		fmt.Printf("LISTENER MESSAGE: %s\n", message)
	})
}

func startServiceWriter(messenger messaging.Messenger, data interface{}){
	go func(){
		for i := 0; true; i++ {
			messenger.Write(fmt.Sprintf("%+v", data))
			duration := time.Second
  		time.Sleep(duration)
		}
	}()
}
