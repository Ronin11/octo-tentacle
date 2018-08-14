package octo

import (
	"log"
	"encoding/json"
)

type Messenger interface {
	Write(message string)
	WriteAndListen(message string, onEvent func(message string))
	Subscribe(onEvent func(message string))
	Unsubscribe()
}

func SendAction(action *Action, network *Network){
	message, err := json.Marshal(action)
	if err != nil {
		log.Fatal(err)
	}
	messenger := CreateMessenger(action.Channel,network)
	messenger.Write(string(message))
}
