package rwi

import (
	"log"

	"github.com/Ronin11/go-rpio"
)

// type Gpio struct{
// 	io rpio
// }

func OutputPin(pinNumber int) OutputRWI{
	pin := rpio.Pin(pinNumber)
	return outputIO{
		pin: pin,
	}
}

type outputIO struct{
	pin rpio.Pin
}

func (io outputIO)Write(state PinState){
	switch state {
		case Low:
			io.pin.Write(rpio.Low)
		case High:
			io.pin.Write(rpio.High)
	}
}

type inputIO struct{
	pin rpio.Pin
}

type pwmIO struct{
	pin rpio.Pin
}

func Setup() {
	if err := rpio.Open(); err != nil {
		log.Fatal(err)
	}
	// return rpio
}

func Close() {
	rpio.Close()
}