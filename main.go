package main

import (
	"time"

	//"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
	"fmt"
	"strconv"
)

func MakePulse(led *gpio.LedDriver, d int) {
	led.On()
	time.AfterFunc(time.Duration(d)*time.Millisecond, func(){led.Off()})
	time.Sleep(10*time.Millisecond)
}

func ledDriver(led *gpio.LedDriver ,durCh chan int){
	d:=<-durCh
	fmt.Println("Set",d,"millisecond blink")
	for{
		select{
			//Изменение параметра дурации
			case d=<-durCh:
				fmt.Println("Set",d,"millisecond blink")
			default:
				MakePulse(led, d)
		}
	}
}

func main() {
	r := raspi.NewAdaptor()
	led := gpio.NewLedDriver(r, "12")

	periodCh:= make(chan int)

	go ledDriver(led, periodCh)

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