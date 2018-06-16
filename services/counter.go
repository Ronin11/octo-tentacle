package services

import (
	"time"
	"strconv"

	"github.com/octo-tentacle/pkg/messaging"
)

func startCounter(messenger messaging.Messenger){
	go func(){
		for i := 0; true; i++ {
			messenger.WriteToChannel("counter", strconv.Itoa(i))
			duration := time.Second
  		time.Sleep(duration)
		}
	}()
}