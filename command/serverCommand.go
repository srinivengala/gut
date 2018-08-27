package command

import (
	"flag"
	"fmt"
	"html"
	"net/http"

	"github.com/mitchellh/cli" // command and subcommand commandline management
)

// implements interface cli.Command
type serverCommand struct {
	UI       cli.Ui // cli.Ui interface
	HttpPort int
}

func (c *serverCommand) parsePort(args []string) error {
	argflags := flag.NewFlagSet("server", flag.ContinueOnError)
	argflags.Usage = func() { c.UI.Output(c.Help()) }

	argflags.IntVar(&c.HttpPort, "port", 9443, "Port for server to listen")
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
    gut server --port=9443`, c.HttpPort)
}

func (c *serverCommand) Run(args []string) int {
	if err := c.parsePort(args); err != nil {
		return 0
	}

	c.UI.Output(fmt.Sprintf("Running server https://127.0.0.1:%d", c.HttpPort))

	http.HandleFunc("/", uiHandler)
	http.HandleFunc("/ui", uiHandler)
	http.HandleFunc("/v1", restHandler)
	if err := http.ListenAndServe(":9443", nil); err != nil {
		return 1
	}
	return 0
}

func (c *serverCommand) Synopsis() string {
	return "delete master key"
}

func uiHandler(w http.ResponseWriter, r *http.Request) {
	msg := fmt.Sprintf(`Hello, %q<br>
	<a href="http://localhost:9443/ui">http://localhost:9443/ui></a><br>
	<a href="http://localhost:9443/v1">http://localhost:9443/v1></a>`, html.EscapeString(r.URL.Path))
	fmt.Fprintf(w, "<html><body><h1>%s</h1></body></html>", msg)
}

func restHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{\"message\":\"Hi there\"}")
}
