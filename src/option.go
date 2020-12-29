package main

import (
	"flag"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const usage = `sleepd - pause for a specified amount of time

Usage:
  sleepd [OPTION]... [NUMBER]

Pause for NUMBER seconds.

Options:
  -exit-code [NUMBER], -e [NUMBER]
      Exit code at the end of sleepd.
      Even if this option is set, exit with the signal number when a signal is received.

  -ignore-signals, -i
      Ignored signals. It should be a comma-separated list of signal numbers.

  -log-interval [NUMBER], -interval [NUMBER]
      Period of the log output. Specifies an integer greater than 1. The unit is seconds. (default: 10 seconds)

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
	waitTime      int
	exitCode      int
	ignoreSignals []int
	logInterval   int
	silent        bool
	showSignal    bool
	showVersion   bool
	showHelp      bool
}

func setFlagBoolBar(flags *flag.FlagSet, p *bool, defaultValue bool, name ...string) {
	for _, n := range name {
		flags.BoolVar(p, n, defaultValue, "")
	}
}

func setFlagStringBar(flags *flag.FlagSet, p *string, defaultValue string, name ...string) {
	for _, n := range name {
		flags.StringVar(p, n, defaultValue, "")
	}
}

func parseOptions(args []string) (*option, error) {
	opt := new(option)
	var rawExitCode string
	var rawIgnoreSignals string
	var rawLogInterval string

	var flags = flag.NewFlagSet("", flag.ContinueOnError)
	setFlagStringBar(flags, &rawExitCode, "0", "exit-code", "e")
	setFlagStringBar(flags, &rawIgnoreSignals, "", "ignore-signals", "i")
	setFlagStringBar(flags, &rawLogInterval, "10", "log-interval", "interval")
	setFlagBoolBar(flags, &opt.silent, false, "silent", "s")
	setFlagBoolBar(flags, &opt.showSignal, false, "show-signal", "S")
	setFlagBoolBar(flags, &opt.showVersion, false, "version", "v")
	setFlagBoolBar(flags, &opt.showHelp, false, "help", "h")

	err := flags.Parse(args)
	if err != nil {
		return nil, err
	}

	exitCode, err := strconv.Atoi(rawExitCode)
	if err != nil {
		return nil, fmt.Errorf("cannot to recognize exit code: %s", rawExitCode)
	}
	opt.exitCode = exitCode

	if len(rawIgnoreSignals) > 0 {
		split := strings.Split(rawIgnoreSignals, ",")
		for _, s := range split {
			sigNum, err := strconv.Atoi(s)
			if err != nil {
				return nil, fmt.Errorf("cannot to recognize signal number: %s", s)
			}
			opt.ignoreSignals = append(opt.ignoreSignals, sigNum)
		}
		sort.Ints(opt.ignoreSignals)
	}

	interval, err := strconv.Atoi(rawLogInterval)
	if err != nil {
		return nil, fmt.Errorf("cannot to recognize log interval: %s", rawLogInterval)
	}
	if interval <= 0 {
		return nil, fmt.Errorf("invalid log interval: %d", interval)
	}
	opt.logInterval = interval

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
