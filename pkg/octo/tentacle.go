package octo

import (
	"encoding/json"
	"io/ioutil"
)

type Characteristics struct {
	Read bool
	Write bool
}

type Trigger struct {
	ChannelName string `json:"channel"`
	Logic string `json:"logic"`

	// Action
	//Threshold
	//Matcher
}

// Config file struct which contains a bunch of updateable info for the services
type Config struct {
	Notify bool `json:"notify"`
	Permissions bool `json:"permissions"`
	DataLevel int `json:"dataLevel"`
	OutputChannels []string `json:"outputChannels"`
	Triggers []Trigger `json:"triggers"`
}

func ReadConfigFile(filename string) (*Config, error){
	file, e := ioutil.ReadFile(filename)
	if e != nil {
			return nil, e
	}

	var config Config
	json.Unmarshal(file, &config)
	return &config, nil
}
