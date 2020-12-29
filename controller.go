package main

import (
	"fmt"
	"time"

	"github.com/nsf/termbox-go"
)

// Default values
const tWork = 5
const tBreak = 5

const animationSpeed = 100 * time.Millisecond

func main() {

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	eventQueue := make(chan termbox.Event)

	// Poll events such as user input and pass it to the channel eventQueue
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

	p := NewPomodoro(tWork, tBreak)

	render(p)

	for {
		select {
		case ev := <-eventQueue:
			if ev.Type == termbox.EventKey {
				switch {
				case ev.Ch == 'q':
					return
				case ev.Ch == 'n':
					p.startWork()
				case ev.Ch == 'b':
					p.startBreak()
				default:
					fmt.Printf("Event %d received.", ev)
				}
			}
		case <-p.timer.C:
			if p.active == work {
				p.startBreak()
			} else if p.active == br {
				p.startWork()
			}

		default:
			render(p)
			time.Sleep(animationSpeed)
		}
	}

}
