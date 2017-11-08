package console

import (
	"os"

	"strings"

	"io"

	contract "github.com/Kretech/contracts/console"
)

type Kernel struct {
	commands       []contract.Command
	defaultCommand contract.Command
}

func NewKernel() *Kernel {
	k := &Kernel{}

	k.RegisterCommand(&RootCommand{})

	k.defaultCommand = k.commands[0]

	return k
}

func (k *Kernel) RegisterCommand(cmd contract.Command) {
	k.commands = append(k.commands, cmd)
	cmd.SetKernel(k)
}

func (k *Kernel) ListCommands() []contract.Command {
	return k.commands
}

func (k *Kernel) Handle(w io.Writer) {

	var zeroArgs []string

	idx := 1
	for ; idx < len(os.Args); idx++ {
		if strings.HasPrefix(os.Args[idx], `-`) {
			zeroArgs = append(zeroArgs, os.Args[idx])
		} else {
			break
		}
	}

	if len(os.Args) < 2 || len(zeroArgs) > 0 {
		k.defaultCommand.Handle(zeroArgs, w)
	}

	var commandName string
	var commandArgs []string

	if idx < len(os.Args) {
		commandName = os.Args[idx]
		commandArgs = os.Args[idx+1:]

		cmd := k.findCommand(commandName)
		cmd.Handle(commandArgs, w)
	}
}

func (k *Kernel) findCommand(name string) contract.Command {
	for i := range k.commands {
		if k.commands[i].Name() == name {
			return k.commands[i]
		}
	}

	return nil
}
