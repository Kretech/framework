package console

import (
	"flag"
	"fmt"
	"io"
	"os"
)

type RootCommand struct {
	BaseCommand
}

func (*RootCommand) Name() string {
	return `krete`
	return os.Args[0]
}

func (*RootCommand) Note() string {
	return `default node`
}

func (c *RootCommand) Handle(args []string, w io.Writer) {

	f := flag.NewFlagSet(c.Name(), flag.ContinueOnError)
	f.SetOutput(w)

	v := f.Bool(`v`, false, `Display this application version`)

	f.Parse(args)

	if *v {
		fmt.Fprintln(w, 0.1)
		return
	}

	f.Usage()
	fmt.Fprintln(w)

	if len(c.kernel.ListCommands()) > 0 {
		fmt.Fprintln(w, cYellow("Available commands:"))
		for _, c := range c.kernel.ListCommands() {
			fmt.Fprintf(w, "  %-32s%s\n", cGreen(c.Name()), c.Note())
		}
	}
}

func cYellow(text string) string {
	return fmt.Sprintf("\x1b[0;33m%s\x1b[0m", text)
}

func cGreen(text string) string {
	return fmt.Sprintf("\x1b[0;32m%s\x1b[0m", text)
}
