package time

import (
	"os"
	"time"
	"strconv"

	"github.com/octo-tentacle/pkg/messaging"
)

func StartService(){
	server := os.Getenv("SERVER")
	messenger := messaging.CreateNatsMessenger("foo", server)
	startTime(messenger)
	messenger2 := messaging.CreateNatsMessenger("foo.two", server)
	startTime(messenger2)
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

