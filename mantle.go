package main

import (
	"fmt"
	"os"
	// "time"
	"os/signal"
	"syscall"

	"github.com/octo-tentacle/pkg/octo"

	"github.com/octo-tentacle/services/sprinkler"
)

func main() {
	fmt.Println("\n~~~~~~~ Starting App ~~~~~~~")

	network := octo.JoinNetwork(os.Getenv("SERVER"), octo.NATSNetwork)
	// err := octo.CreateNatsListener(network.GetServerAddress(), func(message string, subject string) {
	// 	fmt.Printf("Subject: %s \tMessage: %s\n", subject, message)
	// })
	// if err != nil{
	// 	panic(err)
	// }

	sprinklerService.CreateService(network)


	// messenger := octo.CreateMessenger("discovery", network)
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
