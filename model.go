package main

import "time"

// Interval specifies whether work or break is the active interval
type Interval int

const (
	work Interval = iota

	//break
	br
)

// SecondsTimer is a timer modeled in seconds
type SecondsTimer struct {
	timer       *time.Timer
	end         time.Time
	C           <-chan time.Time
	stopped     bool
	stoppedTime time.Duration
}

// NewSecondsTimer returns a new instance of SecondsTimer
func NewSecondsTimer(t time.Duration) *SecondsTimer {
	timer := time.NewTimer(t)
	end := time.Now().Add(t)
	C := timer.C
	timer.Stop()

	return &SecondsTimer{
		timer:   timer,
		end:     end,
		C:       C,
		stopped: true,
	}
}

// Reset sets the SecondsTimer s to the specified time.Duration t
func (s *SecondsTimer) Reset(t time.Duration) {
	s.timer.Reset(t)
	s.end = time.Now().Add(t)
	s.stopped = false
}

// Stop pauses the SecondsTimer s
func (s *SecondsTimer) Stop() {
	s.timer.Stop()
	s.stopped = true
	s.stoppedTime = s.TimeRemaining()
}

// TimeRemaining returns the remaining time of SecondsTimer s in seconds
func (s *SecondsTimer) TimeRemaining() time.Duration {
	if s.stopped {
		return s.stoppedTime
	}
	return s.end.Sub(time.Now())
}

// A Pomodoro stores all information related to a pomodoro timer
type Pomodoro struct {
	tWork, tBreak time.Duration
	timer         *SecondsTimer
	active        Interval
}

// NewPomodoro returns a new instance of a Pomodoro timer
// Arguments tWork and tBreak specify the respective timer intervals in seconds
func NewPomodoro(tWork, tBreak int64) *Pomodoro {
	p := new(Pomodoro)
	p.tBreak = time.Duration(tBreak) * time.Second
	p.tWork = time.Duration(tBreak) * time.Second
	p.timer = NewSecondsTimer(p.tWork)
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
