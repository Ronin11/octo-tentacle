package main

import (
	"fmt"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/octo-tentacle/pkg/messaging"

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

	services.Start(messenger)

	defer messenger.Close()

	exitSignal := make(chan os.Signal)
	signal.Notify(exitSignal, syscall.SIGINT, syscall.SIGTERM)
	<-exitSignal
}