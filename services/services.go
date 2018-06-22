package services

import (
	"github.com/octo-tentacle/pkg/messaging"
)

func Start(messenger messaging.Messenger){
	startTempService(messenger)
	// startCounter(messenger)
	startCounterListener(messenger)
}