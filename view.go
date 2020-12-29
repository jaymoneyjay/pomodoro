package main

import (
	"fmt"
	"time"

	"github.com/nsf/termbox-go"
)

// Layout
const backgroundColor = termbox.ColorBlack
const foregroundColor = termbox.ColorWhite
const defaultColor = termbox.ColorGreen
const width = 50
const height = 30

// Text
var instructions = []string{
	"n: start timer",
	"b: start break",
	"q: quit",
	"s: settings",
}

const title = "POMODORO TIMER"

// Renders the pomodoro timer with the instructions
func render(p *Pomodoro) {
	termbox.Clear(foregroundColor, backgroundColor)

	// Render title
	tbPrint(width/2, 1, defaultColor, backgroundColor, title)

	// Render instructions
	for y, instr := range instructions {
		tbPrint(width, y+1, defaultColor, backgroundColor, instr)
	}

	// Render timer
	var activeString string
	formatedTime := formatSeconds(p.timer.TimeRemaining())
	if p.active == work {
		activeString = fmt.Sprintf("Break in %s", formatedTime)
	} else {
		activeString = fmt.Sprintf("Work in %s", formatedTime)
	}

	tbPrint(width/2, 3, defaultColor, backgroundColor, activeString)

	termbox.Flush()
}

func tbPrint(x, y int, fg, bg termbox.Attribute, text string) {
	for _, c := range text {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}

func formatSeconds(timeRemaining time.Duration) string {
	minutes := int(timeRemaining) / int(time.Minute)
	seconds := (int(timeRemaining) % int(time.Minute)) / int(time.Second)
	return fmt.Sprintf("%d:%d", minutes, seconds)
}
