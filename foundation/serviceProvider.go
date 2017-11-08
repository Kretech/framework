package foundation

import (
	"github.com/Kretech/scaffold/container"
)

type ServiceProvider interface {
	Register(*container.Container)
}
