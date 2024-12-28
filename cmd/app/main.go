package main

import (
	"fmt"

	"github.com/tillpaid/jira-work-log-sender/internal/app"
	"github.com/tillpaid/jira-work-log-sender/internal/app/action"
	"github.com/tillpaid/jira-work-log-sender/internal/jira"
	"github.com/tillpaid/jira-work-log-sender/internal/resource"
	"github.com/tillpaid/jira-work-log-sender/internal/service"
	"github.com/tillpaid/jira-work-log-sender/internal/ui"
)

func main() {
	defer service.HandlePanic()

	application, err := initResources()
	if err != nil {
		ui.EndWindow()
		service.PrintFatalError(err)
	}

	if err := application.Start(); err != nil {
		ui.EndWindow()
		service.PrintFatalError(err)
	}
}

func initResources() (*app.Application, error) {
	window, err := ui.InitializeWindow()
	if err != nil {
		return nil, fmt.Errorf("error initializing window: %v", err)
	}

	config, err := resource.InitConfig()
	if err != nil {
		return nil, fmt.Errorf("error initializing config: %v", err)
	}

	client, err := jira.NewClient(config)
	if err != nil {
		return nil, fmt.Errorf("error initializing jira client: %v", err)
	}

	userInput := app.NewUserInput(window)
	actions := action.NewActions(client, window, config)

	return app.NewApplication(window, client, userInput, actions, config), nil
}
