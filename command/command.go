package command

import (
	"log"
	"os"

	"github.com/mitchellh/cli" // command and subcommand commandline management
)

func basicUI() *cli.BasicUi {
	return &cli.BasicUi{
		Reader:      os.Stdin,
		Writer:      os.Stdout,
		ErrorWriter: os.Stderr,
	}
}

func coloredUI() *cli.ColoredUi {
	return &cli.ColoredUi{
		Ui:          basicUI(),
		OutputColor: cli.UiColorGreen,
	}
}

func putCommandFactory() (cli.Command, error) {
	return &putCommand{
		UI: coloredUI(),
	}, nil
}

func getCommandFactory() (cli.Command, error) {
	return &getCommand{
		UI: coloredUI(),
	}, nil
}

// Run executes commnads
func Run(args []string) int {
	c := cli.NewCLI("gut", "1.0.0")
	c.Args = args[:]
	c.Commands = map[string]cli.CommandFactory{
		"put": putCommandFactory,
		"get": getCommandFactory,
		// or use lambda
		"search": func() (cli.Command, error) {
			return &searchCommand{
				UI: &cli.ColoredUi{
					Ui: &cli.BasicUi{
						Reader:      os.Stdin,
						Writer:      os.Stdout,
						ErrorWriter: os.Stderr,
					},
					OutputColor: cli.UiColorGreen,
				},
			}, nil
		},
		"kill": func() (cli.Command, error) {
			return &killCommand{
				UI: coloredUI(),
			}, nil
		},
		"server": func() (cli.Command, error) {
			return &serverCommand{
				UI:       coloredUI(),
				HttpPort: 9443,
			}, nil
		},
		"class": func() (cli.Command, error) {
			return &classCommand{
				UI: coloredUI(),
			}, nil
		},
		"class enable": func() (cli.Command, error) {
			return &classEnableCommand{
				UI: coloredUI(),
			}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	return exitStatus
}
