package sprinklerService

import (
	"time"
	"encoding/json"

	"github.com/Ronin11/octo-tentacle/pkg/octo"
	"github.com/Ronin11/octo-tentacle/pkg/rwi"
)

type sprinklerAction struct {
	octo.Action
	State sprinklerData `json:"state"`
}

type sprinklerService struct {
	serviceCharacteristic octo.Characteristics
	data sprinklerData
	id int
	config *octo.Config
	pin rwi.OutputRWI
}

type sprinklerData struct {
	SprinklerIsOn bool `json:"sprinklerIsOn"`
	Duration string `json:"duration"`
}

// CreateService ...
func CreateService(config *octo.Config, pin rwi.OutputRWI) octo.Service{
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
		pin: pin,
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
	// go func(){
	// 	for{
	// 		service.data.SprinklerIsOn = !service.data.SprinklerIsOn
	// 		time.Sleep(time.Second * 2)
	// 	}
	// }()
	
	go func(){
 		for {
			if service.data.SprinklerIsOn {
				service.pin.Write(rwi.High)
			}else{
				service.pin.Write(rwi.Low)
			}
			duration := time.Second
  			time.Sleep(duration)
		}
 	}()
}
