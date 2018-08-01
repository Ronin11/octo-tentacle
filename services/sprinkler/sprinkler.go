package sprinklerService

import (
	"time"

	"github.com/octo-tentacle/pkg/octo"

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

var service sprinklerService

// CreateService ...
func CreateService() octo.Service{
	service = sprinklerService{
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

// func OnMessage(message string) string{
// 	var temp sprinklerData
// 	err := json.Unmarshal([]byte(message), &temp)
// 	if err != nil {
// 		return err.Error()
// 	}
// 	service.data = temp
// 	return "okMessage"
// }

func serviceLogic(service *sprinklerService){
	for{
		service.data.SprinklerIsOn = !service.data.SprinklerIsOn
		time.Sleep(time.Second)
	}

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
