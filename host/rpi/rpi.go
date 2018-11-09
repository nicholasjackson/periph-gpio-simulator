package rpi

import (
	"periph.io/x/periph"
	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/pin"
)

var (
	// Raspberry Pi A and B, 26 pin header:
	P1_1  pin.Pin    = &Pin{num: 1}  // max 30mA
	P1_2  pin.Pin    = &Pin{num: 2}  // (filtered)
	P1_3  gpio.PinIO = &Pin{num: 3}  // High, I2C1_SDA
	P1_4  pin.Pin    = &Pin{num: 4}  //
	P1_5  gpio.PinIO = &Pin{num: 5}  // High, I2C1_SCL
	P1_6  pin.Pin    = &Pin{num: 6}  //
	P1_7  gpio.PinIO = &Pin{num: 7}  // High, CLK0
	P1_8  gpio.PinIO = &Pin{num: 8}  // Low,  UART0_TX, UART1_TX
	P1_9  pin.Pin    = &Pin{num: 9}  //
	P1_10 gpio.PinIO = &Pin{num: 10} // Low,  UART0_RX, UART1_RX
	P1_11 gpio.PinIO = &Pin{num: 11} // Low,  UART0_RTS, SPI1_CS1, UART1_RTS
	P1_12 gpio.PinIO = &Pin{num: 12} // Low,  I2S_SCK, SPI1_CS0, PWM0
	P1_13 gpio.PinIO = &Pin{num: 13} // Low,
	P1_14 pin.Pin    = &Pin{num: 14} //
	P1_15 gpio.PinIO = &Pin{num: 15} // Low,
	P1_16 gpio.PinIO = &Pin{num: 16} // Low,
	P1_17 pin.Pin    = &Pin{num: 17} //
	P1_18 gpio.PinIO = &Pin{num: 18} // Low,
	P1_19 gpio.PinIO = &Pin{num: 19} // Low, SPI0_MOSI
	P1_20 pin.Pin    = &Pin{num: 20} //
	P1_21 gpio.PinIO = &Pin{num: 21} // Low, SPI0_MISO
	P1_22 gpio.PinIO = &Pin{num: 22} // Low,
	P1_23 gpio.PinIO = &Pin{num: 23} // Low, SPI0_CLK
	P1_24 gpio.PinIO = &Pin{num: 24} // High, SPI0_CS0
	P1_25 pin.Pin    = &Pin{num: 25} //
	P1_26 gpio.PinIO = &Pin{num: 26} // High, SPI0_CS1

	// Raspberry Pi A+, B+, 2 and later, 40 pin header (also named J8):
	P1_27 gpio.PinIO = &Pin{} // High, I2C0_SDA used to probe for HAT EEPROM, see https://github.com/raspberrypi/hats
	P1_28 gpio.PinIO = &Pin{} // High, I2C0_SCL
	P1_29 gpio.PinIO = &Pin{} // High, CLK1
	P1_30 pin.Pin    = &Pin{} //
	P1_31 gpio.PinIO = &Pin{} // High, CLK2
	P1_32 gpio.PinIO = &Pin{} // Low,  PWM0
	P1_33 gpio.PinIO = &Pin{} // Low,  PWM1
	P1_34 pin.Pin    = &Pin{} //
	P1_35 gpio.PinIO = &Pin{} // Low,  I2S_WS, SPI1_MISO, PWM1
	P1_36 gpio.PinIO = &Pin{} // Low,  UART0_CTS, SPI1_CS2, UART1_CTS
	P1_37 gpio.PinIO = &Pin{} //
	P1_38 gpio.PinIO = &Pin{} // Low,  I2S_DIN, SPI1_MOSI, CLK0
	P1_39 pin.Pin    = &Pin{} //
	P1_40 gpio.PinIO = &Pin{} // Low,  I2S_DOUT, SPI1_CLK, CLK1

	// P5 header on Raspberry Pi A and B, PCB v2:
	P5_1 pin.Pin    = &Pin{}
	P5_2 pin.Pin    = &Pin{}
	P5_3 gpio.PinIO = &Pin{} // Float, I2C0_SDA, I2S_SCK
	P5_4 gpio.PinIO = &Pin{} // Float, I2C0_SCL, I2S_WS
	P5_5 gpio.PinIO = &Pin{} // Low,   I2S_DIN, UART0_CTS, UART1_CTS
	P5_6 gpio.PinIO = &Pin{} // Low,   I2S_DOUT, UART0_RTS, UART1_RTS
	P5_7 pin.Pin    = &Pin{}
	P5_8 pin.Pin    = &Pin{}

	AUDIO_RIGHT         = &Pin{} // Low,   PWM0, SPI2_MISO, UART1_TX
	AUDIO_LEFT          = &Pin{} // Low,   PWM1, SPI2_MOSI, UART1_RX
	HDMI_HOTPLUG_DETECT = &Pin{} // High,
)

var pins = []*Pin{
	P1_1.(*Pin),
	P1_2.(*Pin),
	P1_3.(*Pin),
	P1_4.(*Pin),
	P1_5.(*Pin),
	P1_6.(*Pin),
	P1_7.(*Pin),
	P1_8.(*Pin),
	P1_9.(*Pin),
	P1_10.(*Pin),
	P1_11.(*Pin),
	P1_12.(*Pin),
	P1_13.(*Pin),
	P1_14.(*Pin),
	P1_15.(*Pin),
	P1_16.(*Pin),
	P1_17.(*Pin),
	P1_18.(*Pin),
	P1_19.(*Pin),
	P1_20.(*Pin),
	P1_21.(*Pin),
	P1_22.(*Pin),
	P1_23.(*Pin),
	P1_24.(*Pin),
	P1_25.(*Pin),
	P1_26.(*Pin),
	P1_27.(*Pin),
	P1_28.(*Pin),
	P1_29.(*Pin),
	P1_30.(*Pin),
	P1_31.(*Pin),
	P1_32.(*Pin),
	P1_33.(*Pin),
	P1_34.(*Pin),
	P1_35.(*Pin),
	P1_36.(*Pin),
	P1_37.(*Pin),
	P1_38.(*Pin),
	P1_39.(*Pin),
	P1_40.(*Pin),
	P5_1.(*Pin),
	P5_2.(*Pin),
	P5_3.(*Pin),
	P5_4.(*Pin),
	P5_5.(*Pin),
	P5_6.(*Pin),
	P5_7.(*Pin),
	P5_8.(*Pin),
	AUDIO_RIGHT,
	AUDIO_LEFT,
	HDMI_HOTPLUG_DETECT,
}

// Present returns if the current driver is present
func Present() bool {
	return true
}

// driver implements periph.Driver.
type driver struct {
}

// String returns a string representation for the driver
func (d *driver) String() string {
	return "rpi-simulator"
}

// Prerequisites returns the prerequisites for the driver
func (d *driver) Prerequisites() []string {
	return nil
}

// After does something
func (d *driver) After() []string {
	return []string{"gpio-simulator"}
}

// Init the driver
func (d *driver) Init() (bool, error) {
	return true, nil
}

func init() {
	periph.MustRegister(&drv)
}

var drv driver
