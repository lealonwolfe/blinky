
### TODO 1:Setup



```go
//initializes the machines analog to digital converter
machine.InitADC()

//initializes the led from your machine pin
	led := machine.GPIO0

//configures the settings of the led
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

//initializes the potentiometer from your machine pin
	pot := machine.ADC{Pin: machine.ADC0}

    //configures the settings of the led
	pot.Configure(machine.ADCConfig{})

//sets the threshold
	threshold := uint16(39718)

```

converts the percentage to voltage reading

	(given/65535) * 100 * 3.3v

### TODO 2: Grab the potentiometer value inside a for loop


```go
//repeats infinitely
for {
    //gets value to potentiometer and sets it to variable named value
		value := pot.Get()
        

            //waits 1000 ms before running again
		time.Sleep(1000 * time.Millisecond)
	}


```

### TODO 3: below the value variable turn the led on and off if your value has reached your threshold

```go

//checks if value is greater than the threshold
		if value > threshold {
        //if greater blink the LED

            //turns on led
			led.High()
            //waits 500 ms
			time.Sleep(500 * time.Millisecond)

            //turns led off
			led.Low()
            //waits 500 ms
			time.Sleep(500 * time.Millisecond)
		} else {
            //if less turn off the led

            //turn off led
			led.Low()
		}

```