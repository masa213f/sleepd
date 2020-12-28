# sleepd

`sleepd` is a program for pausing a specified amount of time.

sleepd is similar to the `sleep` command in GNU coreutils, but this program is intended to be used as a dummy application for test.
So it has some useful features for testing.

## Features

- Logging

    For ease of checking operation, sleepd outputs logs to `stdout` periodically (default: every second).
    Events from the outside (e.g. signals) are also output to the log.

- Signal handling

    sleepd handles all signals. If catches a signal, speepd will exit with the signal number as a return code.
    If some signals want to be ignored, it can be specified as command-line arguments.

- Small size container image

    sleepd does not depend on any shared libraries. So its container image is scratch-based. It's easy to download.

## Usage

sleepd is uploadted to Docker Hub. So just run the following command to use it.

```console
$  docker run masa213f/sleepd
```
