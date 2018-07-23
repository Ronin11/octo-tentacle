package sprinklerService

import (
	"log"
	"os"
	"fmt"
	"time"

	"github.com/octo-tentacle/pkg/messaging"
	"github.com/octo-tentacle/pkg/octo"
)

type sprinklerDatum struct {
	sprinklerCharacteristic octo.Characteristics
	sprinklerIsOn bool
}

const serviceChannel = "sprinkler"
var data []sprinklerDatum


// Start ...
func Start(){
	config, err := octo.ReadConfigFile("./services/sprinkler/config.json")
	if err != nil{
		log.Fatal(err)
	}
	config.Notify = true
	// fmt.Println(config.Stream)
	data = append(data, sprinklerDatum{
		sprinklerCharacteristic: octo.Characteristics{Read: true, Write: true},
		sprinklerIsOn: false,
	})
	// fmt.Println(fmt.Sprintf("%+v", data))

	server := os.Getenv("SERVER")
	messenger := messaging.CreateNatsMessenger(serviceChannel, server)
	startWriter(messenger)
	startListener(messenger)

	queryMessenger := messaging.CreateNatsMessenger(fmt.Sprintf("%s.services", serviceChannel), server)
	startQueryListener(queryMessenger)
}

func startQueryListener(messenger messaging.Messenger){
	messenger.Subscribe(func(message string){
		// fmt.Printf("QUERY MESSAGE: %s\n", message)
		if(message == "?"){
			messenger.Write(fmt.Sprintf("%+v", data))
		}
	})
}

func startListener(messenger messaging.Messenger){
	messenger.Subscribe(func(message string){
		fmt.Printf("LISTENER MESSAGE: %s\n", message)
	})
}


func startWriter(messenger messaging.Messenger){
	go func(){
		for i := 0; true; i++ {
			// messenger.Write(strconv.Itoa(i))
			duration := time.Second
  		time.Sleep(duration)
		}
	}()
}

