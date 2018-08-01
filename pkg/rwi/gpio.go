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

// func PinIO(pinNumber int, mode PinMode) RWI{
// 	pin := rpio.Pin(pinNumber)
// 	fmt.Println(pin)
// 	switch mode {
// 		case Input:
// 			return inputIO{
// 				pin: pin,
// 			}
// 		case Output:
// 			return outputIO{
// 				pin: pin,
// 			}
// 		case PWM:
// 			return pwmIO{
// 				pin: pin,
// 			}
// 		default:
// 			return nil
// 	}
// }