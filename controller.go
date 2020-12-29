package main

import (
	"fmt"
	"time"

	"github.com/nsf/termbox-go"
)

func main() {

	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	eventQueue := make(chan termbox.Event)

	// Poll events such as user input and pass it to the channel eventQueue
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

	render()

	for {
		select {
		case ev := <-eventQueue:
			switch {
			case ev.Ch == 'q':
				return
			default:
				fmt.Printf("Event %d received.", ev)
			}

		default:
			render()
			time.Sleep(animationSpeed)
		}
	}

}
