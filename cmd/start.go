package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/urfave/cli"
)

func executeCommand(e execCommand) {
	cmd := e.cmd
	args := e.args
	if err := exec.Command(cmd, args...).Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func StartCommand() cli.Command {
	return cli.Command{
		Name:        "start",
		Aliases:     []string{""},
		Usage:       "proxy start",
		Description: "Start CNTLM Proxy",
		Flags: []cli.Flag{
			cli.IntFlag{
				Name:        "port, p",
				Value:       3128,
				Usage:       "set custom CNTLM `PORT`",
				Destination: &port,
			},
			cli.BoolFlag{
				Name:        "all, a",
				Usage:       "set all CNTLM config",
				Destination: &setAll,
			},
		},
		Action: func(c *cli.Context) {
			if port != 3128 {
				UpdateCNTLMFile(port)
			}
			fmt.Println(c)
			p := fmt.Sprintf("http://localhost:%v", port)
			a := []string{"config", "set", "proxy", p}
			e := execCommand{cmd: "npm", args: a}
			executeCommand(e)
			fmt.Println("CNTLM Proxy Started On", p)
		},
	}
}
