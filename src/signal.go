package main

import (
	"os"
	"syscall"
)

var signalToInt = map[os.Signal]int{
	syscall.SIGHUP:    1,
	syscall.SIGINT:    2,
	syscall.SIGQUIT:   3,
	syscall.SIGILL:    4,
	syscall.SIGTRAP:   5,
	syscall.SIGABRT:   6, // syscall.SIGIOT
	syscall.SIGBUS:    7,
	syscall.SIGFPE:    8,
	syscall.SIGKILL:   9,
	syscall.SIGUSR1:   10,
	syscall.SIGSEGV:   11,
	syscall.SIGUSR2:   12,
	syscall.SIGPIPE:   13,
	syscall.SIGTERM:   15,
	syscall.SIGSTKFLT: 16,
	syscall.SIGCHLD:   17, // syscall.SIGCLD
	syscall.SIGCONT:   18,
	syscall.SIGSTOP:   19,
	syscall.SIGTSTP:   20,
	syscall.SIGTTIN:   21,
	syscall.SIGTTOU:   22,
	syscall.SIGURG:    23,
	syscall.SIGXCPU:   24,
	syscall.SIGXFSZ:   25,
	syscall.SIGVTALRM: 26,
	syscall.SIGPROF:   27,
	syscall.SIGWINCH:  28,
	syscall.SIGIO:     29, // syscall.SIGPOLL
	syscall.SIGPWR:    30,
	syscall.SIGUNUSED: 31, // syscall.SIGSYS
}

func getSignalNumber(sig os.Signal) int {
	return signalToInt[sig]
}
