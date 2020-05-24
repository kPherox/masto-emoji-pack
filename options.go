package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/pborman/getopt/v2"
)

type Options struct {
	Help    bool
	Version bool
	Servers []string
	Split   bool
}

func Help() {
	getopt.PrintUsage(os.Stdout)
	os.Exit(0)
}

func Version() {
	fmt.Fprintln(os.Stdout, "masto-emoji-pack v0.0.1")
	os.Exit(0)
}

func Usage(err error) {
	fmt.Fprintln(os.Stderr, err)
	getopt.Usage()
	os.Exit(2)
}

func Parse() {
	if err := getopt.Getopt(nil); err != nil {
		Usage(err)
	}
}

func parseOptions() (options Options) {
	getopt.SetParameters("DOMAIN...")
	help := getopt.BoolLong("help", 'h', "show help message")
	version := getopt.BoolLong("version", 'v', "show version info")
	split := getopt.BoolLong("split", 's', "split emoji pack via category")

	Parse()

	options = Options{
		Help:    *help,
		Version: *version,
		Servers: getopt.Args(),
		Split:   *split,
	}

	if options.Help {
		Help()
	}
	if options.Version {
		Version()
	}
	if len(options.Servers) == 0 {
		Usage(errors.New("must be specified: DOMAIN..."))
	}

	return
}
