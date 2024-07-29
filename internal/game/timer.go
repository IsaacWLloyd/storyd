package game

import (
	"time"
)

type Timer struct {
	Duration  time.Duration
	StartTime time.Time
}

func NewTimer(duration time.Duration) *Timer {
	return &Timer{
		Duration: duration,
	}
}

func (t *Timer) Start() {
	t.StartTime = time.Now()
}

func (t *Timer) Reset(duration time.Duration) {
	t.Duration = duration
	t.Start()
}

func (t *Timer) IsExpired() bool {
	return time.Since(t.StartTime) > t.Duration
}

func (t *Timer) RemainingTime() time.Duration {
	elapsed := time.Since(t.StartTime)
	if elapsed > t.Duration {
		return 0
	}
	return t.Duration - elapsed
}
