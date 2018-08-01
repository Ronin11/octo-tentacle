package main

import (
	"fmt"
	"os"
	"time"
	"os/signal"
	"syscall"

	"github.com/octo-tentacle/pkg/octo"

	"github.com/octo-tentacle/services/sprinkler"
)

func main() {
	fmt.Println("\n~~~~~~~ Starting App ~~~~~~~")

	network := octo.JoinNetwork(os.Getenv("SERVER"), octo.NATSNetwork)
	err := octo.CreateNatsListener(network.GetServerAddress(), func(message string, subject string) {
		fmt.Printf("Subject: %s \tMessage: %s\n", subject, message)
	})
	if err != nil{
		panic(err)
	}

	sprinklerService.CreateService(network)
	// network.AddService(sprinklerService.CreateService())


	messenger := octo.CreateMessenger("discovery", network)
	go func(){
		duration := time.Second
		time.Sleep(duration)
		var services []string
		messenger.Subscribe(func(message string){
			if message != "?" {
				services = append(services, message)
			}
		})
		messenger.Write("?")
		time.Sleep(duration * 5)
		fmt.Println("Available Services: ", services)
	}()

	exitSignal := make(chan os.Signal)
	signal.Notify(exitSignal, syscall.SIGINT, syscall.SIGTERM)
	<-exitSignal
}
