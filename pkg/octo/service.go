package octo


import (
	"fmt"
	"time"
)

// Service ...
type Service interface{
	// Init()
	// OnMessage(message string)
	// DiscoveryMessage() string
}

// Characteristics ...
type Characteristics struct {
	Read bool `json:"read"`
	Write bool `json:"write"`
}

// CreateService ...
func CreateService(network *Network){
	server := network.GetServerAddress()
	fmt.Println(server)
}

func startServiceDiscoveryListener(messenger Messenger, data interface{}){
	messenger.Subscribe(func(message string){
		if(message == "?"){
			messenger.Write(fmt.Sprintf("%+v", data))
		}
	})
}

func startServiceListener(messenger Messenger){
	messenger.Subscribe(func(message string){
		// OnMessage(message)
	})
}

func startServiceWriter(messenger Messenger, data interface{}){
	go func(){
		for i := 0; true; i++ {
			messenger.Write(fmt.Sprintf("%+v", data))
			duration := time.Second
  		time.Sleep(duration)
		}
	}()
}
