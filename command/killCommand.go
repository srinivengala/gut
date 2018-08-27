package command

import (
	"github.com/mitchellh/cli" // command and subcommand commandline management
)

// implements interface cli.Command
type killCommand struct {
	UI cli.Ui // cli.Ui interface
}

func (c *killCommand) Help() string {
	return "deletes master key to gut repository"
}

func (c *killCommand) Run(args []string) int {
	c.UI.Output("Running kill ...")
	return 0 //return cli.RunResultHelp to show help and exit
}

func (c *killCommand) Synopsis() string {
	return "delete master key"
}
