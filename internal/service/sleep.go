package service

import "time"

func SleepMilliseconds(milliseconds int) {
	time.Sleep(time.Duration(milliseconds) * time.Millisecond)
}
