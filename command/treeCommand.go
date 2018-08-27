package command

import (
	"github.com/mitchellh/cli" // command and subcommand commandline management
)

// implements interface cli.Command
type treeCommand struct {
	UI cli.Ui // cli.Ui interface
}

func (c *treeCommand) Help() string {
	return "displays gut directory tree for enabled classes of files"
}

func (c *treeCommand) Run(args []string) int {
	c.UI.Output("Running tree ...")
	return 0 //return cli.RunResultHelp to show help and exit
}

func (c *treeCommand) Synopsis() string {
	return "show enabled class files as direcory tree"
}
