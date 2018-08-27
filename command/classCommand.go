package command

import (
	"github.com/mitchellh/cli" // command and subcommand commandline management
)

// implements interface cli.Command
type classCommand struct {
	UI cli.Ui // cli.Ui interface
}

func (c *classCommand) Help() string {
	return "manage file class flags.\nOnly enabled file classes are allowed to be accessed."
}

func (c *classCommand) Run(args []string) int {
	c.UI.Output("Running class ...")
	return 0 //return cli.RunResultHelp to show help and exit
}

func (c *classCommand) Synopsis() string {
	return "manage file class"
}
