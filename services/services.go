package services

import (
	// "github.com/octo-tentacle/pkg/messaging"

	// "github.com/octo-tentacle/services/counter"
	"github.com/octo-tentacle/services/time"
)

func Start(){
	time.StartService()
	// startCounter(messenger)
	// startCounterListener(messenger)
}