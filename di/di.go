package di

import (
	"fmt"

	"go.uber.org/dig"

	"main/service"
	"main/tools"
)

type App struct{ run func() }

func (a *App) Run() { a.run() }

func BuildApp() (*App, error) {
	c := dig.New()

	if err := c.Provide(func() service.Zoo { return service.NewMoscowZoo() }); err != nil {
		return nil, err
	}

	if err := c.Provide(service.NewVetClinic); err != nil {
		return nil, err
	}

	if err := c.Provide(tools.NewConsole); err != nil {
		return nil, err
	}

	var app *App
	if err := c.Invoke(func(ui *tools.Console) {
		app = &App{run: ui.Run}
	}); err != nil {
		return nil, fmt.Errorf("invoke: %w", err)
	}

	return app, nil
}
