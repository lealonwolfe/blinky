package main

import (
	"fmt"
	"machine"
	"time"

	"tinygo.org/x/drivers/dht"
	"tinygo.org/x/drivers/hd44780i2c"
)

func printForever(err error) {
	for {
		println(err)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {

	machine.I2C0.Configure(machine.I2CConfig{
		Frequency: 100 * machine.KHz,
	})

	lcd := hd44780i2c.New(machine.I2C0, 0x27)
	lcd.Configure(hd44780i2c.Config{
		Width:       16, // required
		Height:      2,  // required
		CursorOn:    true,
		CursorBlink: true,
	})

	lcd.Print([]byte(" TinyGo\n  LCD Test "))

	time.Sleep(time.Millisecond * 7000)
	//////////////////////////////////
	machine.InitADC()

	pin := machine.GP4
	dhtSensor := dht.New(pin, dht.DHT22)

	led := machine.GPIO0
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	pot := machine.ADC{Pin: machine.ADC0}
	// (given/65535) * 100 * 3.3v
	pot.Configure(machine.ADCConfig{})

	threshold := uint16(39718)

	for {

		value := pot.Get()
		voltage := (float32(value) / 65535.0) * 100.0 * 3.3
		temp, hum, err := dhtSensor.Measurements()
		println("VALUE", value, "VOLTAGE", voltage)

		if err == nil {
			fmt.Printf("Temperature: %02d.%d°C, Humidity: %02d.%d%%\n", temp/10, temp%10, hum/10, hum%10)
		} else {
			fmt.Printf("Could not take measurements from the sensor: %s\n", err.Error())
		}

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
