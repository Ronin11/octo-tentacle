package main

import (
	// "time"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Ronin11/octo-tentacle/pkg/octo"
	"github.com/Ronin11/octo-tentacle/pkg/rwi"

	"github.com/Ronin11/octo-tentacle/services/sprinkler"
)

func main() {
	fmt.Println("\n~~~~~~~ Starting Tentacle ~~~~~~~")

	rwi.Setup()
	sprinklerPin0 := rwi.OutputPin(24)
	sprinklerPin1 := rwi.OutputPin(25)
	defer rwi.Close()

	// sprinklerPin0 := rwi.TestOutputRWI{}
	// sprinklerPin1 := rwi.TestOutputRWI{}

	sprinklerConfig0 := octo.ReadConfigFile("./services/sprinkler/config.json")
	sprinklerConfig1 := octo.ReadConfigFile("./services/sprinkler/config.json")
	sprinklerService0 := sprinklerService.CreateService(sprinklerConfig0, sprinklerPin0)
	sprinklerService1 := sprinklerService.CreateService(sprinklerConfig1, sprinklerPin1)
	
	network := octo.JoinNetwork(os.Getenv("SERVER"), octo.NATSNetwork)
	network.AddService(sprinklerService0)
	network.AddService(sprinklerService1)

	exitSignal := make(chan os.Signal)
	signal.Notify(exitSignal, syscall.SIGINT, syscall.SIGTERM)
	<-exitSignal
}
