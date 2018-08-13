package main

import (
	"time"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/octo-tentacle/pkg/octo"
	"github.com/octo-tentacle/pkg/rwi"

	"github.com/octo-tentacle/services/sprinkler"
)

func main() {
	fmt.Println("\n~~~~~~~ Starting App ~~~~~~~")

	// rwi.Setup()
	// sprinklerPin1 := rwi.OutputPin(24)
	// defer rwi.Close()

	sprinklerConfig1 := octo.ReadConfigFile("./services/sprinkler/config.json")
	// sprinklerService1 := sprinklerService.CreateService(sprinklerConfig1, sprinklerPin1)
	sprinklerService1 := sprinklerService.CreateService(sprinklerConfig1, rwi.TestOutputRWI{})
	network := octo.JoinNetwork(os.Getenv("SERVER"), octo.NATSNetwork)
	network.AddService(sprinklerService1)
	// network.AddService(sprinklerService.CreateService(config))


	messenger := octo.CreateMessenger("sprinkler.backyard.0.input", network)
	duration := time.Second
	time.Sleep(duration)
	messenger.Write(`{"Name":"Action Description","State":{"sprinklerIsOn": true,"Duration":"SomeDuration"},"onDone":{"name": "ON DONE"}}`)
	messenger.Subscribe(func(message string){
		fmt.Println("RESPONSE: ", message)
	})

	exitSignal := make(chan os.Signal)
	signal.Notify(exitSignal, syscall.SIGINT, syscall.SIGTERM)
	<-exitSignal
}
