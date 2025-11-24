package main

import (
	"learn-go/mocking/countdown"
	"os"
)

func main() {
	sleeper := &countdown.DefaultSleeper{}
	countdown.Countdown(os.Stdout, sleeper)
}
