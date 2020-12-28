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
  -show-signal, -S
      When used with -silent, -s, it makes sleepd show an message if it catches signal.

  -silent, -s
      Silent or quiet mode. Don't show some messages. Makes sleepd mute.
      Use --show-signal, -S in addition to this option to disable info messages but still show signaled messages.

  -help, -h
      Display this help and exit.

  -version, -v
      Display program version and exit.

GitHub repository URL: https://github.com/masa213f/sleepd
`

type option struct {
	waitTime    int
	silent      bool
	showSignal  bool
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
	setFlagBoolBar(flags, &opt.showSignal, false, "show-signal", "S")
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
