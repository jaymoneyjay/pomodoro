package main

import (
	"time"

	"github.com/nsf/termbox-go"
)

// Layout
const backgroundColor = termbox.ColorBlack
const foregroundColor = termbox.ColorWhite
const defaultColor = termbox.ColorGreen
const width = 50
const height = 30
const animationSpeed = 100 * time.Millisecond

// Text
var instructions = []string{
	"n: start timer",
	"b: start break",
	"q: quit",
	"s: settings",
}

const title = "POMODORO TIMER"

// Renders the pomodoro timer with the instructions
func render() { //p *Pomodoro) {
	termbox.Clear(foregroundColor, backgroundColor)

	// Render title
	tbPrint(width/2, 1, foregroundColor, backgroundColor, title)

}

func tbPrint(x, y int, fg, bg termbox.Attribute, text string) {
	for _, c := range text {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}
