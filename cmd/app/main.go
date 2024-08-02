package main

import (
	"fmt"
	"github.com/rthornton128/goncurses"
	"github.com/tillpaid/paysera-log-time-golang/internal/app"
	"github.com/tillpaid/paysera-log-time-golang/internal/resource"
	"github.com/tillpaid/paysera-log-time-golang/internal/service"
	"github.com/tillpaid/paysera-log-time-golang/internal/ui"
)

func main() {
	config, screen := initResources()
	defer ui.EndScreen()

	if err := app.StartApp(config, screen); err != nil {
		ui.EndScreen()
		service.PrintFatalError(err)
	}
}

func initResources() (*resource.Config, *goncurses.Window) {
	config, err := resource.InitConfig()
	if err != nil {
		service.PrintFatalError(err)
	}

	screen, err := ui.InitializeScreen()
	if err != nil {
		service.PrintFatalError(fmt.Errorf("error initializing screen: %v", err))
	}

	return config, screen
}
