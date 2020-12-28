package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"
)

var Version string

func log(format string, a ...interface{}) {
	fmt.Println(time.Now().Format(time.RFC3339) + " " + fmt.Sprintf(format, a...))
}

func logInfo(opt *option, format string, a ...interface{}) {
	if !opt.silent {
		log(format, a...)
	}
}

func logSignal(opt *option, format string, a ...interface{}) {
	if !opt.silent || opt.showSignal {
		log(format, a...)
	}
}

func main() {
	opt, err := parseOptions(os.Args[1:])
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		fmt.Println(usage)
		os.Exit(1)
	}
	if opt.showHelp {
		fmt.Println(usage)
		os.Exit(0)
	}
	if opt.showVersion {
		fmt.Printf("sleepd %s\n", Version)
		os.Exit(0)
	}

	var signals []os.Signal
	for sig := range signalToInt {
		signals = append(signals, sig)
	}
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, signals...)

	ignoreSignal := map[int]bool{}
	for _, sig := range opt.ignoreSignals {
		ignoreSignal[sig] = true
	}

	timer := time.NewTimer(time.Duration(opt.waitTime) * time.Second)
	if opt.waitTime <= 0 {
		timer.Stop()
	}

	ticker := time.NewTicker(time.Second)
	if opt.silent {
		ticker.Stop()
	}

	if len(opt.ignoreSignals) == 0 {
		logInfo(opt, "start (catch all signals)")
	} else {
		var sigs []string
		for _, s := range opt.ignoreSignals {
			sigs = append(sigs, strconv.Itoa(s))
		}
		logInfo(opt, "start (ignore: %s)", strings.Join(sigs, ", "))
	}

	done := make(chan int, 1)
	go func() {
		for {
			select {
			case sig := <-sigCh:
				sigNum := getSignalNumber(sig)
				if ignoreSignal[sigNum] {
					logSignal(opt, "catch signal: %d (ignored)", sigNum)
				} else {
					logSignal(opt, "catch signal: %d", sigNum)
					done <- sigNum
				}
			case <-ticker.C:
				logInfo(opt, "sleeping...")
			case <-timer.C:
				done <- 0
			}
		}
	}()
	ret := <-done

	logInfo(opt, "exit (code: %d)", ret)
	os.Exit(ret)
}
