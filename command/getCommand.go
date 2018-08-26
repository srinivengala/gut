package command

import (
	"github.com/mitchellh/cli" // command and subcommand commandline management
)

// implements interface cli.Command
type getCommand struct {
	UI cli.Ui // cli.Ui interface
}

func (c *getCommand) Help() string {
	return "get file out of the gut repository."
}

func (c *getCommand) Run(args []string) int {
	c.UI.Output("Running get ...")
	return 0 //return cli.RunResultHelp to show help and exit
}

func (c *getCommand) Synopsis() string {
	return "checkout file"
}
