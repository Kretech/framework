package foundation

import (
	"github.com/Kretech/framework/event"
	"github.com/Kretech/scaffold/container"
)

var app *Application

type Application struct {
	container.Container
}

func App() *Application {

	if app == nil {
		app = &Application{
			*container.NewContainer(),
		}

		app.registerBaseServiceProviders()
		app.registerCoreContainerAlias()
	}

	return app
}

func (app *Application) registerBaseServiceProviders() {
	app.registerServiceProvider(&event.ServiceProvider{})
}

func (app *Application) registerCoreContainerAlias() {

}

func (app *Application) registerServiceProvider(provider ServiceProvider) {
	provider.Register(&app.Container)
}
