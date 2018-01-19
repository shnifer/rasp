package main

import (
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	r := raspi.NewAdaptor()
	led := gpio.NewLedDriver(r, "12")

	work := func() {
		gobot.Every(2*time.Second, func() {
			per:=time.After(time.Second)
			for{
				select{
					case <-time.After(10*time.Millisecond):
						led.Toggle()
					case <- per:
						return
				}
			}
		})
	}

	robot := gobot.NewRobot("blinkBot",
		[]gobot.Connection{r},
		[]gobot.Device{led},
		work,
	)

	robot.Start()
}