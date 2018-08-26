package command

import (
	"github.com/mitchellh/cli" // command and subcommand commandline management
)

// implements interface cli.Command
type searchCommand struct {
	UI cli.Ui // cli.Ui interface
}

func (c *searchCommand) Help() string {
	return "search file in the gut. stores file in gut repository."
}

func (c *searchCommand) Run(args []string) int {
	c.UI.Output("Running search ...")
	return 0 //return cli.RunResultHelp to show help and exit
}

func (c *searchCommand) Synopsis() string {
	return "search for file(s)"
}
