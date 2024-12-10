package service

import "github.com/tillpaid/paysera-log-time-golang/internal/ui"

func HandlePanic() {
	if r := recover(); r != nil {
		ui.EndWindow()
		panic(r)
	}
}
