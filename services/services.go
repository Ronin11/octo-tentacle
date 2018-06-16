package services

import (
	"github.com/octo-tentacle/pkg/messaging"
)

func Start(messenger messaging.Messenger){
	startCounter(messenger)
	startCounterListener(messenger)
}