package main

import (
	"fmt"
	"github.com/aybabtme/jsonlab/cmd"
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "jsonlab"
	app.Version = "v0.0.1"
	app.Author = "Antoine Grondin"
	app.Email = "antoinegrondin@gmail.com"
	app.Usage = `A small laboratory to play with your JSON.`
	app.Action = cli.ShowAppHelp
	app.Commands = []cli.Command{
		cmd.Strings,
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "critial error: %v", err)

	}
}
