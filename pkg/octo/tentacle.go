package octo

import (
	"encoding/json"
	"io/ioutil"
)

type Characteristics struct {
	Read bool
	Write bool
}

// Config file struct which contains a bunch of updateable info for the services
type Config struct {
	Notify bool `json:"notify"`
	Permissions bool `json:"permissions"`
	Stream string `json:"stream"`
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