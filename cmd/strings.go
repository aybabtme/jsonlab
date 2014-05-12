package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/cli"
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	srcFlag = "src"
	dstFlag = "dst"
	Strings = cli.Command{
		Name:  "strings",
		Usage: "manipulates JSON strings",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  srcFlag,
				Value: "",
				Usage: "Source file to read the JSON string from. Default to stdin.",
			},
			cli.StringFlag{
				Name:  dstFlag,
				Value: "",
				Usage: "Destination to write the unescaped string to. Default to stdout.",
			},
		},

		Subcommands: []cli.Command{
			cli.Command{
				Name:   "unescape",
				Usage:  "takes a JSON encoded string and prints if as a decoded string",
				Action: stringUnescape,
			},

			cli.Command{
				Name:   "escape",
				Usage:  "takes a string and encode it to an escaped JSON string",
				Action: stringEscape,
			},
		},

		Action: func(c *cli.Context) { cli.ShowCommandHelp(c, c.Command.Name) },
	}
)

func withOutIn(c *cli.Context, f func(out io.Writer, in io.Reader)) {

	var in *os.File
	var out *os.File

	if c.IsSet(srcFlag) {
		var err error
		in, err = os.Open(c.String(srcFlag))
		if err != nil {
			log.Fatalf("error: opening source file, %v", err)
		}
	} else {
		in = os.Stdin
	}
	defer in.Close()

	if c.IsSet(dstFlag) {
		var err error
		out, err = os.Open(c.String(dstFlag))
		if err != nil {
			log.Fatalf("error: opening destination file, %v", err)
		}
	} else {
		out = os.Stdout
	}
	defer out.Close()

	f(out, in)
}

func stringUnescape(c *cli.Context) {
	withOutIn(c, func(out io.Writer, in io.Reader) {
		var dummy struct{ Val string }

		if err := json.NewDecoder(in).Decode(&dummy.Val); err != nil {
			log.Fatalf("error: decoding JSON string, %v", err)
		}

		_, err := fmt.Fprint(out, dummy.Val)
		if err != nil {
			log.Fatalf("error: writing to output, %v", err)
		}
	})
}

func stringEscape(c *cli.Context) {
	withOutIn(c, func(out io.Writer, in io.Reader) {
		val, err := ioutil.ReadAll(in)
		if err != nil {
			log.Fatalf("error: reading from input, %v", err)
		}

		if err := json.NewEncoder(out).Encode(string(val)); err != nil {
			log.Fatalf("error: encoding JSON string, %v", err)
		}
	})
}
