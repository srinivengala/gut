package command

import (
	"github.com/mitchellh/cli" // command and subcommand commandline management
)

// implements interface cli.Command
type putCommand struct {
	UI cli.Ui // cli.Ui interface
}

func (c *putCommand) Help() string {
	return "put file in the gut repository"
}

func (c *putCommand) Run(args []string) int {
	c.UI.Output("Running put ...")
	return 0 //return cli.RunResultHelp to show help and exit
}

func (c *putCommand) Synopsis() string {
	return "commit file"
}
