package main

import (
	"github.com/tillpaid/paysera-log-time-golang/internal/import_data"
	"github.com/tillpaid/paysera-log-time-golang/internal/resource"
	"github.com/tillpaid/paysera-log-time-golang/internal/service"
	"github.com/tillpaid/paysera-log-time-golang/internal/ui"
	"log"
)

func main() {
	config, err := resource.InitConfig()
	if err != nil {
		service.PrintFatalError(err)
	}

	screen, err := ui.InitializeScreen()
	if err != nil {
		log.Fatalf("Error initializing screen: %v", err)
	}
	defer ui.EndScreen()

	workLogs, err := import_data.ParseWorkLogs(config)
	if err != nil {
		ui.EndScreen()
		service.PrintFatalError(err)
	}

	workLogs = service.ModifyWorkLogsTime(workLogs)

	if err = ui.DrawTable(screen, workLogs); err != nil {
		ui.EndScreen()
		service.PrintFatalError(err)
	}

	screen.GetChar()
}
