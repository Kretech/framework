package event

import (
	"github.com/Kretech/scaffold/container"
	"github.com/Kretech/scaffold/events"
)

func GetRegister() func(*container.Container) {
	return func(container *container.Container) {
		container.Singleton(`events`, func() interface{} {
			return events.NewDispatcher()
		})
	}
}
