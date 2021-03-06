package command

import (
	"flag"
	"fmt"

	"github.com/mitchellh/cli" // command and subcommand commandline management
	"github.com/srinivengala/gut/ui"
)

// implements interface cli.Command
type serverCommand struct {
	UI       cli.Ui // cli.Ui interface
	HTTPPort int
}

func (c *serverCommand) parsePort(args []string) error {
	argflags := flag.NewFlagSet("server", flag.ContinueOnError)
	argflags.Usage = func() { c.UI.Output(c.Help()) }

	argflags.IntVar(&c.HTTPPort, "port", 9443, "Port for server to listen")
	if err := argflags.Parse(args); err != nil {
		//return cli.RunResultHelp
		return err
	}
	return nil
}

func (c *serverCommand) Help() string {
	return fmt.Sprintf(`
    gut server [options]

Options:
    port : default port %d

Example:
    gut server --port=9443`, c.HTTPPort)
}

func (c *serverCommand) Run(args []string) int {
	if err := c.parsePort(args); err != nil {
		return 0
	}

	c.UI.Output(fmt.Sprintf("Running server https://127.0.0.1:%d", c.HTTPPort))

	return ui.StartServer()
}

func (c *serverCommand) Synopsis() string {
	return "starts server"
}
