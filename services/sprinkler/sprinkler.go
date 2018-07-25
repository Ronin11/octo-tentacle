package sprinklerService

import (
	"log"
	// "os"
	"fmt"
	"time"

	"github.com/octo-tentacle/pkg/messaging"
	"github.com/octo-tentacle/pkg/octo"
)

type sprinklerData struct {
	sprinklerCharacteristic octo.Characteristics
	sprinklerIsOn bool
}

const serviceChannel = "sprinkler"
var data sprinklerData


// Start ...
func Start(){
	config, err := octo.ReadConfigFile("./services/sprinkler/config.json")
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(config.Triggers)
	// fmt.Println(fmt.Sprintf("%+v", data))

	// server := os.Getenv("SERVER")
	octo.CreateService(data)

	// for _, channel := range config.OutputChannels {
	// 	messenger := messaging.CreateNatsMessenger(channel, server)
	// 	startWriter(messenger)
	// }
}

// func startListener(messenger messaging.Messenger){
// 	messenger.Subscribe(func(message string){
// 		fmt.Printf("LISTENER MESSAGE: %s\n", message)
// 	})
// }


func startWriter(messenger messaging.Messenger){
	go func(){
		for i := 0; true; i++ {
			messenger.Write("Sprinkler")
			duration := time.Second
  		time.Sleep(duration)
		}
	}()
}

