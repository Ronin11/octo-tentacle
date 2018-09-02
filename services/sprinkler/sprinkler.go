package sprinklerService

import (
	"fmt"
	// "os"
	"time"
	"encoding/json"

	"github.com/Ronin11/octo-tentacle/pkg/octo"
	"github.com/Ronin11/octo-tentacle/pkg/rwi"
)

type SprinklerData struct {
	SprinklerIsOn bool `json:"sprinklerIsOn"`
	Duration int64 `json:"duration"`
}

type sprinklerService struct {
	serviceCharacteristic octo.Characteristics
	data SprinklerData
	id int
	config *octo.Config
	pin rwi.OutputRWI
}

// CreateService ...
func CreateService(config *octo.Config, pin rwi.OutputRWI) octo.Service{
	service := sprinklerService{
		serviceCharacteristic: octo.Characteristics{
			Read: true,
			Write: true,
		},
		id: 0,
		data: SprinklerData{
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
	var action SprinklerAction
	json.Unmarshal([]byte(message), &action)
	service.data = action.State

	timer := time.NewTimer(time.Duration(service.data.Duration) * time.Second)
	go func() {
		<-timer.C
		service.data.Duration = 0
		service.data.SprinklerIsOn = false
		
		v, ok := action.OnDone.(octo.Action)
		fmt.Println(v)
		fmt.Println(ok)

		// var nextAction SprinklerAction
		// json.Unmarshal([]byte(action.OnDone), &nextAction)
		// fmt.Println(nextAction)
		// if ok {
		// 	fmt.Println(onDone.GetChannel())
		// 	// network := octo.JoinNetwork(os.Getenv("SERVER"), octo.NATSNetwork)
		// 	// octo.SendAction(onDone, network)
		// }
	}()
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
