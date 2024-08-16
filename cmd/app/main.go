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
	config, screen, _ := initResources()

	if err := app.StartApp(config, screen); err != nil {
		ui.EndScreen()
		service.PrintFatalError(err)
	}
}

func initResources() (*resource.Config, *goncurses.Window, *jira.Client) {
	config, err := resource.InitConfig()
	if err != nil {
		service.PrintFatalError(err)
	}

	client, err := jira.NewClient(config)
	if err != nil {
		service.PrintFatalError(fmt.Errorf("error initializing jira client; %v", err))
	}

	screen, err := ui.InitializeScreen()
	if err != nil {
		service.PrintFatalError(fmt.Errorf("error initializing screen: %v", err))
	}

	return config, screen, client
}
