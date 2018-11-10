package rpi

import (
	"bufio"
	"io"
	"os"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
	termbox "github.com/nsf/termbox-go"
	"periph.io/x/periph/conn/gpio"
)

var piZero = `
######### Raspbery Pi Model Zero GPIO simulator ###########

 +| +| -|14|15|18| -|23|24| -|25|08|07|01| -|12| -|16|20|21
				 ## ## ##    ## ##    ## ## ## ##    ##    ## ## ##
	 ## ## ##    ## ## ##    ## ## ##    ## ## ## ## ## ##
 +|02|03|04| -|17|27|22| +|10|09|11| -|00|05|06|13|19|16| -
`

var termBoxInitalized bool
var output *os.File

func init() {
	termBoxInitalized = true

	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	draw()

	os.Remove("./out.log")

	// redirect stdout to the logger
	output, err = os.OpenFile("./out.log", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	os.Stdout = output
	os.Stderr = output

	go startRefresh()

	logLines = make([]string, 0)
	go tailLog("./out.log")
}

func startRefresh() {
	t := time.Tick(200 * time.Millisecond)
	for range t {
		draw()
	}
}

func tailLog(filename string) error {
	for i, r := range "Output:" {
		termbox.SetCell(i, 9, r, termbox.ColorYellow, termbox.ColorBlack)
	}

	file, _ := os.Open(filename)
	watcher, _ := fsnotify.NewWatcher()
	defer watcher.Close()
	_ = watcher.Add(filename)

	r := bufio.NewReader(file)
	for {
		by, err := r.ReadBytes('\n')
		if err != nil && err != io.EOF {
			return err
		}

		writeLogline(string(by))

		if err != io.EOF {
			continue
		}
		if err = waitForChange(watcher); err != nil {
			return err
		}
	}
}

var logLines []string

func writeLogline(l string) {
	if l == "" {
		return
	}

	// TODO: if log is more than 120 characters split onto multiple lines

	if len(logLines) > 10 {
		logLines = logLines[1:]
	}

	logLines = append(logLines, time.Now().String()+" "+l)

	for r, line := range logLines {
		for c, b := range line {
			termbox.SetCell(c, r+10, rune(b), termbox.ColorWhite, termbox.ColorBlack)
		}
	}
}

func waitForChange(w *fsnotify.Watcher) error {
	for {
		select {
		case event := <-w.Events:
			if event.Op&fsnotify.Write == fsnotify.Write {
				return nil
			}
		case err := <-w.Errors:
			return err
		}
	}
}

func draw() {
	lines := strings.Split(piZero, "\n")

	// print the header
	for row, l := range lines[:4] {
		for col, r := range l {
			termbox.SetCell(col, row, r, termbox.ColorWhite, termbox.ColorBlack)
		}
	}

	// print the pins in the first row
	drawPins(lines[3], 4)
	drawPins(lines[6], 5)

	for col, r := range lines[6] {
		termbox.SetCell(col, 6, r, termbox.ColorWhite, termbox.ColorBlack)
	}

	termbox.Flush()
}

func drawPins(line string, index int) {
	for _, p := range pins {
		pos := strings.Index(line, p.Name())
		var color = termbox.ColorWhite

		if pos != -1 {
			if p.Read() == gpio.High {
				color = termbox.ColorRed
			}

			termbox.SetCell(pos+1, index, '#', color, termbox.ColorBlack)
			termbox.SetCell(pos+2, index, '#', color, termbox.ColorBlack)
		}
	}
}
