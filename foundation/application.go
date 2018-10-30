package foundation

import (
	"sync"

	"github.com/Kretech/framework/event"
	"github.com/Kretech/scaffold/container"
)

var (
	app         *Application
	appInitOnce sync.Once
)

type Application struct {
	container.Container
}

func init() {
	appInitOnce.Do(func() {
		app = &Application{
			Container: *container.NewContainer(),
		}

		app.registerBaseServiceProviders()
		app.registerCoreContainerAlias()
	})
}

// Default singleton application
func App() *Application {

	// appInitOnce.Do(func() {
	// 	app = &Application{
	// 		Container: *container.NewContainer(),
	// 	}
	//
	// 	app.registerBaseServiceProviders()
	// 	app.registerCoreContainerAlias()
	// })

	return app
}

func (app *Application) registerBaseServiceProviders() {
	app.RegisterService(event.GetRegister())
}

func (app *Application) registerCoreContainerAlias() {

}

// Deprecated
func (app *Application) RegisterServiceProvider(provider ServiceProvider) {
	provider.Register(&app.Container)
}

func (app *Application) RegisterService(register ServiceRegister) {
	register(&app.Container)
}
