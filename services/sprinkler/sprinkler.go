package sprinklerService

import (
	"time"
	"log"
	"fmt"

	"github.com/octo-tentacle/pkg/octo"

	"github.com/Ronin11/go-rpio"
)

type sprinklerService struct {
	serviceCharacteristic octo.Characteristics
	data sprinklerData
}

type sprinklerData struct {
	sprinklerIsOn bool
}

var config octo.Config
var service sprinklerService

// CreateService ...
func CreateService(network *octo.Network) octo.Service{
	service = sprinklerService{
		serviceCharacteristic: octo.Characteristics{
			Read: true,
			Write: true,
		},
		data: sprinklerData{
			sprinklerIsOn: false,
		},
	}

	config, err := octo.ReadConfigFile("./services/sprinkler/config.json")
	if err != nil{
		log.Fatal(err)
	}

	for _, channel := range config.OutputChannels {
		messenger := octo.CreateMessenger(channel.Name, network)
		startServiceWriter(messenger)
	}

	discoveryMessenger := octo.CreateMessenger("discovery", network)
	startServiceDiscoveryListener(discoveryMessenger)
	
	// go serviceLogic(config)
	return &service
}

func startServiceDiscoveryListener(messenger octo.Messenger){
	messenger.Subscribe(func(message string){
		if(message == "?"){
			messenger.Write(fmt.Sprintf("%+v", service.data))
		}
	})
}

func startServiceListener(messenger octo.Messenger){
	messenger.Subscribe(func(message string){
		// OnMessage(message)
	})
}

func startServiceWriter(messenger octo.Messenger){
	go func(){
		for i := 0; true; i++ {
			messenger.Write(fmt.Sprintf("%+v", service.data))
			duration := time.Second
  		time.Sleep(duration)
		}
	}()
}

func serviceLogic(messenger octo.Messenger){
	// for{
	// 	data.sprinklerIsOn = !data.sprinklerIsOn
	// 	time.Sleep(time.Second * 5)
	// }
	const pin = rpio.Pin(18)
	go func(){
		// Open and map memory to access gpio, check for errors
		if err := rpio.Open(); err != nil {
				fmt.Println(err)
				os.Exit(1)
		}
 		// Set pin to output mode
		pin.Output()
		// Unmap gpio memory when done
		defer rpio.Close()
 		for i := 0; true; i++ {
			messenger.Write(fmt.Sprintf("%+v", service.data))
			oof := i%2
			if(oof > 0){
				pin.Write(rpio.High)
			}else{
				pin.Write(rpio.Low)
			}
			messenger.Write(fmt.Sprintf("%+v", service.data))
			duration := time.Second
  			time.Sleep(duration)
		}
 	}()
}

// func (s sprinklerService) Init(){

// }

// func (s sprinklerService) OnMessage(message string){
// 	fmt.Println("MSG: ", message)
// }
