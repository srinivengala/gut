package command

import (
	"github.com/mitchellh/cli" // command and subcommand commandline management
)

// implements interface cli.Command
type classEnableCommand struct {
	UI cli.Ui // cli.Ui interface
}

func (c *classEnableCommand) Help() string {
	return "gut class enable [class]..."
}

func (c *classEnableCommand) Run(args []string) int {
	c.UI.Output("Running class enable ...")
	return 0 //return cli.RunResultHelp to show help and exit
}

func (c *classEnableCommand) Synopsis() string {
	return "enable class(es)"
}
