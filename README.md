# Periph.io RPI GPIO Simulator
Simple drop in replacement for Perip.io Raspbery pi host.  When it is not possible to test on a real raspbery pi the simulator will show the status of the GPIO pins.

High GPIO pins will be shown in red, Low pins in white.

```text
######### Raspbery Pi Model Zero GPIO simulator ###########

 +| +| -|14|15|18| -|23|24| -|25|08|07|01| -|12| -|16|20|21
         ## ## ##    ## ##    ## ## ## ##    ##    ## ## ##
   ## ## ##    ##    ##    ## ## ##    ## ## ## ## ## ##
 +|02|03|04| -|17|27|22| +|10|09|11| -|00|05|06|13|19|16| -


Output:
2018-11-09 17:31:32.220435 +0000 GMT m=+12.025952666 Set pin High
2018-11-09 17:31:33.222656 +0000 GMT m=+13.028164482 Set pin Low
2018-11-09 17:31:34.223266 +0000 GMT m=+14.028764333 Set pin High
2018-11-09 17:31:35.226843 +0000 GMT m=+15.032331878 Set pin Low
2018-11-09 17:31:36.231492 +0000 GMT m=+16.036971008 Set pin High
2018-11-09 17:31:37.236634 +0000 GMT m=+17.042103158 Set pin Low
2018-11-09 17:31:38.240648 +0000 GMT m=+18.046107154 Set pin High
2018-11-09 17:31:39.243725 +0000 GMT m=+19.049173638 Set pin Low
2018-11-09 17:31:40.248845 +0000 GMT m=+20.054284152 Set pin High
2018-11-09 17:31:41.253253 +0000 GMT m=+21.058682455 Set pin Low
2018-11-09 17:31:42.254859 +0000 GMT m=+22.060278559 Set pin High
```

## Usage
To use the simulator replace the standard import for the Raspberry Pi host:  
`"periph.io/x/periph/host/rpi"`  
with this version:   
`"github.com/nicholasjackson/periph-gpio-simulator/host/rpi"`

```go
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
  // create a custom logger which uses the redirected std out otherwise
  // the output display will be messed up
	logger := log.New(os.Stdout, "", log.LUTC)

	// Load all the drivers:
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

  // alternate pins low and high every 500ms
  go func() {
    lev := gpio.High
	  t := time.Tick(500 * time.Millisecond)

	  for range t {
	    logger.Println("Set pin",lev)
	    rpi.P1_15.Out(lev)

      if lev == gpio.High {
        lev = gpio.Low
      } else {
        lev = gpio.High
      }
    }
  }()

  // Block until a key is pressed
  c := make(chan os.Signal, 1)
  signal.Notify(c)

  s := <-c
}
```
