// DCSO IOCee IOC Extractor
// Copyright (c) 2017, DCSO GmbH

package main

import (
	"bufio"
	"fmt"
	"os"
	"github.com/DCSO/iocee"
	"gopkg.in/urfave/cli.v1"
)

func exitWithError(message string) {
	fmt.Fprintf(os.Stderr, "Error: %s \n", message)
	os.Exit(-1)
}

func extractFromStdin(interactive bool, withSource bool) {
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

	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" && interactive {
			break
		}
		results := iocee.Parse(line)
		for _, result := range results {
			f.WriteString(result)
			if withSource {
				f.WriteString("\t")
				f.WriteString(line)
			}
			f.WriteString("\n")
		}
	}

}

func mainAction(c *cli.Context) {
	interactive := c.GlobalBool("interactive")
	withSource := c.GlobalBool("with-source")
	extractFromStdin(interactive, withSource)
}

func main() {

	app := cli.NewApp()
	app.Name = "IOCee"
	app.Usage = "Utility to extract IOCs from text streams"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "interactive,i",
			Usage: "interactively add values to the filter",
		},
		cli.BoolFlag{
			Name:  "with-source,s",
			Usage: "return the line in which a given IOC was found with the IOC",
		},
	}
	app.Action = mainAction
	app.Run(os.Args)

}
