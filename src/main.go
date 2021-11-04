package main

import (
	"fmt"
	"time"

	rpio "github.com/stianeikeland/go-rpio"
)

func positionToInt(x int, y int) int {
	//TODO implement translation from coordinates to binary for led array
	return 16
}

func main() {
	fmt.Println("opening gpio")
	err := rpio.Open()
	if err != nil {
		panic(fmt.Sprint("unable to open gpio", err.Error()))
	}

	defer rpio.Close()

	// Test connection to gpio works
	pin := rpio.Pin(18)
	pin.Output()
	pin.Toggle()
	time.Sleep(time.Second / 1)
	pin.Toggle()

	// Test call to method works
	intToSendToPi := positionToInt(0, 0)
	pin = rpio.Pin(intToSendToPi)
	pin.Output()

	for x := 0; x < 20; x++ {
		pin.Toggle()
		time.Sleep(time.Second / 5)
	}
}
