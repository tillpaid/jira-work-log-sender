package main

import (
	"github.com/tillpaid/paysera-log-time-golang/internal/service"
	"github.com/tillpaid/paysera-log-time-golang/internal/ui"
	"log"
)

func main() {
	screen, err := ui.InitializeScreen()
	if err != nil {
		log.Fatalf("Error initializing screen: %v", err)
	}
	defer ui.EndScreen()

	workLogs := service.GetWorkLogs()
	ui.DrawTable(screen, workLogs)
	screen.GetChar()
}
