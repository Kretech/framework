package event

import (
	"github.com/Kretech/scaffold/container"
	"github.com/Kretech/scaffold/events"
)

type ServiceProvider struct {
}

func NewServiceProvider() *ServiceProvider {
	return &ServiceProvider{}
}

func (*ServiceProvider) Register(container *container.Container) {
	container.Singleton(`events`, func() interface{} {
		return events.NewDispatcher()
	})
}
