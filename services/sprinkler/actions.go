package sprinklerService

import (
	"github.com/Ronin11/octo-tentacle/pkg/octo"
)

type SprinklerAction struct {
	octo.Action
	State SprinklerData `json:"state"`
}

func (s SprinklerAction) GetChannel() string{
	return s.Channel
}