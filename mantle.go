package main

import (
	"github.com/octo-tentacle/pkg/octo"
	"fmt"
	// "time"
	"os"
	"os/signal"
	"syscall"

	"github.com/octo-tentacle/pkg/messaging"

	"github.com/octo-tentacle/services/sprinkler"
)

func main() {
	fmt.Println("\n~~~~~~~ Starting App ~~~~~~~")

	network := octo.CreateNetwork(os.Getenv("SERVER"))
	err := messaging.CreateNatsListener(network.GetServerAddress(), func(message string, subject string) {
		fmt.Printf("Subject: %s \tMessage: %s\n", subject, message)
	})
	if err != nil{
		panic(err)
	}

	sprinklerService.CreateService(network)


	// messenger := messaging.CreateNatsMessenger("discovery", server)
	// go func(){
	// 	duration := time.Second * 5
	// 	time.Sleep(duration)
	// 	messenger.WriteAndListen("?", func(msg string){
	// 		fmt.Println("asdf: ", msg)
	// 	})
	// }()

	exitSignal := make(chan os.Signal)
	signal.Notify(exitSignal, syscall.SIGINT, syscall.SIGTERM)
	<-exitSignal
}
