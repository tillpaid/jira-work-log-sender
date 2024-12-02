package main

import (
	"fmt"

	"github.com/rthornton128/goncurses"
	"github.com/tillpaid/paysera-log-time-golang/internal/app"
	"github.com/tillpaid/paysera-log-time-golang/internal/jira"
	"github.com/tillpaid/paysera-log-time-golang/internal/resource"
	"github.com/tillpaid/paysera-log-time-golang/internal/service"
	"github.com/tillpaid/paysera-log-time-golang/internal/ui"
	"github.com/tillpaid/paysera-log-time-golang/internal/ui/pages"
)

func main() {
	config, loading, screen, client := initResources()

	if err := app.StartApp(client, config, screen, loading); err != nil {
		ui.EndScreen()
		service.PrintFatalError(err)
	}
}

func initResources() (*resource.Config, *pages.Loading, *goncurses.Window, *jira.Client) {
	screen, err := ui.InitializeScreen()
	if err != nil {
		ui.EndScreen()
		service.PrintFatalError(fmt.Errorf("error initializing screen: %v", err))
	}

	loading := pages.NewLoading(screen)
	loading.PrintBorder()

	loading.PrintRow("Initializing configuration...", 0)
	config, err := resource.InitConfig()
	if err != nil {
		ui.EndScreen()
		service.PrintFatalError(err)
	}

	loading.PrintRow("Initializing jira client...", 0)
	client, err := jira.NewClient(config)
	if err != nil {
		ui.EndScreen()
		service.PrintFatalError(fmt.Errorf("error initializing jira client; %v", err))
	}

	return config, loading, screen, client
}
