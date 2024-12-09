package main

import (
	"fmt"

	"github.com/rthornton128/goncurses"
	"github.com/tillpaid/paysera-log-time-golang/internal/app"
	"github.com/tillpaid/paysera-log-time-golang/internal/jira"
	"github.com/tillpaid/paysera-log-time-golang/internal/resource"
	"github.com/tillpaid/paysera-log-time-golang/internal/service"
	"github.com/tillpaid/paysera-log-time-golang/internal/ui"
)

func main() {
	defer panicHandler()

	config, window, client := initResources()

	if err := app.StartApp(client, config, window); err != nil {
		ui.EndWindow()
		service.PrintFatalError(err)
	}
}

func initResources() (*resource.Config, *goncurses.Window, *jira.Client) {
	window, err := ui.InitializeWindow()
	if err != nil {
		ui.EndWindow()
		service.PrintFatalError(fmt.Errorf("error initializing window: %v", err))
	}

	config, err := resource.InitConfig()
	if err != nil {
		ui.EndWindow()
		service.PrintFatalError(err)
	}

	client, err := jira.NewClient(config)
	if err != nil {
		ui.EndWindow()
		service.PrintFatalError(fmt.Errorf("error initializing jira client; %v", err))
	}

	return config, window, client
}

func panicHandler() {
	if r := recover(); r != nil {
		ui.EndWindow()
		panic(r)
	}
}
