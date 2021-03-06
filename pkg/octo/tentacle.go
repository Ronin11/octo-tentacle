package octo

import (
	"log"
	"encoding/json"
	"io/ioutil"
)

// Action ...
type Action struct {
	Channel string `json:"channel"`
	Name string `json:"name"`
	OnDone interface{} `json:"onDone"`
}

// GetChannel ...
func (action *Action) GetChannel() string{
	return action.Channel
}

// Trigger ...
type Trigger struct {
	ChannelName string `json:"channel"`
	Logic string `json:"logic"`
	TargetAction Action `json:"action"`
	//Threshold
	//Matcher
}

//OutputChannel ...
type OutputChannel struct {
	Name string `json:"name"`
	Verbosity int `json:"verbosity"`
}

// Config file struct which contains a bunch of updateable info for the services
type Config struct {
	Service string `json:"service"`
	Group string `json:"group"`
	InputChannels []string `json:"inputChannels"`
	OutputChannels []OutputChannel `json:"outputChannels"`
	Triggers []Trigger `json:"triggers"`
}

// ReadConfigFile ...
func ReadConfigFile(filename string) (*Config){
	file, e := ioutil.ReadFile(filename)
	if e != nil {
			log.Fatal(e)
	}

	var config Config
	json.Unmarshal(file, &config)
	return &config
}
