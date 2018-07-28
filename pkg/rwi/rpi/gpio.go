package rpi

import (
	"os"

	"github.com/Ronin11/go-rpio"
)

struct Gpio struct{
	io rpio
}

func Setup() Gpio{
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return rpio
}

func PinIO(pinNumber int, )