// DCSO Threat Intelligence Engine
// Copyright (c) 2016, DCSO GmbH

package main

import (
	"bufio"
	"dcso.de/iocee"
	"fmt"
	"gopkg.in/urfave/cli.v1"
	"os"
)

func exitWithError(message string) {
	fmt.Fprintf(os.Stderr, "Error: %s \n", message)
	os.Exit(-1)
}

func extractFromStdin(interactive bool) {
	//we determine if the program is run interactively or within a pipe
	stat, _ := os.Stdin.Stat()
	var isTerminal = (stat.Mode() & os.ModeCharDevice) != 0
	//if we are not in an interactive session and this is a terminal, we quit
	if !interactive && isTerminal {
		return
	}
	if interactive {
		fmt.Println("Interactive mode: Enter a blank line [by pressing ENTER] to exit (values will not be stored otherwise).")
	}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" && interactive {
			break
		}
		results := iocee.Parse(line)
		for _, result := range results {
			fmt.Println(result)
		}
	}

}

func mainAction(c *cli.Context) {
	interactive := c.GlobalBool("interactive")
	extractFromStdin(interactive)
}

func main() {

	app := cli.NewApp()
	app.Name = "IOCee"
	app.Usage = "Utility to extract IOCs from text streams"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "interactive",
			Usage: "interactively add values to the filter",
		},
	}
	app.Action = mainAction
	app.Run(os.Args)

}
