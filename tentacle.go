package main

import (
	"fmt"

	"gpio"
)

func main() {
	fmt.Println("Starting App")
	gpio.Setup()

	defer gpio.Cleanup()
}