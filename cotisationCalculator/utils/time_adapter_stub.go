package utils

import "time"

type TestTime struct {
}

func NewTestTime() TestTime {
	return TestTime{}
}

func (t *TestTime) Sleep(d time.Duration) {
	// do nothing
}
