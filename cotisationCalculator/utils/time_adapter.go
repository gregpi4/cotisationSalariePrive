package utils

import "time"

type TimeOperations interface {
	Sleep(d time.Duration)
}

type Time struct {
}

func (t *Time) Sleep(d time.Duration) {
	time.Sleep(d)
}
