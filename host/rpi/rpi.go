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

var (
	SO_1   pin.Pin    = pin.GROUND    // GND
	SO_2   pin.Pin    = pin.INVALID   // EMMC_DISABLE_N
	SO_3   gpio.PinIO = &Pin{num: 0}  // GPIO0
	SO_4   pin.Pin    = pin.INVALID   // NC, SDX_VDD, NC
	SO_5   gpio.PinIO = &Pin{num: 1}  // GPIO1
	SO_6   pin.Pin    = pin.INVALID   // NC, SDX_VDD, NC
	SO_7   pin.Pin    = pin.GROUND    // GND
	SO_8   pin.Pin    = pin.GROUND    // GND
	SO_9   gpio.PinIO = &Pin{num: 2}  // GPIO2
	SO_10  pin.Pin    = pin.INVALID   // NC, SDX_CLK, NC
	SO_11  gpio.PinIO = &Pin{num: 3}  // GPIO3
	SO_12  pin.Pin    = pin.INVALID   // NC, SDX_CMD, NC
	SO_13  pin.Pin    = pin.GROUND    // GND
	SO_14  pin.Pin    = pin.GROUND    // GND
	SO_15  gpio.PinIO = &Pin{num: 4}  // GPIO4
	SO_16  pin.Pin    = pin.INVALID   // NC, SDX_D0, NC
	SO_17  gpio.PinIO = &Pin{num: 5}  // GPIO5
	SO_18  pin.Pin    = pin.INVALID   // NC, SDX_D1, NC
	SO_19  pin.Pin    = pin.GROUND    // GND
	SO_20  pin.Pin    = pin.GROUND    // GND
	SO_21  gpio.PinIO = &Pin{num: 6}  // GPIO6
	SO_22  pin.Pin    = pin.INVALID   // NC, SDX_D2, NC
	SO_23  gpio.PinIO = &Pin{num: 7}  // GPIO7
	SO_24  pin.Pin    = pin.INVALID   // NC, SDX_D3, NC
	SO_25  pin.Pin    = pin.GROUND    // GND
	SO_26  pin.Pin    = pin.GROUND    // GND
	SO_27  gpio.PinIO = &Pin{num: 8}  // GPIO8
	SO_28  gpio.PinIO = &Pin{num: 28} // GPIO28
	SO_29  gpio.PinIO = &Pin{num: 9}  // GPIO9
	SO_30  gpio.PinIO = &Pin{num: 29} // GPIO29
	SO_31  pin.Pin    = pin.GROUND    // GND
	SO_32  pin.Pin    = pin.GROUND    // GND
	SO_33  gpio.PinIO = &Pin{num: 10} // GPIO10
	SO_34  gpio.PinIO = &Pin{num: 30} // GPIO30
	SO_35  gpio.PinIO = &Pin{num: 11} // GPIO11
	SO_36  gpio.PinIO = &Pin{num: 31} // GPIO31
	SO_37  pin.Pin    = pin.GROUND    // GND
	SO_38  pin.Pin    = pin.GROUND    // GND
	SO_39  pin.Pin    = pin.DC_IN     // GPIO0-27_VDD
	SO_40  pin.Pin    = pin.DC_IN     // GPIO0-27_VDD
	SO_41  pin.Pin    = pin.DC_IN     // GPIO28-45_VDD
	SO_42  pin.Pin    = pin.DC_IN     // GPIO28-45_VDD
	SO_43  pin.Pin    = pin.GROUND    // GND
	SO_44  pin.Pin    = pin.GROUND    // GND
	SO_45  gpio.PinIO = &Pin{num: 12} // GPIO12
	SO_46  gpio.PinIO = &Pin{num: 32} // GPIO32
	SO_47  gpio.PinIO = &Pin{num: 13} // GPIO13
	SO_48  gpio.PinIO = &Pin{num: 33} // GPIO33
	SO_49  pin.Pin    = pin.GROUND    // GND
	SO_50  pin.Pin    = pin.GROUND    // GND
	SO_51  gpio.PinIO = &Pin{num: 14} // GPIO14
	SO_52  gpio.PinIO = &Pin{num: 34} // GPIO34
	SO_53  gpio.PinIO = &Pin{num: 15} // GPIO15
	SO_54  gpio.PinIO = &Pin{num: 35} // GPIO35
	SO_55  pin.Pin    = pin.GROUND    // GND
	SO_56  pin.Pin    = pin.GROUND    // GND
	SO_57  gpio.PinIO = &Pin{num: 16} // GPIO16
	SO_58  gpio.PinIO = &Pin{num: 36} // GPIO36
	SO_59  gpio.PinIO = &Pin{num: 17} // GPIO17
	SO_60  gpio.PinIO = &Pin{num: 37} // GPIO37
	SO_61  pin.Pin    = pin.GROUND    // GND
	SO_62  pin.Pin    = pin.GROUND    // GND
	SO_63  gpio.PinIO = &Pin{num: 18} // GPIO18
	SO_64  gpio.PinIO = &Pin{num: 38} // GPIO38
	SO_65  gpio.PinIO = &Pin{num: 19} // GPIO19
	SO_66  gpio.PinIO = &Pin{num: 39} // GPIO39
	SO_67  pin.Pin    = pin.GROUND    // GND
	SO_68  pin.Pin    = pin.GROUND    // GND
	SO_69  gpio.PinIO = &Pin{num: 20} // GPIO20
	SO_70  gpio.PinIO = &Pin{num: 40} // GPIO40
	SO_71  gpio.PinIO = &Pin{num: 21} // GPIO21
	SO_72  gpio.PinIO = &Pin{num: 41} // GPIO41
	SO_73  pin.Pin    = pin.GROUND    // GND
	SO_74  pin.Pin    = pin.GROUND    // GND
	SO_75  gpio.PinIO = &Pin{num: 22} // GPIO22
	SO_76  gpio.PinIO = &Pin{num: 42} // GPIO42
	SO_77  gpio.PinIO = &Pin{num: 23} // GPIO23
	SO_78  gpio.PinIO = &Pin{num: 43} // GPIO43
	SO_79  pin.Pin    = pin.GROUND    // GND
	SO_80  pin.Pin    = pin.GROUND    // GND
	SO_81  gpio.PinIO = &Pin{num: 24} // GPIO24
	SO_82  gpio.PinIO = &Pin{num: 44} // GPIO44
	SO_83  gpio.PinIO = &Pin{num: 25} // GPIO25
	SO_84  gpio.PinIO = &Pin{num: 45} // GPIO45
	SO_85  pin.Pin    = pin.GROUND    // GND
	SO_86  pin.Pin    = pin.GROUND    // GND
	SO_87  gpio.PinIO = &Pin{num: 26} // GPIO26
	SO_88  pin.Pin    = pin.INVALID   // HDMI_HPD_N_1V8, HDMI_HPD_N_1V8, GPIO46_1V8
	SO_89  gpio.PinIO = &Pin{num: 27} // GPIO27
	SO_90  pin.Pin    = pin.INVALID   // EMMC_EN_N_1V8, EMMC_EN_N_1V8, GPIO47_1V8
	SO_91  pin.Pin    = pin.GROUND    // GND
	SO_92  pin.Pin    = pin.GROUND    // GND
	SO_93  pin.Pin    = pin.INVALID   // DSI0_DN1
	SO_94  pin.Pin    = pin.INVALID   // DSI1_DP0
	SO_95  pin.Pin    = pin.INVALID   // DSI0_DP1
	SO_96  pin.Pin    = pin.INVALID   // DSI1_DN0
	SO_97  pin.Pin    = pin.GROUND    // GND
	SO_98  pin.Pin    = pin.GROUND    // GND
	SO_99  pin.Pin    = pin.INVALID   // DSI0_DN0
	SO_100 pin.Pin    = pin.INVALID   // DSI1_CP
	SO_101 pin.Pin    = pin.INVALID   // DSI0_DP0
	SO_102 pin.Pin    = pin.INVALID   // DSI1_CN
	SO_103 pin.Pin    = pin.GROUND    // GND
	SO_104 pin.Pin    = pin.GROUND    // GND
	SO_105 pin.Pin    = pin.INVALID   // DSI0_CN
	SO_106 pin.Pin    = pin.INVALID   // DSI1_DP3
	SO_107 pin.Pin    = pin.INVALID   // DSI0_CP
	SO_108 pin.Pin    = pin.INVALID   // DSI1_DN3
	SO_109 pin.Pin    = pin.GROUND    // GND
	SO_110 pin.Pin    = pin.GROUND    // GND
	SO_111 pin.Pin    = pin.INVALID   // HDMI_CLK_N
	SO_112 pin.Pin    = pin.INVALID   // DSI1_DP2
	SO_113 pin.Pin    = pin.INVALID   // HDMI_CLK_P
	SO_114 pin.Pin    = pin.INVALID   // DSI1_DN2
	SO_115 pin.Pin    = pin.GROUND    // GND
	SO_116 pin.Pin    = pin.GROUND    // GND
	SO_117 pin.Pin    = pin.INVALID   // HDMI_D0_N
	SO_118 pin.Pin    = pin.INVALID   // DSI1_DP1
	SO_119 pin.Pin    = pin.INVALID   // HDMI_D0_P
	SO_120 pin.Pin    = pin.INVALID   // DSI1_DN1
	SO_121 pin.Pin    = pin.GROUND    // GND
	SO_122 pin.Pin    = pin.GROUND    // GND
	SO_123 pin.Pin    = pin.INVALID   // HDMI_D1_N
	SO_124 pin.Pin    = pin.INVALID   // NC
	SO_125 pin.Pin    = pin.INVALID   // HDMI_D1_P
	SO_126 pin.Pin    = pin.INVALID   // NC
	SO_127 pin.Pin    = pin.GROUND    // GND
	SO_128 pin.Pin    = pin.INVALID   // NC
	SO_129 pin.Pin    = pin.INVALID   // HDMI_D2_N
	SO_130 pin.Pin    = pin.INVALID   // NC
	SO_131 pin.Pin    = pin.INVALID   // HDMI_D2_P
	SO_132 pin.Pin    = pin.INVALID   // NC
	SO_133 pin.Pin    = pin.GROUND    // GND
	SO_134 pin.Pin    = pin.GROUND    // GND
	SO_135 pin.Pin    = pin.INVALID   // CAM1_DP3
	SO_136 pin.Pin    = pin.INVALID   // CAM0_DP0
	SO_137 pin.Pin    = pin.INVALID   // CAM1_DN3
	SO_138 pin.Pin    = pin.INVALID   // CAM0_DN0
	SO_139 pin.Pin    = pin.GROUND    // GND
	SO_140 pin.Pin    = pin.GROUND    // GND
	SO_141 pin.Pin    = pin.INVALID   // CAM1_DP2
	SO_142 pin.Pin    = pin.INVALID   // CAM0_CP
	SO_143 pin.Pin    = pin.INVALID   // CAM1_DN2
	SO_144 pin.Pin    = pin.INVALID   // CAM0_CN
	SO_145 pin.Pin    = pin.GROUND    // GND
	SO_146 pin.Pin    = pin.GROUND    // GND
	SO_147 pin.Pin    = pin.INVALID   // CAM1_CP
	SO_148 pin.Pin    = pin.INVALID   // CAM0_DP1
	SO_149 pin.Pin    = pin.INVALID   // CAM1_CN
	SO_150 pin.Pin    = pin.INVALID   // CAM0_DN1
	SO_151 pin.Pin    = pin.GROUND    // GND
	SO_152 pin.Pin    = pin.GROUND    // GND
	SO_153 pin.Pin    = pin.INVALID   // CAM1_DP1
	SO_154 pin.Pin    = pin.INVALID   // NC
	SO_155 pin.Pin    = pin.INVALID   // CAM1_DN1
	SO_156 pin.Pin    = pin.INVALID   // NC
	SO_157 pin.Pin    = pin.GROUND    // GND
	SO_158 pin.Pin    = pin.INVALID   // NC
	SO_159 pin.Pin    = pin.INVALID   // CAM1_DP0
	SO_160 pin.Pin    = pin.INVALID   // NC
	SO_161 pin.Pin    = pin.INVALID   // CAM1_DN0
	SO_162 pin.Pin    = pin.INVALID   // NC
	SO_163 pin.Pin    = pin.GROUND    // GND
	SO_164 pin.Pin    = pin.GROUND    // GND
	SO_165 pin.Pin    = pin.INVALID   // USB_DP
	SO_166 pin.Pin    = pin.INVALID   // TVDAC
	SO_167 pin.Pin    = pin.INVALID   // USB_DM
	SO_168 pin.Pin    = pin.INVALID   // USB_OTGID
	SO_169 pin.Pin    = pin.GROUND    // GND
	SO_170 pin.Pin    = pin.GROUND    // GND
	SO_171 pin.Pin    = pin.INVALID   // HDMI_CEC
	SO_172 pin.Pin    = pin.INVALID   // VC_TRST_N
	SO_173 pin.Pin    = pin.INVALID   // HDMI_SDA
	SO_174 pin.Pin    = pin.INVALID   // VC_TDI
	SO_175 pin.Pin    = pin.INVALID   // HDMI_SCL
	SO_176 pin.Pin    = pin.INVALID   // VC_TMS
	SO_177 pin.Pin    = pin.INVALID   // RUN
	SO_178 pin.Pin    = pin.INVALID   // VC_TDO
	SO_179 pin.Pin    = pin.INVALID   // VDD_CORE (DO NOT CONNECT)
	SO_180 pin.Pin    = pin.INVALID   // VC_TCK
	SO_181 pin.Pin    = pin.GROUND    // GND
	SO_182 pin.Pin    = pin.GROUND    // GND
	SO_183 pin.Pin    = pin.V1_8      // 1V8
	SO_184 pin.Pin    = pin.V1_8      // 1V8
	SO_185 pin.Pin    = pin.V1_8      // 1V8
	SO_186 pin.Pin    = pin.V1_8      // 1V8
	SO_187 pin.Pin    = pin.GROUND    // GND
	SO_188 pin.Pin    = pin.GROUND    // GND
	SO_189 pin.Pin    = pin.DC_IN     // VDAC
	SO_190 pin.Pin    = pin.DC_IN     // VDAC
	SO_191 pin.Pin    = pin.V3_3      // 3V3
	SO_192 pin.Pin    = pin.V3_3      // 3V3
	SO_193 pin.Pin    = pin.V3_3      // 3V3
	SO_194 pin.Pin    = pin.V3_3      // 3V3
	SO_195 pin.Pin    = pin.GROUND    // GND
	SO_196 pin.Pin    = pin.GROUND    // GND
	SO_197 pin.Pin    = pin.DC_IN     // VBAT
	SO_198 pin.Pin    = pin.DC_IN     // VBAT
	SO_199 pin.Pin    = pin.DC_IN     // VBAT
	SO_200 pin.Pin    = pin.DC_IN     // VBAT
)

