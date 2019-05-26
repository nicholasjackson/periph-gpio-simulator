package main

import (
	"log"
	"os"
	"time"

	"github.com/nicholasjackson/periph-gpio-simulator/host/rpi"
	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/host"
)

func main() {
	logger := log.New(os.Stdout, "", log.Lmicroseconds)

	// Load all the drivers:
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	go func() {
		l := gpio.High
		for {
			logger.Println("Set pin", l)
			rpi.SO_51.Out(l)
			if l == gpio.High {
				l = gpio.Low
			} else {
				l = gpio.High
			}
			time.Sleep(1 * time.Second)
		}
	}()

	rpi.SO_9.In(gpio.PullDown, gpio.BothEdges)
	rpi.SO_9.WaitForEdge(-1)
	logger.Println("Pin state changed")

	for {
	}
}
