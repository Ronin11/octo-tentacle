package octo


import (
	"fmt"
	"time"
	"log"
	"encoding/json"
)

// Service ...
type Service interface {
	AddToNetwork(*Network)
	GetData() interface{}
	GetConfig() *Config
	GetID() int
	SetID(int)
}

// Characteristics ...
type Characteristics struct {
	Read bool `json:"read"`
	Write bool `json:"write"`
}

func SetServiceId(service Service, network *Network) {
	var id = 0
	idMessenger := CreateMessenger(
		fmt.Sprintf("%s.discovery", service.GetConfig().Service), 
		network,
	)
	var services []string
	idMessenger.Subscribe(func(message string){
		if message != "?" {
			services = append(services, message)
		}
	})
	idMessenger.Write("?")
	time.Sleep(time.Second * 2)
	if len(services) > 0 {
		id = len(services)
	}
	service.SetID(id)
}

func CreateServicWriters(service Service, network *Network){
	for _, channel := range service.GetConfig().OutputChannels {
		messenger := CreateMessenger(
			fmt.Sprintf("%s.%s.%d.%s", 
				service.GetConfig().Service, service.GetConfig().Group, 
				service.GetID(), channel.Name,
			), 
			network,
		)
		go func(){
			for i := 0; true; i++ {
				data, err := json.Marshal(service.GetData())
				if err != nil {
					log.Fatal(err)
				}
				messenger.Write(string(data))
				//TODO SEND DATA ON INTERVAL DEFINED BY CONFIG
				duration := time.Second
				time.Sleep(duration)
			}
		}()
	}
}

// const okMessage = "Ok"
func CreateServiceListeners(service Service, network *Network){
	for _, channelName := range service.GetConfig().InputChannels {
		messenger := CreateMessenger(
			fmt.Sprintf("%s.%s.%d.%s", 
				service.GetConfig().Service, service.GetConfig().Group, 
				service.GetID(), channelName,
			),
			network,
		)
		messenger.Subscribe(func(message string){
			//TODO MAKE THIS NOT BAD
			fmt.Println(
				fmt.Sprintf("%s.%s.%d.%s: %s", 
					service.GetConfig().Service, service.GetConfig().Group, 
					service.GetID(), channelName, message,
				),
			)
			// if message != okMessage {
			// 	messenger.Write(OnMessage(message))
			// }
		})
	}
}

func CreateDiscoveryListeners(service Service, network *Network){
	serviceDiscoveryMessenger := CreateMessenger(
		fmt.Sprintf("%s.discovery", service.GetConfig().Service),
		network,
	)
	serviceDiscoveryMessenger.Subscribe(func(message string){
		if(message == "?"){
			serviceDiscoveryMessenger.Write(
				fmt.Sprintf("%s.%s.%d:%+v", 
					service.GetConfig().Service, service.GetConfig().Group, 
					service.GetID(), service.GetData(),
				),
			)
		}
	})
	discoveryMessenger := CreateMessenger("discovery", network)
	discoveryMessenger.Subscribe(func(message string){
		if(message == "?"){
			discoveryMessenger.Write(
				fmt.Sprintf("%s.%s.%d:%+v", 
					service.GetConfig().Service, service.GetConfig().Group, 
					service.GetID(), service.GetData(),
				),
			)
		}
	})
}



// CreateService ...
// func CreateService(network *Network){
// 	server := network.GetServerAddress()
// 	fmt.Println(server)
// }

// func startServiceDiscoveryListener(messenger Messenger, data interface{}){
// 	messenger.Subscribe(func(message string){
// 		if(message == "?"){
// 			messenger.Write(fmt.Sprintf("%+v", data))
// 		}
// 	})
// }

// func startServiceListener(messenger Messenger){
// 	messenger.Subscribe(func(message string){
// 		// OnMessage(message)
// 	})
// }

// func startServiceWriter(messenger Messenger, data interface{}){
// 	go func(){
// 		for i := 0; true; i++ {
// 			messenger.Write(fmt.Sprintf("%+v", data))
// 			duration := time.Second
//   		time.Sleep(duration)
// 		}
// 	}()
// }
