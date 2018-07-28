package octo

import (
	"encoding/json"
	"io/ioutil"
)

// Action ...
type Action struct {
	Name string `json:"name"`
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
	OutputChannels []OutputChannel `json:"outputChannels"`
	Triggers []Trigger `json:"triggers"`
}

// ReadConfigFile ...
func ReadConfigFile(filename string) (*Config, error){
	file, e := ioutil.ReadFile(filename)
	if e != nil {
			return nil, e
	}

	var config Config
	json.Unmarshal(file, &config)
	return &config, nil
}
