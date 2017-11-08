package console

import "github.com/Kretech/contracts/console"

type BaseCommand struct {
	kernel console.Kernel
}

func (c *BaseCommand) SetKernel(kernel console.Kernel) {
	c.kernel = kernel
}
