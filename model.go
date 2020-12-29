package main

import "time"

// Interval specifies whether work or break is the active interval
type Interval int

const (
	work Interval = iota

	//break
	br
)

// A Pomodoro stores all information related to a pomodoro timer
type Pomodoro struct {
	tWork, tBreak time.Duration
	timer         *time.Timer
	active        Interval
}

// newPomodoro returns a new instance of a Pomodoro timer
// Arguments tWork and tBreak specify the respective timer intervals in seconds
func newPomodoro(tWork, tBreak int64) *Pomodoro {
	p := new(Pomodoro)
	p.tBreak = time.Duration(tBreak)
	p.tWork = time.Duration(tBreak)
	p.timer = time.NewTimer(p.tWork)
	p.timer.Stop()
	p.active = work
	return p
}

// startWork starts the work interval of the pomodoro timer
func (p *Pomodoro) startWork() {
	p.timer.Reset(p.tWork)
	p.active = work
}

// startBreak starts the break interval of the pomodoro timer
func (p *Pomodoro) startBreak() {
	p.timer.Reset(p.tBreak)
	p.active = br
}

// setWorkInterval changes the work interval of the pomodoro to tWork
func (p *Pomodoro) setWorkInterval(tWork int64) {
	p.tWork = time.Duration(tWork)
}

// setBreakInterval changes the break interval of the pomodoro to tBreak
func (p *Pomodoro) setBreakInterval(tBreak int64) {
	p.tBreak = time.Duration(tBreak)
}
