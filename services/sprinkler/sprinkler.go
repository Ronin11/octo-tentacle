package sprinklerService

import (
	"log"
	"time"
	"encoding/json"

	"github.com/octo-tentacle/pkg/octo"
	// "github.com/octo-tentacle/pkg/rwi"

	"github.com/Ronin11/go-rpio"
)

type sprinklerService struct {
	serviceCharacteristic octo.Characteristics
	id int
	data sprinklerData
	config *octo.Config
}

type sprinklerData struct {
	SprinklerIsOn bool `json:"sprinklerIsOn"`
	Duration string `json:"duration"`
}

// CreateService ...
func CreateService(config *octo.Config) octo.Service{
	service := sprinklerService{
		serviceCharacteristic: octo.Characteristics{
			Read: true,
			Write: true,
		},
		id: 0,
		data: sprinklerData{
			SprinklerIsOn: false,
		},
		config: config,
	}
	
	go serviceLogic(&service)
	return &service
}

func (service *sprinklerService) GetConfig() *octo.Config{
	return service.config
}

func (service *sprinklerService) GetData() interface{}{
	return &service.data
}

func (service *sprinklerService) GetID() int{
	return service.id
}

func (service *sprinklerService) SetID(newID int){
	service.id = newID
}

func (service *sprinklerService) OnMessage(message string){
	var action sprinklerAction
	json.Unmarshal([]byte(message), &action)
	service.data = action.State
}

func (service *sprinklerService) AddToNetwork(network *octo.Network){
	octo.SetServiceId(service, network)
	octo.CreateDiscoveryListeners(service, network)
	octo.CreateServiceWriters(service, network)
	octo.CreateServiceListeners(service, network)
}

func serviceLogic(service *sprinklerService){
	go func(){
		for{
			service.data.SprinklerIsOn = !service.data.SprinklerIsOn
			time.Sleep(time.Second * 2)
		}
	}()
	
	go func(){
		// rwi.Setup()
		// pin := rwi.OutputPin(18)
		// defer rwi.Close()
		pin := rpio.Pin(4)
		pin2 := rpio.Pin(24)
		if err := rpio.Open(); err != nil {
			log.Fatal(err)
		}
	
		// Unmap gpio memory when done
		defer rpio.Close()
	
		// Set pin to output mode
		pin.Output()
		pin2.Output()

 		for {
			if service.data.SprinklerIsOn {
				pin.Write(rpio.High)
				pin2.Write(rpio.High)
				// pin.Write(rwi.High)
			}else{
				pin.Write(rpio.Low)
				pin2.Write(rpio.Low)
				// pin.Write(rwi.Low)
			}
			duration := time.Second
  			time.Sleep(duration)
		}
 	}()
}
