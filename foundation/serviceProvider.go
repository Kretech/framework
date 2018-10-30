package foundation

import (
	"github.com/Kretech/scaffold/container"
)

type ServiceProvider interface {
	Register(*container.Container)
}

// 简单形式的
type ServiceRegister func(*container.Container)

func (f ServiceRegister) Register(c *container.Container) {
	f(c)
}
