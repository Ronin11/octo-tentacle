package sprinklerService

import (
	"github.com/octo-tentacle/pkg/octo"
)

type sprinklerAction struct {
	octo.Action
	State sprinklerData `json:"state"`
}