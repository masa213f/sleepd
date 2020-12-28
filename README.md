# sleepd

`sleepd` is a program for pausing a specified amount of time.

`sleepd` is intended to be used as a dummy application running on Kubernetes.
This command is similar to the `sleep` command in GNU coreutils, but there are some useful features for Kubernetes.

- Logging
    For ease of checking operation, `sleepd` outputs logs to `stdout` periodically (default: every second).
    Events from the outside (e.g. signals) are also output to the log.

- Signal handling
    `sleepd` handles all signals. If catches a signal, `speepd` will exit with the signal number as a return code.
    If some signals want to be ignored, it can be specified as command-line arguments.

- Small size container image
    `sleepd` does not depend on any shared Libraries. So its container image is scratch-based. It's easy to download. 
