package sprinklerService

import (
	"time"

	"github.com/octo-tentacle/pkg/octo"
	"github.com/octo-tentacle/pkg/rwi"

	// "github.com/Ronin11/go-rpio"
)

type sprinklerService struct {
	serviceCharacteristic octo.Characteristics
	id int
	data sprinklerData
	config *octo.Config
}

type sprinklerData struct {
	SprinklerIsOn bool `json:"sprinklerIsOn"`
}

// CreateService ...
func CreateService() octo.Service{
	service := sprinklerService{
		serviceCharacteristic: octo.Characteristics{
			Read: true,
			Write: true,
		},
		id: 0,
		data: sprinklerData{
			SprinklerIsOn: false,
		},
		config: octo.ReadConfigFile("./services/sprinkler/config.json"),
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

func (service *sprinklerService) AddToNetwork(network *octo.Network){
	octo.SetServiceId(service, network)
	octo.CreateDiscoveryListeners(service, network)
	octo.CreateServicWriters(service, network)
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
