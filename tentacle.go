package main

import (
	"fmt"
	"flag"

	"github.com/octo-tentacle/pkg/messaging"
	"github.com/octo-tentacle/pkg/gpio"

	"github.com/octo-tentacle/services"
)

func main() {
	fmt.Println("Starting App")
	
	server := flag.String("server", "http://127.0.0.1:4200", "NATS server address")
	flag.Parse()

	messenger, err := messaging.GetNatsMessenger(*server)
	if(err != nil){
		panic(err)
	}

	gpio.Setup()

	services.Start(messenger)

	defer messenger.Close()
	defer gpio.Cleanup()

	for{

	}
}