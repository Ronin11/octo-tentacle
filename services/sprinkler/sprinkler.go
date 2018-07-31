package sprinklerService

import (
	"time"
	"log"
	"fmt"
	"encoding/json"

	"github.com/octo-tentacle/pkg/octo"

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
	config, err = octo.ReadConfigFile("./services/sprinkler/config.json")
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
	
	// go serviceLogic(config)
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

func serviceLogic(messenger octo.Messenger){
	// for{
	// 	data.sprinklerIsOn = !data.sprinklerIsOn
	// 	time.Sleep(time.Second * 5)
	// }

	// const pin = rpio.Pin(18)
	// go func(){
	// 	// Open and map memory to access gpio, check for errors
	// 	if err := rpio.Open(); err != nil {
	// 			fmt.Println(err)
	// 			os.Exit(1)
	// 	}
 	// 	// Set pin to output mode
	// 	pin.Output()
	// 	// Unmap gpio memory when done
	// 	defer rpio.Close()
 	// 	for i := 0; true; i++ {
	// 		messenger.Write(fmt.Sprintf("%+v", service.data))
	// 		oof := i%2
	// 		if(oof > 0){
	// 			pin.Write(rpio.High)
	// 		}else{
	// 			pin.Write(rpio.Low)
	// 		}
	// 		messenger.Write(fmt.Sprintf("%+v", service.data))
	// 		duration := time.Second
  // 			time.Sleep(duration)
	// 	}
 	// }()
}
