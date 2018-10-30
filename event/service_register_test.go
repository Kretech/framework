package event

import (
	"testing"

	containerPkg "github.com/Kretech/scaffold/container"

	"github.com/Kretech/scaffold/events"
	. "github.com/smartystreets/goconvey/convey"
)

func TestServiceProvider(t *testing.T) {

	Convey("TestEventServiceProvider", t, func() {
		c := containerPkg.NewContainer()
		p := GetRegister()
		p(c)

		dispatcher := c.Make(`events`).(*events.Dispatcher)

		dispatcher.ListenFunc(`start`, func() {})
		dispatcher.Dispatch(`start`)
		// dispatcher.TestTool.ShouldDispatched(`start`)
	})
}
