package rpi

import (
	"bufio"
	"bytes"
	"io"
	"log"
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
var eventQueue chan termbox.Event

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

	eventQueue = make(chan termbox.Event)
	go monitorEvents()
}

func monitorEvents() {
	for {
		eventQueue <- termbox.PollEvent()
	}
}

func startRefresh() {
	defer termbox.Close()

	ticker := time.NewTicker(200 * time.Millisecond)

	for {
		select {
		case e := <-eventQueue:
			if e.Key == termbox.KeyCtrlC {
				log.Println("Exit")
				ticker.Stop()
				os.Exit(0)
			}
		case <-ticker.C:
			draw()
		}
	}
}

func tailLog(filename string) error {
	for i, r := range "Log (log is also written to ./out.log:" {
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
var logLineMax = 120

func writeLogline(l string) {
	if l == "" {
		return
	}

	// Split the message by max log length
	lines := splitSubN(l, 120)

	// remove old lines from the buffer
	if len(logLines)+len(lines) > 10 {
		logLines = logLines[len(lines):]
	}

	// append the lines to the buffer
	for _, line := range lines {
		logLines = append(logLines, line)
	}

	for r, line := range logLines {
		for c := 0; c < logLineMax; c++ {
			ru := ' '
			if c < len(line) {
				ru = rune(line[c])
			}
			termbox.SetCell(c, r+10, ru, termbox.ColorWhite, termbox.ColorBlack)
		}
	}
}

func splitSubN(s string, n int) []string {
	sub := ""
	subs := []string{}

	runes := bytes.Runes([]byte(s))
	l := len(runes)
	for i, r := range runes {
		sub = sub + string(r)
		if (i+1)%n == 0 {
			subs = append(subs, sub)
			sub = ""
		} else if (i + 1) == l {
			subs = append(subs, sub)
		}
	}

	return subs
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

	// set the background
	for c := 0; c < 59; c++ {
		termbox.SetCell(c, index, ' ', termbox.ColorBlack, termbox.ColorBlack)
	}

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
