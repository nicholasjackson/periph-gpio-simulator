package rpi

import (
	"fmt"
	"time"

	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpiostream"
	"periph.io/x/periph/conn/physic"
	"periph.io/x/periph/conn/pin"
)

// Pin is a simulated pin
type Pin struct {
	num   int
	level gpio.Level
	edge  gpio.Edge
}

//DefaultPull something
func (p *Pin) DefaultPull() gpio.Pull {
	return gpio.PullUp
}

// Drive something
func (p *Pin) Drive() physic.ElectricCurrent {
	return 0
}

// FastOut something
func (p *Pin) FastOut(l gpio.Level) {

}

// FastRead something
func (p *Pin) FastRead() gpio.Level {
	return p.level
}

// Func something
func (p *Pin) Func() pin.Func {
	return ""
}

// Function something
func (p *Pin) Function() string {
	return ""
}

// Halt something
func (p *Pin) Halt() error {
	return nil
}

// Hysteresis something
func (p *Pin) Hysteresis() bool {
	return true
}

// In something
func (p *Pin) In(pull gpio.Pull, edge gpio.Edge) error {
	if pull == gpio.PullDown {
		p.level = gpio.Low
	}

	if pull == gpio.PullUp {
		p.level = gpio.High
	}

	p.edge = edge

	return nil
}

// Name something
func (p *Pin) Name() string {
	return fmt.Sprintf("|%02d|", p.num)
}

// Number something
func (p *Pin) Number() int {
	return p.num
}

// Out something
func (p *Pin) Out(l gpio.Level) error {
	p.level = l
	return nil
}

// PWM something
func (p *Pin) PWM(duty gpio.Duty, freq physic.Frequency) error {
	return nil
}

// Pull something
func (p *Pin) Pull() gpio.Pull {
	return gpio.PullUp
}

// Read something
func (p *Pin) Read() gpio.Level {
	return p.level
}

// SetFunc something
func (p *Pin) SetFunc(f pin.Func) error {
	return nil
}

// SlewLimit something
func (p *Pin) SlewLimit() bool {
	return true
}

// StreamIn something
func (p *Pin) StreamIn(pull gpio.Pull, s gpiostream.Stream) error {
	return nil
}

// StreamOut something
func (p *Pin) StreamOut(s gpiostream.Stream) error {
	return nil
}

// String something
func (p *Pin) String() string {
	return ""
}

// SupportedFuncs something
func (p *Pin) SupportedFuncs() []pin.Func {
	return []pin.Func{}
}

// WaitForEdge somethings
func (p *Pin) WaitForEdge(timeout time.Duration) bool {
	c := p.Read() // get the initial state
	running := true

	if timeout != -1 {
		time.AfterFunc(timeout, func() {
			running = false
		})
	}

	for running {
		s := p.Read()
		// if we are waiting for a rising edge or both
		if s != c && s == gpio.High && (p.edge == gpio.RisingEdge || p.edge == gpio.BothEdges) {
			return true
		}

		// if we are waiting for a falling edge or both
		if s != c && s == gpio.Low && (p.edge == gpio.FallingEdge || p.edge == gpio.BothEdges) {
			return true
		}
	}

	return false
}