// define a reference to the pins in an array for the UI
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
	SO_3.(*Pin),
	SO_5.(*Pin),
	SO_9.(*Pin),
	SO_11.(*Pin),
	SO_15.(*Pin),
	SO_17.(*Pin),
	SO_21.(*Pin),
	SO_23.(*Pin),
	SO_27.(*Pin),
	SO_28.(*Pin),
	SO_29.(*Pin),
	SO_30.(*Pin),
	SO_33.(*Pin),
	SO_34.(*Pin),
	SO_35.(*Pin),
	SO_36.(*Pin),
	SO_45.(*Pin),
	SO_46.(*Pin),
	SO_47.(*Pin),
	SO_48.(*Pin),
	SO_51.(*Pin),
	SO_52.(*Pin),
	SO_53.(*Pin),
	SO_54.(*Pin),
	SO_57.(*Pin),
	SO_58.(*Pin),
	SO_59.(*Pin),
	SO_60.(*Pin),
	SO_63.(*Pin),
	SO_64.(*Pin),
	SO_65.(*Pin),
	SO_66.(*Pin),
	SO_60.(*Pin),
	SO_69.(*Pin),
	SO_70.(*Pin),
	SO_71.(*Pin),
	SO_72.(*Pin),
	SO_75.(*Pin),
	SO_76.(*Pin),
	SO_77.(*Pin),
	SO_78.(*Pin),
	SO_81.(*Pin),
	SO_82.(*Pin),
	SO_83.(*Pin),
	SO_84.(*Pin),
	SO_87.(*Pin),
	SO_89.(*Pin),
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
