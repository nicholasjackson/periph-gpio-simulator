package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
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
			rpi.P1_15.Out(l)
			if l == gpio.High {
				l = gpio.Low
			} else {
				l = gpio.High
			}
			time.Sleep(1 * time.Second)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c)

	// Block until a signal is received.
	s := <-c

	fmt.Println("Got signal", s)
}
