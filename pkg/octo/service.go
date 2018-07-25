package octo


import (
	"fmt"
	"os"

	"github.com/octo-tentacle/pkg/messaging"
)

func CreateService(data interface{}){
	server := os.Getenv("SERVER")
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
