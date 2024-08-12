package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	finalWord      = "Go!"
	countDownStart = 3
)

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}
func Countdown(buffer io.Writer, sleeper Sleeper) {
	for i := countDownStart; i > 0; i-- {
		fmt.Fprint(buffer, i)
		sleeper.Sleep()
	}
	fmt.Fprint(buffer, finalWord)
}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
