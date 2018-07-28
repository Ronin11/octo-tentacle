package sprinklerService

import (
	// "time"
	"log"

	"github.com/octo-tentacle/pkg/octo"
)

type sprinklerData struct {
	serviceCharacteristic octo.Characteristics
	sprinklerIsOn bool
}

var data sprinklerData

// CreateService ...
func CreateService() *octo.Service{
	config, err := octo.ReadConfigFile("./services/sprinkler/config.json")
	if err != nil{
		log.Fatal(err)
	}
	data.serviceCharacteristic = octo.Characteristics{
		Read: true,
		Write: true,
	}
	
	// octo.CreateService(network, config, &data)
	
	// go serviceLogic(config)
	return struct {
		onMessage func(string)
	}
}

func serviceLogic(config *octo.Config){
	// for{
	// 	data.sprinklerIsOn = !data.sprinklerIsOn
	// 	time.Sleep(time.Second * 5)
	// }
}

func onMessage(message string){

}
