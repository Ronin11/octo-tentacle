package services

import (
	"fmt"
	"time"

	"github.com/octo-tentacle/pkg/messaging"

	"github.com/yryz/ds18b20"
)

func startTempService(messenger messaging.Messenger) {
	go func(){
		sensors, err := ds18b20.Sensors()
		if err != nil {
			panic(err)
		}
		for{
			for _, sensor := range sensors {
				t, err := ds18b20.Temperature(sensor)
				if err == nil {
					messenger.WriteToChannel("counter", fmt.Sprintf("sensor: %s temperature: %.2f°C\n", sensor, t))
				}
			}
			time.Sleep(time.Second)
		}
	}()
}