package time

import (
	"os"
	"fmt"
	"time"
	"strconv"

	"github.com/octo-tentacle/pkg/messaging"
)

func StartService(){
	server := os.Getenv("SERVER")
	messenger := messaging.CreateNatsMessenger("foo", server)
	startTime(messenger)
	startTimeListener(messenger)
}

func startTime(messenger messaging.Messenger){
	go func(){
		for i := 0; true; i++ {
			messenger.Write(strconv.Itoa(i))
			duration := time.Second
  		time.Sleep(duration)
		}
	}()
}

func startTimeListener(messenger messaging.Messenger){
	messenger.Subscribe(func(message string){
		fmt.Printf("TIME RECIEVED MESSAGE: %s\n", message)
	})
}