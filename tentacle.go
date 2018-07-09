package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/octo-tentacle/pkg/messaging"

	// "github.com/octo-tentacle/services/counter"
	"github.com/octo-tentacle/services/time"
)

func startServices(){
	time.StartService()
	// startCounter(messenger)
	// startCounterListener(messenger)
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

	exitSignal := make(chan os.Signal)
	signal.Notify(exitSignal, syscall.SIGINT, syscall.SIGTERM)
	<-exitSignal
}
