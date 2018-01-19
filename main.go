package main

import (
	"time"

	//"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
	"fmt"
	"strconv"
)

func main() {
	//d := 1000//milliseconds
	r := raspi.NewAdaptor()
	led := gpio.NewLedDriver(r, "12")
/*
	work := func() {
		gobot.Every(2*time.Second, func() {
			per:=time.After(1*time.Second)
			for{
				select{
					case <-time.After(time.Duration(d)*time.Millisecond):
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
*/
	periodCh:= make(chan int)
	go func(ch chan int){
		d:=<-ch
		fmt.Println("Set",d,"millisecond blink")
//		per:=time.After(1*time.Second)
		for{
			select{
			case d=<-ch:
				fmt.Println("Set",d,"millisecond blink")
			case <-time.After(time.Duration(d)*time.Millisecond):
				led.Toggle()
//			case <- per:
//				return
			}
		}
	} (periodCh)

	var str string
	for {
		fmt.Scanln(&str)
		fmt.Println("you wrote ", str)
		switch str {
		case "quit", "exit":
			//robot.Stop()
			return
		default:
			fmt.Println("try to conv")
			v,err:=strconv.Atoi(str)
			if err==nil {
				periodCh<-v
			} else {
				fmt.Println("but can't")
			}
		}
	}
}