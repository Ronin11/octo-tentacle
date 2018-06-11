package main

import (
	"fmt"
	"os"

	"github.com/octo-tentacle/messaging"
	"github.com/octo-tentacle/gpio"
)

func test(message string){
	fmt.Printf("MESSAGE: %s\n", message)
}

func main() {
	fmt.Println("Starting App")
	serverAddress := os.Args[1]
	messenger := messaging.GetNatsMessenger(serverAddress)

	messenger.SubscribeToChannel("foo", test)
	messenger.WriteToChannel("foo", "WOOO")

	messenger2 := messaging.GetNatsMessenger(serverAddress)
	messenger2.WriteToChannel("foo", "WOOO2")

	gpio.Setup()

	defer messenger.Close()
	defer gpio.Cleanup()

	for{

	}
}