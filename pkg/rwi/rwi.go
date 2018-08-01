package rwi

// RWI ...
type RWI interface{}

type OutputRWI interface{
	Write(state PinState)
}

type PinState int
const (
	Low PinState = iota
	High
)

type PinMode int

const (
	Input PinMode = iota
	Output
	PWM
	// RX
	// TX
)