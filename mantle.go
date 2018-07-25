package main

import (
	"fmt"
	"time"
	"os"
	"os/signal"
	"syscall"

	"github.com/octo-tentacle/pkg/messaging"

	"github.com/octo-tentacle/services/sprinkler"
	"github.com/octo-tentacle/services/soil"
	// "github.com/octo-tentacle/services/temp"
)

func startServices(){
	sprinklerService.Start()
	soilService.Start()
	// tempService.Start()
}

func main() {
	fmt.Println("\n~~~~~~~ Starting App ~~~~~~~")

	server := os.Getenv("SERVER")
	err := messaging.CreateNatsListener(server, func(message string, subject string) {
		fmt.Printf("Subject: %s \tMessage: %s\n", subject, message)
	})
	if err != nil{
		panic(err)
	}

	startServices()


	messenger := messaging.CreateNatsMessenger("discovery", server)
	go func(){
		duration := time.Second * 5
		time.Sleep(duration)
		messenger.WriteAndListen("?", func(msg string){
			fmt.Println("asdf: ", msg)
		})
	}()

	exitSignal := make(chan os.Signal)
	signal.Notify(exitSignal, syscall.SIGINT, syscall.SIGTERM)
	<-exitSignal
}
