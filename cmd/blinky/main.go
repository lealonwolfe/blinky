package main

import (
	"machine"
	"time"
)

func main() {

	machine.InitADC()

	led := machine.GPIO0
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	pot := machine.ADC{Pin: machine.ADC0}
	// (given/65535) * 100 * 3.3v
	pot.Configure(machine.ADCConfig{})

	threshold := uint16(39718)

	for {

		value := pot.Get()

		if value > threshold {

			led.High()
			time.Sleep(500 * time.Millisecond)

			led.Low()
			time.Sleep(500 * time.Millisecond)
		} else {

			led.Low()
		}

		time.Sleep(1000 * time.Millisecond)
	}
}
