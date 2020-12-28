package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

const usage = `sleepd - pause for a specified amount of time

Usage:
  sleepd [OPTION]... [NUMBER]

Pause for NUMBER seconds.

Options:
  -show-error, -S
      When used with -silent, -s, it makes sleepd show an error message if it fails.

  -silent, -s
      Silent or quiet mode. Don't show info or error messages.  Makes sleepd mute.
      Use --show-error, -S in addition to this option to disable info messages but still show error messages.

  -help, -h
      Display this help and exit.

  -version, -v
      Display program version and exit.

GitHub repository URL: https://github.com/masa213f/sleepd
`

type option struct {
	waitTime    int
	silent      bool
	showError   bool
	showVersion bool
	showHelp    bool
}

func setFlagBoolBar(flags *flag.FlagSet, p *bool, defaultValue bool, name ...string) {
	for _, n := range name {
		flags.BoolVar(p, n, defaultValue, "")
	}
}

func parseOptions(args []string) (*option, error) {
	opt := new(option)

	var flags = flag.NewFlagSet("", flag.ContinueOnError)
	setFlagBoolBar(flags, &opt.silent, false, "silent", "s")
	setFlagBoolBar(flags, &opt.showError, false, "show-error", "S")
	setFlagBoolBar(flags, &opt.showVersion, false, "version", "v")
	setFlagBoolBar(flags, &opt.showHelp, false, "help", "h")

	err := flags.Parse(args)
	if err != nil {
		return nil, err
	}

	if flags.NArg() > 1 {
		return nil, fmt.Errorf("too many args: %s", strings.Join(flags.Args(), ", "))
	}

	if flags.NArg() == 1 {
		num, err := strconv.Atoi(flags.Arg(0))
		if err != nil {
			return nil, err
		}
		if num < 0 {
			return nil, fmt.Errorf("invalid number: %d", num)
		}
		opt.waitTime = num
	}

	return opt, nil
}
