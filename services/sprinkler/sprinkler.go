package sprinklerService

import (
	"time"
	"log"
	"fmt"
	"encoding/json"

	"github.com/octo-tentacle/pkg/octo"
	"github.com/octo-tentacle/pkg/rwi"

	// "github.com/Ronin11/go-rpio"
)

type sprinklerService struct {
	serviceCharacteristic octo.Characteristics
	data sprinklerData
}

type sprinklerData struct {
	SprinklerIsOn bool `json:"sprinklerIsOn"`
}

var config *octo.Config
var service sprinklerService

// CreateService ...
func CreateService(network *octo.Network) octo.Service{
	const configFile = "./services/sprinkler/config.json"
	service = sprinklerService{
		serviceCharacteristic: octo.Characteristics{
			Read: true,
			Write: true,
		},
		data: sprinklerData{
			SprinklerIsOn: false,
		},
	}
	var id = 0
	var err error
	config, err = octo.ReadConfigFile(configFile)
	if err != nil{
		log.Fatal(err)
	}

	idMessenger := octo.CreateMessenger(
		fmt.Sprintf("%s.discovery", config.Service), 
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

	
	for _, channel := range config.OutputChannels {
		messenger := octo.CreateMessenger(
			fmt.Sprintf("%s.%s.%d.%s", config.Service, config.Group, id, channel.Name), 
			network,
		)
		startServiceWriter(messenger)
	}

	for _, channelName := range config.InputChannels {
		messenger := octo.CreateMessenger(
			fmt.Sprintf("%s.%s.%d.%s", config.Service, config.Group, id, channelName),
			network,
		)
		startServiceListener(messenger)

		// go func(){
		// 	for i := 0; true; i++ {
		// 		if i % 2 == 0 {
		// 			messenger.Write(`{"sprinklerIsOn":true}`)
		// 		} else {
		// 			messenger.Write(`{"sprinklerIsOn":false}`)
		// 		}
		// 		duration := time.Second
		// 		time.Sleep(duration)
		// 	}
		// }()
	}
	serviceDiscoveryMessenger := octo.CreateMessenger(
		fmt.Sprintf("%s.discovery", config.Service),
		network,
	)
	serviceDiscoveryMessenger.Subscribe(func(message string){
		if(message == "?"){
			serviceDiscoveryMessenger.Write(
				fmt.Sprintf("%s.%s.%d:%+v", config.Service, config.Group, id, service.data),
			)
		}
	})
	discoveryMessenger := octo.CreateMessenger("discovery", network)
	discoveryMessenger.Subscribe(func(message string){
		if(message == "?"){
			discoveryMessenger.Write(
				fmt.Sprintf("%s.%s.%d:%+v", config.Service, config.Group, id, service.data),
			)
		}
	})
	
	go serviceLogic()
	return &service
}

const okMessage = "Ok"
func startServiceListener(messenger octo.Messenger){
	messenger.Subscribe(func(message string){
		if message != okMessage {
			messenger.Write(OnMessage(message))
		}
	})
}

func startServiceWriter(messenger octo.Messenger){
	go func(){
		for i := 0; true; i++ {
			data, err := json.Marshal(service.data)
			if err != nil {
				log.Fatal(err)
			}
			messenger.Write(string(data))
			duration := time.Second
  		time.Sleep(duration)
		}
	}()
}

func OnMessage(message string) string{
	var temp sprinklerData
	err := json.Unmarshal([]byte(message), &temp)
	if err != nil {
		return err.Error()
	}
	service.data = temp
	return okMessage
}

func serviceLogic(){
	go func(){
		for{
			service.data.SprinklerIsOn = !service.data.SprinklerIsOn
			time.Sleep(time.Second * 2)
		}
	}()
	
	go func(){
		rwi.Setup()
		pin := rwi.OutputPin(12)
		defer rwi.Close()
 		for {
			if service.data.SprinklerIsOn {
				pin.Write(rwi.High)
			}else{
				pin.Write(rwi.Low)
			}
			duration := time.Second
  			time.Sleep(duration)
		}
 	}()
}
