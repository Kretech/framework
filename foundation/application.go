package foundation

import (
	"github.com/Kretech/contracts/config"
	"github.com/Kretech/framework/event"
	"github.com/Kretech/scaffold/container"
)

var app *Application

type Application struct {
	container.Container

	Conf config.Config
}

func App() *Application {

	if app == nil {
		app = &Application{
			Container: *container.NewContainer(),
		}

		app.registerBaseServiceProviders()
		app.registerCoreContainerAlias()
	}

	return app
}

func (app *Application) registerBaseServiceProviders() {
	//app.RegisterServiceProvider(&configImpl.TomlServiceProvider{})
	//app.SetConfig(app.Make("config").(config.Config))

	app.RegisterServiceProvider(&event.ServiceProvider{})
}

func (app *Application) registerCoreContainerAlias() {

}

func (app *Application) RegisterServiceProvider(provider ServiceProvider) {
	provider.Register(&app.Container)
}

func (app *Application) SetConfig(conf config.Config) {
	app.Conf = conf
}
